package main
import (
    "fmt"
    "strings"
    "math"
    "time"
)

func getMaxTag(posUnigrams []PosUniGram,previous []float64) string {
    maxVal := 0.0
    maxTag := posUnigrams[0].PosTag
    for idx,j := range posUnigrams {
        if(previous[idx] != 0 && previous[idx] > maxVal) {
            maxVal = previous[idx]
            maxTag = j.PosTag
        }
    }
    return maxTag
}

func printTags(posUnigrams []Ngrams) {
    for _,j := range posUnigrams {
        fmt.Printf("%s ",j.Ngram)
    }
    fmt.Printf("\n\n")
}

func getMaxChunkTag(posUnigrams []Ngrams,previous []float64) string {
    maxVal := 0.0
    maxTag := posUnigrams[0].Ngram
    for idx,j := range posUnigrams {
        if(previous[idx] != 0 && previous[idx] > maxVal) {
            maxVal = previous[idx]
            maxTag = j.Ngram
        }
    }
    return maxTag
}

func getPOSUnseen(database_name string, ngram_length int) float64 {
    frequecy_distribution := getFrequencyAggregation(database_name)
    return float64(frequecy_distribution[1])/float64(ngram_length)
}

func processWord(word string,unseen_prob float64,posUnigrams []PosUniGram, previous []float64) []float64 {
    next := make([]float64,len(previous))
    possibleTags := getWordPosgram(word)
    if (len(possibleTags) == 0) {
        possibleTags = getWordPosgram("nil")
    }
    for idx,j := range posUnigrams {
        for _,k := range possibleTags {
            if (k.PosTag == j.PosTag) {
                //next[idx] = math.Log(float64(k.Count)) - math.Log(float64(j.Count)) + previous[idx]
                next[idx] = float64(k.Count)*previous[idx]/float64(j.Count)
            }
        }
    }
    //fmt.Printf("%v\n",next)
    return next
}

func fastProcessWord(word string,unseen_prob float64,wordPosgram []PosWordGram,posUnigrams []PosUniGram, previous []float64) []float64 {
    next := make([]float64,len(previous))
    possibleTags := make([]PosWordGram,0)
    for _,i := range wordPosgram {
        if (i.Word == word) {
            possibleTags = append(possibleTags,i)
        }
    }

    if (len(possibleTags) == 0) {
        possibleTags = getWordPosgram("nil")
    }
    for idx,j := range posUnigrams {
        for _,k := range possibleTags {
            if (k.PosTag == j.PosTag) {
                //next[idx] = math.Log(float64(k.Count)) - math.Log(float64(j.Count)) + previous[idx]
                next[idx] = float64(k.Count)*previous[idx]/float64(j.Count)
            }
        }
    }
    //fmt.Printf("%v\n",next)
    return next
}

func bulkPoswordgrams(tokens []string) []PosWordGram {
    possibleTags := getAllWordPosgram(tokens)
    return possibleTags
}

func bulkPosgrams(posUnigrams []PosUniGram) []Ngrams {
    list := make([]string,0)
    for _,i := range posUnigrams {
        for _,j := range posUnigrams {
            list =  append(list,i.PosTag+" "+j.PosTag)
        }
    }
    result := getAllNgram(list,"posTags")
    return result
}

func processTag(tag string,val float64,unseen_prob float64, posUnigrams []PosUniGram,previous []float64) []float64 {
    next := make([]float64,82)
    list := make([]string,82)
    for idx,j := range posUnigrams {
        list[idx] =  tag+" "+j.PosTag
    }
    result := getAllNgram(list,"posTags")
    fmt.Printf("%#v\n",len(result))
    i := 0
    for idx,j := range posUnigrams {
        change := true
        for _,k := range result {
            if (k.Ngram == (tag+" "+j.PosTag)) {
                i += 1
                next[idx] = math.Max(float64(k.Count)*val/float64(j.Count),previous[idx])
                change = false
            }
        }
        if(change) {
            next[idx] = math.Max(unseen_prob,previous[idx])
        }
    }
    //fmt.Printf("%v\n",next)
    return next
}

func fastProcessTag(tag string,tag_count int,val float64,unseen_prob float64,posgrams []Ngrams, posUnigrams []PosUniGram,previous []float64) []float64 {
    next := make([]float64,82)
    result := make([]Ngrams,0)
    for _,j := range posgrams {
        temp := strings.Split(j.Ngram," ")
        if(temp[0] == tag) {
            result = append(result,j)
        }
    }
    // fmt.Printf("%#v\n",len(result))
    i := 0
    for idx,j := range posUnigrams {
        change := true
        for _,k := range result {
            if (k.Ngram == (tag+" "+j.PosTag)) {
                i += 1
                next[idx] = math.Max(float64(k.Count)*val/float64(tag_count),previous[idx])
                change = false
            }
        }
        if(change) {
            next[idx] = math.Max(unseen_prob,previous[idx])
        }
    }
    //fmt.Printf("%v\n",next)
    return next
}

func processEmission(word string,posTag string,unseen_prob float64,chunkUnigrams []Ngrams, previous []float64) []float64 {
    next := make([]float64,len(previous))
    possibleTags := getChunkPosgram(word,posTag)
    if (len(possibleTags) == 0) {
        possibleTags = getChunkPosgram("nil",posTag)
    }
    // fmt.Printf("%s %s\n",word,posTag)
    for idx,j := range chunkUnigrams {
        for _,k := range possibleTags {
            if (k.ChunkTag == j.Ngram) {
                // fmt.Printf("%s %s\n",k.ChunkTag,j.Ngram)
                // next[idx] = math.Log(float64(j.Count)) - math.Log(float64(k.Count)) - previous[idx]
                next[idx] = float64(k.Count)*previous[idx]/float64(j.Count)
            }
        }
    }
    // fmt.Printf("Emission:%v\n",next)
    return next
}

func processTransition(tag string,tag_count int,val float64,unseen_prob float64, chunkUnigrams []Ngrams,previous []float64) []float64 {
    next := make([]float64,len(previous))
    list := make([]string,len(previous))
    for idx,j := range chunkUnigrams {
        list[idx] =  tag+" "+j.Ngram
    }
    // fmt.Printf("%v\n",tag)
    result := getAllNgram(list,"chunkngram")

    i := 0
    // fmt.Printf("%d\n",tag_count)
    for idx,j := range chunkUnigrams {
        change := true
        for _,k := range result {
            if (k.Ngram == (tag+" "+j.Ngram)) {
                i += 1
                //next[idx] = math.Min(math.Log(float64(j.Count)) - math.Log(float64(k.Count)) - math.Log(val),previous[idx])
                next[idx] = math.Max(float64(k.Count)*val/float64(tag_count),previous[idx])
                change = false
            }
        }
        if(change) {
            next[idx] = math.Max(unseen_prob,previous[idx])
            // fmt.Printf("_++++++++++++++++%0.2f\n",next[idx])
        }
    }
    // fmt.Printf("Transition:%v\n",next)
    return next
}

func getPOSTags(sentence string) []string {
    posTags := make([]string,0)
    previous := make([]float64,82)
    next := make([]float64,82)
    regularexp := GetRegex()
    tokens := ProcessSentences(sentence,regularexp)
    tokens = append(tokens,"<\\s>")
    if(len(tokens) > 2) {
        fmt.Printf("%#v\n",tokens)
        values := getAllPosUnigrams()
        tag := "starts"
        //posTags = append(posTags,tag)
        // to-do change this to a dynamic query for total
        // unseen_tag_prob := getPOSUnseen("posTags",1045943)
        start := time.Now()
        unseen_word_prob := getPOSUnseen("wordPosgram",1101375)
        unseen_tag_prob := 0.0
        //unseen_word_prob := 0.0
        wordgram := bulkPosgrams(values)
        for idx,j := range values {
            count := getNgram(tag + " " + j.PosTag,"posTags")
            //previous[idx] = math.Log(float64(count)) - math.Log(float64(j.Count))
            previous[idx] = float64(count)/float64(j.Count)
        }
        // fmt.Printf("%#v\n",previous)
        wordPosgram := bulkPoswordgrams(tokens)
        previous = fastProcessWord(tokens[1],unseen_word_prob,wordPosgram, values,previous)
        posTags = append(posTags,getMaxTag(values,previous))
        for _,i := range tokens[2:len(tokens)-1] {
            for idx,j := range values {
                //start := time.Now()
                next = fastProcessTag(j.PosTag,j.Count,previous[idx],unseen_tag_prob,wordgram,values,next)
                //elapsed := time.Since(start)
                //fmt.Println("Transition:%s\n",elapsed)
            }
            //next = processWord(i,unseen_word_prob, values,next)

            next = fastProcessWord(i,unseen_word_prob,wordPosgram, values,next)
            // elapsed := time.Since(start)
            // fmt.Println("Emission:%s\n",elapsed)
            posTags = append(posTags,getMaxTag(values,next))
            //fmt.Printf("POS:%s\n",printMaxTag(values,next))
            previous = next
            next = make([]float64,82)
        }
        elapsed := time.Since(start)
        fmt.Println(elapsed)
    }
    //posTags = append(posTags,"ends")
    return posTags
}

func getChunkTags(tokens []string,posTags []string) []string {
    chunkTags := make([]string,0)
    previous := make([]float64,22)
    next := make([]float64,22)
    //fmt.Printf("%#v\n",tokens)
    values := getAllChunkUnigrams()
    tag := "start_chunk"

    // to-do change this to a dynamic query for total
    // unseen_tag_prob := getPOSUnseen("posTags",1045943)
    //unseen_word_prob := getPOSUnseen("wordPosgram",1101375)
    unseen_word_prob := 0.0
    unseen_tag_prob := getPOSUnseen("chunkgrams",187939)
    //unseen_word_prob := 0.0
    if(len(tokens) == len(posTags)) {
        for idx,j := range values {
            fmt.Printf(tag + " " + j.Ngram+"\n")
            count := getNgram(tag + " " + j.Ngram,"chunkngram")
            previous[idx] = float64(count)/float64(j.Count)
            //previous[idx] = float64(count)/float64(j.Count)
        }
        // fmt.Printf("Previous:%#v\n",previous)
        // printTags(values)
        previous = processEmission(tokens[0],posTags[0],unseen_word_prob, values,previous)
        chunkTags = append(chunkTags,getMaxChunkTag(values,previous))
        for i,_ := range tokens[1:len(tokens)] {
            for idx,j := range values {
                next = processTransition(j.Ngram,j.Count,previous[idx],unseen_tag_prob,values,next)
            }

            next = processEmission(tokens[i+1],posTags[i+1],unseen_word_prob, values,next)
            chunkTags = append(chunkTags,getMaxChunkTag(values,next))
            // fmt.Printf("CHUNK:%s\n",getMaxChunkTag(values,next))
            previous = next
            next = make([]float64,22)
        }
    }
    return chunkTags
}

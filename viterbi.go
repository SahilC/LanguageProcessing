package main
import (
    "fmt"
    "strings"
    "math"
)
func printMaxTag(posUnigrams []PosUniGram,previous []float64) string {
    maxVal := math.Log(0)
    maxTag := posUnigrams[0].PosTag
    for idx,j := range posUnigrams {
        if(previous[idx] != 0 && math.Exp(previous[idx]) > maxVal) {
            maxVal = math.Exp(previous[idx])
            maxTag = j.PosTag
        }
    }
    return maxTag
}

func processWord(word string,unseen_prob float64,posUnigrams []PosUniGram, previous []float64) []float64 {
    next := make([]float64,82)
    possibleTags := getWordPosgramCount(word)
    for idx,j := range posUnigrams {
        changed := true
        for _,k := range possibleTags {
            if (k.PosTag == j.PosTag) {
                //next[idx] = math.Log(float64(k.Count)) - math.Log(float64(j.Count)) + previous[idx]
                next[idx] = float64(k.Count)*previous[idx]/float64(j.Count)
                changed = false
            }
        }
        if(changed) {
            next[idx] = unseen_prob*previous[idx]
        }
    }
    //fmt.Printf("%v",next)
    return next
}

func getPOSUnseen(database_name string, ngram_length int) float64 {
    frequecy_distribution := getFrequencyAggregation(database_name)
    return float64(frequecy_distribution[1])/float64(ngram_length)
}

func processTag(tag string,val float64,unseen_prob float64, posUnigrams []PosUniGram,previous []float64) []float64 {
    next := make([]float64,82)
    i := 0
    for idx,j := range posUnigrams {
            i += 1
            count := getPosgram( tag+" "+j.PosTag)
            //fmt.Printf("%.2f %.2f\n",math.Log(float64(count)) - math.Log(float64(j.Count)) + val,previous[idx])
            //next[idx] = math.Max(math.Log(float64(count)) - math.Log(float64(j.Count)) + val,previous[idx])
            if (count != 0) {
                next[idx] = math.Max(float64(count)*val/float64(j.Count),previous[idx])
            } else {
                next[idx] = math.Max(unseen_prob,previous[idx])
            }
    }
    //fmt.Printf("%v\n",i)
    return next
}

func fastProcessTag(tag string,val float64,unseen_prob float64, posUnigrams []PosUniGram,previous []float64) []float64 {
    next := make([]float64,82)
    list := make([]string,82)
    for idx,j := range posUnigrams {
        list[idx] =  tag+" "+j.PosTag
    }
    result := getAllPosgram(list)
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

func viterbi(sentence string) []string {
    posTags := make([]string,0)
    previous := make([]float64,82)
    next := make([]float64,82)
    regularexp := GetRegex()
    tokens := ProcessLines(sentence,regularexp)
    tokens = append(tokens,"<\\s>")
    fmt.Printf("%s\n",strings.Join(tokens," "))
    values := getAllPosUnigrams()
    tag := "starts"

    // to-do change this to a dynamic query for total
    unseen_tag_prob := getPOSUnseen("posTags",1045943)
    unseen_word_prob := getPOSUnseen("wordPosgram",1101375)

    for idx,j := range values {
        count := getPosgram(tag + " " + j.PosTag)
        //previous[idx] = math.Log(float64(count)) - math.Log(float64(j.Count))
        previous[idx] = float64(count)/float64(j.Count)
    }
    previous = processWord(tokens[1],unseen_word_prob, values,previous)
    posTags = append(posTags,printMaxTag(values,previous))
    for _,i := range tokens[2:len(tokens)-1] {
        for idx,j := range values {
            next = fastProcessTag(j.PosTag,previous[idx],unseen_tag_prob,values,next)
        }
        next = processWord(i,unseen_word_prob, values,next)
        posTags = append(posTags,printMaxTag(values,next))
        //fmt.Printf("POS:%s\n",printMaxTag(values,next))
        previous = next
        next = make([]float64,82)
    }
    return posTags
}

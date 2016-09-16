package main
import (
    "fmt"
    "math"
    //"strings"
)
func processWord(word string,posUnigrams []PosUniGram, previous []float64) []float64 {
    next := make([]float64,82)
    possibleTags := getWordPosgramCount(word)
    for idx,j := range posUnigrams {
        for _,k := range possibleTags {
            if (k.PosTag == j.PosTag) {
                //fmt.Printf("%s %d\n",k.PosTag,k.Count)
                next[idx] = math.Log(float64(k.Count)) - math.Log(float64(j.Count)) + previous[idx]
            }
        }
        // fmt.Printf("%s %d\n",j.PosTag,j.Count)
    }
    //fmt.Printf("%v",next)
    return next
}

func processTag(tag string,posUnigrams []PosUniGram,val float64,previous []float64) []float64 {
    next := make([]float64,82)
    //possibleTags := getWordPosgramCount(word)
    for idx,j := range posUnigrams {
        //for _,k := range possibleTags {
            count := getPosgram( tag+" "+j.PosTag)
            //fmt.Printf("%s %s:%d\n",j.PosTag,tag,count)
            next[idx] = math.Max(math.Log(float64(count)) - math.Log(float64(j.Count)) + val,previous[idx])
        //}
        // fmt.Printf("%s %d\n",j.PosTag,j.Count)
    }
    //fmt.Printf("%s %v\n",tag,next)
    return next
}

func viterbi(sentence string) {
    previous := make([]float64,82)
    next := make([]float64,82)
    regularexp := GetRegex()
    tokens := ProcessLine(sentence,regularexp)
    tokens = append(tokens,"<\\s>")
    values := getAllPosUnigrams()
    tag := "starts"
    for idx,j := range values {
        count := getPosgram(tag + " " + j.PosTag)
        previous[idx] = math.Log(float64(count)) - math.Log(float64(j.Count))
    }
    // fmt.Printf("%v\n",tokens[0])
    previous = processWord(tokens[1],values,previous)

    for _,i := range tokens[2:len(tokens)-1] {
        fmt.Printf("======================%v\n",previous)
        fmt.Printf("======================%s\n",i)

        for idx,j := range values {
            next = processTag(j.PosTag,values,previous[idx],next)
        }
        next = processWord(i,values,next)

        previous = next
        next = make([]float64,82)
    }
    for idx,_ := range previous {
        fmt.Printf("%s %.2f\n",values[idx].PosTag,previous[idx])
    }
    // fmt.Printf("%v\n",previous)
    // fmt.Printf("%v\n",values)
}

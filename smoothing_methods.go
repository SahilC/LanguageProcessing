package main
import (
        "strings"
        "fmt"
    )
func getUnseenProbability(ngram map[string] int) float64 {
    var frequecy_distribution map[int] int = bucketNgrams(ngram)
    return float64(frequecy_distribution[1])/float64(len(ngram))
}

func estimateSentenceProbability(ngram map[string] int) {
    var tokens map[string] int = buildNGram(TokenizeSentences("Test/Test.txt"),3)
    var lessergram map[string] int = buildNGram(TokenizeSentences("Corpora/pg105.txt"),2)
    var smooth_ngram map[string] float64 = GoodTuring(ngram)
    var unseen float64 = getUnseenProbability(ngram)
    var probability = 1.0
    for k,_ := range tokens {
        if(smooth_ngram[k] == 0.0) {
            probability *= unseen
            //fmt.Printf("%s:%0.2f----------\n",k,unseen)
        } else {
            temp := strings.Split(k," ")
            //fmt.Printf(strings.Join(temp[1:len(temp)]," ")+"\n")
            probability *= smooth_ngram[k]/float64(lessergram[strings.Join(temp[0:len(temp)-1]," ")])
            //fmt.Printf("%s:%0.2f**********\n",k,smooth_ngram[k]/float64(lessergram[strings.Join(temp[0:len(temp)-1]," ")]))
        }
    }
    fmt.Printf("probability of sentence:%0.30f",probability)
}

func bucketNgrams(ngram map[string] int) map[int] int {
    var frequecy_distribution = make(map[int] int,0)
    for _,v := range ngram {
        frequecy_distribution[v] += 1
    }
    return frequecy_distribution
}

func GoodTuring(ngram map[string] int) map[string] float64 {
    var frequecy_distribution map[int] int = bucketNgrams(ngram)
    var smooth_ngram = make(map[string] float64)
    i := 0
    for k,v := range ngram {
        if(i <= 5) {
            smooth_ngram[k] = float64(v+1)*(float64(frequecy_distribution[v+1])/float64(frequecy_distribution[v]))
        } else {
            smooth_ngram[k] = float64(v)
        }
        i += 1
        // fmt.Printf("%s:%0.2f\n",k,smooth_ngram[k])
    }
    return smooth_ngram
    // fmt.Printf("%d %0.2f\n",ngram["determined to go"],smooth_ngram["determined to go"])
    // fmt.Printf("%d %0.2f\n",ngram["heard a soft"],smooth_ngram["heard a soft"])
    // fmt.Printf("%d %0.2f\n",ngram["the bottom ."],smooth_ngram["the bottom ."])
    // fmt.Printf("%d %0.2f\n",ngram["it must be"],smooth_ngram["it must be"])
}

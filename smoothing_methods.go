package main
// import "fmt"

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
        if(i <= 10) {
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

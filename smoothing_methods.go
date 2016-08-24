package main
import "fmt"
func bucketNgrams(ngram map[string] int) map[int] int {
    var frequecy_distribution = make(map[int] int,0)
    for _,v := range ngram {
        frequecy_distribution[v] += 1
    }
    return frequecy_distribution
}

func GoodTuring(ngram map[string] int) {
    var frequecy_distribution map[int] int = bucketNgrams(ngram)
    for k,v := range ngram {
        ngram[k] = (v+1)*(float64(frequecy_distribution[v+1])/frequecy_distribution[v])
    }
}

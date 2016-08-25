package main
// import (
// 				"bufio"
// 				"os"
// 				"fmt"
// )

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Printf("Enter Sentence to be Evaluated:\n")
	//text, _ := reader.ReadString('\n')
	filename := "Corpora/pg105.txt"
	var tokens map[string] int = buildNGram(TokenizeSentences(filename),6)
	// for k,v := range tokens {
	// 	fmt.Printf("%d:%s\n",v,k);
	// }
	//fmt.Printf(text)
	var ngram map[string] float64 = GoodTuring(tokens)
	var sorted = sortFloatNGrams(ngram)
	//var sorted = sortNGrams(tokens)
	//drawHistogram("histogram",sorted)
	PlotLogLog(sorted,"Austen")
}

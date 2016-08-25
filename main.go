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
	var tokens map[string] int = buildNGram(TokenizeSentences(),3)
	// for k,v := range tokens {
	// 	fmt.Printf("%d:%s\n",v,k);
	// }
	//fmt.Printf(text)
	var ngram map[string] float64 = GoodTuring(tokens)
	var sorted = sortFloatNGrams(ngram)
	//drawHistogram("histogram",sorted)
	PlotLogLog(sorted)
}

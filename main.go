package main
import (
		"io/ioutil"
		//"fmt"
)
// import (
// 				"bufio"
// 				"os"
// 				"fmt"
// )
func WalkCorpora() {
    files, _ := ioutil.ReadDir("Corpora/")
	vals := make([][]float64 ,0)
	filenames := make([]string,0)
    for _, f := range files {
			filenames = append(filenames,f.Name())
			filename := "Corpora/"+f.Name()
			var tokens map[string] int = buildNGram(TokenizeSentences(filename),6)
			// for k,v := range tokens {
			// 	fmt.Printf("%d:%s\n",v,k);
			// }
			//fmt.Printf(text)
			var ngram map[string] float64 = GoodTuring(tokens)
			var sorted = sortFloatNGrams(ngram)
			vals = append(vals,sorted)
			//var sorted = sortNGrams(tokens)
			//drawHistogram("histogram",sorted)
    }
	PlotLogLog(vals,"Multiplot",filenames)
}

func main() {
	WalkCorpora()
	// //reader := bufio.NewReader(os.Stdin)
	// //fmt.Printf("Enter Sentence to be Evaluated:\n")
	// //text, _ := reader.ReadString('\n')
	// filename := "Corpora/pg105.txt"
	// var tokens map[string] int = buildNGram(TokenizeSentences(filename),6)
	// // for k,v := range tokens {
	// // 	fmt.Printf("%d:%s\n",v,k);
	// // }
	// //fmt.Printf(text)
	// var ngram map[string] float64 = GoodTuring(tokens)
	// var sorted = sortFloatNGrams(ngram)
	// //var sorted = sortNGrams(tokens)
	// //drawHistogram("histogram",sorted)
	// PlotLogLog(sorted,"Austen")
}

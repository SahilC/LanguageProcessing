package main
import (
		"io/ioutil"
		"fmt"
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
			fmt.Printf(f.Name()+"\n");
			filenames = append(filenames,f.Name())
			filename := "Corpora/"+f.Name()
			var tokens map[string] int = buildNGram(TokenizeSentences(filename),3)
			// for k,v := range tokens {
			// 	fmt.Printf("%d:%s\n",v,k);
			// }
			//fmt.Printf(text)
			var ngram map[string] float64 = GoodTuring(tokens)
			var sorted = sortFloatNGrams(ngram)
			drawHistogram(f.Name(),sorted)
			vals = append(vals,sorted)
			//var sorted = sortNGrams(tokens)
    }
	PlotLogLog(vals,"Multiplot",filenames)
}

func main() {
	//ReadBrown()
	//runPOSTests()
	RunHMMChunkerTests()
	// ReadCONLL()
	// fmt.Printf("%#v",generateLMSentence())
	// randomChunkWalk()
	//getFrequencyAggregation("wordPosgram")
	//fmt.Printf("%#v",viterbi("My name is Sahil"))
	//WalkCorpora()
	// //reader := bufio.NewReader(os.Stdin)
	// //fmt.Printf("Enter Sentence to be Evaluated:\n")
	// //text, _ := reader.ReadString('\n')
	//  filename := "Corpora/pg105.txt"
	//  var tokens map[string] int = buildNGram(TokenizeSentences(filename),3)
	//  estimateSentenceProbability(tokens)
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

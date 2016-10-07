package main

import (
		"fmt"
		"strings"
		"io/ioutil"
		"sort"
		"regexp"
		"strconv"
		"github.com/gonum/plot"
		"github.com/gonum/plot/vg"
		"github.com/gonum/plot/plotter"
		"github.com/gonum/plot/plotutil"
)

type Ngram struct {
    ngram string
    count int
}

/*
	unigramHistogram:-
		Desciption- Reads the corpus,and tokenizes it.
		From the tokens generates a unigram model, which is then sorted
		for rank, and then ploted to chart which is saved in histogram.png.
		Run this function for solution to b. of assignment 2.
*/
func unigramHistogram() {
	var tokens [][]string = ReadFile()
	var ngram = buildNGram(tokens,1)
	var sorted = sortNGrams(ngram)
	drawHistogram("histogram",sorted)
}

/*
	randomWalk:-
		Input- A list of lists which contain tokenized strings
		Output- A sequence of random sentences generated from the different N-grams(1-6)
		Desciption- This function performs a random walk on different N-gram models
		to generate sentences. It first interpolates the current N-gram for the <s> tag.
		Once found, it tries to find the word which is most likely to follow that sequence of characters.
		By iteratively querying the model, we generate a sequence of tokens which have been seen together.
		Solution for the d. part of assignment 2 is given here.
*/

func randomWalk(tokens [][]string) {
	sortedNgrams:=[]Ngram{}
	var ngrams []map[string] int
	var ngram map[string] int
	var count int = 0
	for n:=1;n<=6;n++ {
		ngram = buildNGram(tokens,n)
		ngrams = append(ngrams,ngram)
		sortedNgrams = sortStructNgram(ngram)
		if(n > 1) {
			for _,k := range sortedNgrams {
				var temp = strings.Split(k.ngram," ")
				if(strings.EqualFold(temp[0],"<s>" )) {
					var currentToken string = ""
					var sentence string = strings.Join(temp[:len(temp)]," ")
					for !strings.Contains(sentence,"<\\s>") {
						for _,a := range sortedNgrams {
							var temp2 []string = strings.Split(a.ngram, " ")
							//fmt.Printf(sentence+"\n")
							if(strings.EqualFold(strings.Join(temp[1:]," "),strings.Join(temp2[:len(temp2)-1]," ")) && !strings.Contains(sentence,a.ngram)) {
								currentToken = temp2[len(temp2)-1]
								sentence = sentence+ " " + currentToken
								temp = temp2
								//fmt.Printf(a.ngram+"%d\n",a.count)
								break
							}
						}
					}
					count += 1
					fmt.Printf("%d. ",n)
					fmt.Printf(sentence+"\n")
					if(count > 5) {
						count = 0
						break
					}
				}
			}
		}
	}
}

// func randomWalk(tokens [][]string) {
// 	var ngrams []map[string] int
// 	var count int = 0
// 	for n:=1;n<=6;n++ {
// 		ngrams = append(ngrams,buildNGram(tokens,n))
// 		if(n > 1) {
// 			for k,_ := range ngrams[n-1] {
// 				var temp = strings.Split(k," ")
// 				if(strings.EqualFold(temp[0],"<s>" )) {
// 					var currentToken string = ""
// 					var sentence string = strings.Join(temp[:len(temp)-1]," ")
// 					for !strings.EqualFold(currentToken,"<\\s>") {
// 						for a,_ := range ngrams[n-1] {
// 							var temp2 []string = strings.Split(a, " ")
// 							if(strings.EqualFold(strings.Join(temp[1:]," "),strings.Join(temp2[:len(temp2)-1]," "))) {
// 								currentToken = temp2[len(temp2)-1]
// 								sentence = sentence+ " " + currentToken
// 								temp = temp2
// 								//fmt.Printf(sentence+"\n")
// 								break
// 							}
// 						}
// 					}
// 					count += 1
// 					fmt.Printf("%d. ",n)
// 					fmt.Printf(sentence+"\n")
// 					if(count > 5) {
// 						count = 0
// 						break
// 					}
// 				}
// 			}
// 		}
// 	}
// }

/*
	estimateEndProbabilities :-
		Input- A list of lists which contain tokenized strings
		Output- A set of N-grams with the count of number of times that N-gram has occured
		with the <\s> following it, to the count of the number of times that N-gram has occured
		That is the N-gram count for (X,<\s>) to the (N-1)-gram count of X.
		Desciption- This code iterates over all N-grams(1-6) to see which ones contain the X<\s>
		token, then if they do, we display the count of that N-gram and query the N-1-gram for
		count of X. This is the solution to a. from the Assignment
*/
func estimateEndProbabilities(tokens [][]string) {
	//fmt.Printf("%v",ngram)
	//fmt.Printf("%#v",tokens)
	var ngrams []map[string] int
	for n := 1; n<= 6; n++ {
		var count int = 0
		ngrams = append(ngrams,buildNGram(tokens,n))
		for k,v := range ngrams[n-1] {
			var temp = strings.Split(k," ")
			if(strings.EqualFold(temp[len(temp)-1],"<\\s>" )) {
				count += 1
				if(n > 1) {
					var gramCount = ngrams[n-2][strings.Join(temp[:len(temp)-1]," ")]
					fmt.Printf("%s %d:%d\n",strings.Join(temp[:len(temp)-1]," "),v,gramCount)
				}
			}
		}
		fmt.Printf("Count: %d",count);
	}
}

/*
	plotEndFrequencies -
	Output - Figures generated with the corrosponding N-grams are numbered as histogramN.png
	Description - This function reads the tokens from the file, and for all N-grams (2-6),
	generates a plot for all elements that end with <\s>. The code searches through the N-grams
	to see which end with <\s>, and only considers those for the rank-vs-freq plot. Again, we have to
	sort the values to get an efficent plot. This is the answer to the c. part of the assignment 2.
*/
func plotEndFrequencies() {
	var tokens [][]string = ReadFile()
	var ngram map[string] int
	var temp map[string] int
	var sorted []float64
	for i := 2;i<=6;i ++ {
		ngram = buildNGram(tokens,i)
		temp = make(map[string]int)
		for k,v:= range ngram {
			var t1 = strings.Split(k," ")
			if(strings.EqualFold(t1[len(t1)-1],"<\\s>" )) {
				temp[k] = v
			}
		}
		sorted = sortNGrams(temp)
		drawHistogram("histogram"+strconv.Itoa(i),sorted)
	}
}

/*
	TokenizeSentences :-
	Description -
		Uses smart Heuristics to split sentences. Replaces "." which follows abbriviations with <p>, and reorders
		sentences within quotes to be more flexible to splitting. Splits on <end> symbol, to give new sentences.
*/
func TokenizeSentences(filename string) [][]string {
	dat, err := ioutil.ReadFile(filename)
    if err != nil {
        panic(err)
    }
    var abbriviations = regexp.MustCompile(`(Mr|St|Mrs|Ms|Dr|Inc|Ltd|Jr|Sr|Co)[.]`)
    var  str string = strings.Replace(string(dat),"\n"," ",-1)
    str = strings.Replace(str,"\r"," ",-1)
    str = abbriviations.ReplaceAllString(str,"$1<p>")
    str = strings.Replace(str,".\"","\".",-1)
    str = strings.Replace(str,"!\"","\"!",-1)
    str = strings.Replace(str,"?\"","\"?",-1)
    str = strings.Replace(str,"?","?<end>",-1)
    str = strings.Replace(str,"!","!<end>",-1)
    str = strings.Replace(str,".",".<end>",-1)
    str = strings.Replace(str,"<p>",".",-1)
    var s [] string = strings.Split(str,"<end>")
    var tokens = make([][]string,0)
    var matches = make([]string, 0, 2000)
    var regularexp = GetRegex()
    for _,i := range s {
    	matches = ProcessPOSLine(i,regularexp)
    	matches = append(matches,"<\\s>")
    	tokens = append(tokens,matches)
    }
    return tokens
}

func drawHistogram(fileName string, values []float64) {
    groupA := make(plotter.Values,0)
    //vals := make([]string,0)
    for _,i := range values {
    	groupA = append(groupA,i)
    }
    p, err := plot.New()
    if err != nil {
        panic(err)
    }
    p.Title.Text = "Bar chart"
    p.Y.Label.Text = "Heights"

    w := vg.Points(2)

    barsA, err := plotter.NewBarChart(groupA, w)
    if err != nil {
        panic(err)
    }
    barsA.LineStyle.Width = vg.Length(1)
    barsA.Color = plotutil.Color(1)
    barsA.Offset = -w

    p.Add(barsA)
    //p.NominalX(vals)

    if err := p.Save(10*vg.Inch, 5*vg.Inch, fileName+".png"); err != nil {
        panic(err)
    }
}
/*
	sortNGrams - Returns a list of values sorted in terms of frequencies
*/

func sortFloatNGrams(ngram map[string] float64) []float64 {
	//groupA := make([]float64,0)
    n := map[float64][]string{}
    var a []float64
    for k, v := range ngram {
        n[v] = append(n[v], k)
    }
    for k := range n {
        a = append(a, k)
    }
    sort.Sort(sort.Reverse(sort.Float64Slice(a)))
	// for _, k := range a {
	// 	// for _,s := range n[k] {
	// 	// 	fmt.Printf("%s %d\n",s,k)
	// 	// }
    //     groupA = append(groupA,float64(k))
    // }
    return a
}

func sortNGrams(ngram map[string] int) []float64 {
	groupA := make([]float64,0)
    n := map[int][]string{}
    var a []int
    for k, v := range ngram {
        n[v] = append(n[v], k)
    }
    for k := range n {
        a = append(a, k)
    }
    sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for _, k := range a {
		// for _,s := range n[k] {
		// 	fmt.Printf("%s %d\n",s,k)
		// }
        groupA = append(groupA,float64(k))
    }
    return groupA
}
/*
	buildNgram - Iterates over tokens to produce different N-grams. Returns a map of String to Int representing
	N-gram to count pairs.
*/
func buildNGram(tokens [][]string, n int) map[string] int {
	var ngram = make(map[string] int)
	for k := range tokens {
    	for j := 0; j <= (len(tokens[k]) - n); j++ {
    		ngram[strings.Join(tokens[k][j:j+n]," ")] += 1
    	}
    }
    return ngram
}

func sortStructNgram(ngram map[string] int) []Ngram{
    n := map[int][]string{}
    var a []int
    //fmt.Printf("Entering Sort\n")
    sortedNgrams:=[]Ngram{}
    for k, v := range ngram {
        n[v] = append(n[v], k)
    }
    for k := range n {
        a = append(a, k)
    }
    sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for _, k := range a {
		for _,s := range n[k] {
			var temp Ngram
			temp.ngram = s
			temp.count = k
			sortedNgrams = append(sortedNgrams,temp)
			//fmt.Printf("%s %d\n",s,k)
		}
		//fmt.Printf("Done")
    }
    //fmt.Printf("Out")
    return sortedNgrams
}
// func main() {
// 	//var tokens = TokenizeSentences()
// 	//var ngram = buildNGram(tokens,1)
// 	//var sorted = sortNGrams(ngram)
// 	//drawHistogram("histogram",sorted)
// 	//estimateEndProbabilities(tokens)
// 	//randomWalk(tokens)
// 	//plotEndFrequencies()
// }

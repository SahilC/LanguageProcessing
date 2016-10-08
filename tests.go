package main

import (
    "fmt"
    "strings"
    "io/ioutil"
)

func runPOSTests() {
    dat, err := ioutil.ReadFile("Test/testSet.txt")
    corpus_location := "/home/sahil/nltk_data/corpora/brown"
    if err != nil {
        panic(err)
    }
    var regularexp = GetRegex()
    //fmt.Printf(string(dat))
    var s [] string = strings.Split(string(dat),"\n")
    for _,i := range s {
        dat,_ := ioutil.ReadFile(corpus_location+"/"+i)
        //fmt.Printf(string(dat))
        s := strings.Split(string(dat),"\n")
        matches := make([]string, 0, 2000)
        for _,i := range s {
        	matches = ProcessPOSLine(i,regularexp)
            if(len(matches) > 1) {
                matches = append(matches,"<\\s>/ends")
                sentence := make([]string,0)
                posTags := make([]string,0)
                for _,k := range matches {
                    temp := strings.Split(k,"/")
                    sentence = append(sentence,temp[0])
                    posTags = append(posTags,temp[len(temp)-1])
                }
                //fmt.Printf("%#v\n",i)
                returnTags := getPOSTags(strings.Join(sentence[1:len(sentence)-1]," "))
                fmt.Printf("%s\n%#v\n%#v\n=====================\n",strings.Join(sentence," "),returnTags,posTags)
            }
        }
    }
}


func RunHMMChunkerTests() {
    corpus_location := "/home/sahil/nltk_data/corpora/conll2000/test.txt"
    dat,_ := ioutil.ReadFile(corpus_location)
    //fmt.Printf(string(dat))
    s := strings.Split(string(dat),"\n")
    matches := make([]string, 0)
    for _,i := range s {
        line := strings.Split(string(i)," ")
        //fmt.Printf("%#v\n",line)
        if(len(line) == 3) {
            if ( line[0] != "." && line[1] != "." && line[2] != "O") {
                matches = append(matches,line[0])
            } else {
                returnTags := getPOSTags(strings.Join(matches," "))
                chunkTags := getChunkTags(matches,returnTags)
                fmt.Printf("%#v\n",chunkTags)
                matches = make([]string, 0)
            }
        }
    }
}

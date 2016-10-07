package main

import ("io/ioutil"
        "strings"
        "regexp"
        "html"
        "fmt"
)
/*
	GetRegex :-
	 	Builds a Regex string to handle various scenarios like unicodeEmojis, ASCII smileys, contactions like 's
	 	urls,emails, phone numbers, mentions and hashtags
*/
func GetRegex() string {
	regexps :=  []string {
         "[[:alnum:]]+\\/([[:alpha:]]+[[:punct:]]*[[:alpha:]]*[[:punct:]]*)", //POSTags
		 "[\\x{2712}\\x{2714}\\x{2716}\\x{271d}\\x{2721}\\x{2728}\\x{2733}\\x{2734}\\x{2744}\\x{2747}\\x{274c}\\x{274e}\\x{2753}-\\x{2755}\\x{2757}\\x{2763}\\x{2764}\\x{2795}-\\x{2797}\\x{27a1}\\x{27b0}\\x{27bf}\\x{2934}\\x{2935}\\x{2b05}-\\x{2b07}\\x{2b1b}\\x{2b1c}\\x{2b50}\\x{2b55}\\x{3030}\\x{303d}\\x{1f004}\\x{1f0cf}\\x{1f170}\\x{1f171}\\x{1f17e}\\x{1f17f}\\x{1f18e}\\x{1f191}-\\x{1f19a}\\x{1f201}\\x{1f202}\\x{1f21a}\\x{1f22f}\\x{1f232}-\\x{1f23a}\\x{1f250}\\x{1f251}\\x{1f300}-\\x{1f321}\\x{1f324}-\\x{1f393}\\x{1f396}\\x{1f397}\\x{1f399}-\\x{1f39b}\\x{1f39e}-\\x{1f3f0}\\x{1f3f3}-\\x{1f3f5}\\x{1f3f7}-\\x{1f4fd}\\x{1f4ff}-\\x{1f53d}\\x{1f549}-\\x{1f54e}\\x{1f550}-\\x{1f567}\\x{1f56f}\\x{1f570}\\x{1f573}-\\x{1f579}\\x{1f587}\\x{1f58a}-\\x{1f58d}\\x{1f590}\\x{1f595}\\x{1f596}\\x{1f5a5}\\x{1f5a8}\\x{1f5b1}\\x{1f5b2}\\x{1f5bc}\\x{1f5c2}-\\x{1f5c4}\\x{1f5d1}-\\x{1f5d3}\\x{1f5dc}-\\x{1f5de}\\x{1f5e1}\\x{1f5e3}\\x{1f5ef}\\x{1f5f3}\\x{1f5fa}-\\x{1f64f}\\x{1f680}-\\x{1f6c5}\\x{1f6cb}-\\x{1f6d0}\\x{1f6e0}-\\x{1f6e5}\\x{1f6e9}\\x{1f6eb}\\x{1f6ec}\\x{1f6f0}\\x{1f6f3}\\x{1f910}-\\x{1f918}\\x{1f980}-\\x{1f984}\\x{1f9c0}\\x{3297}\\x{3299}\\x{a9}\\x{ae}\\x{203c}\\x{2049}\\x{2122}\\x{2139}\\x{2194}-\\x{2199}\\x{21a9}\\x{21aa}\\x{231a}\\x{231b}\\x{2328}\\x{2388}\\x{23cf}\\x{23e9}-\\x{23f3}\\x{23f8}-\\x{23fa}\\x{24c2}\\x{25aa}\\x{25ab}\\x{25b6}\\x{25c0}\\x{25fb}-\\x{25fe}\\x{2600}-\\x{2604}\\x{260e}\\x{2611}\\x{2614}\\x{2615}\\x{2618}\\x{261d}\\x{2620}\\x{2622}\\x{2623}\\x{2626}\\x{262a}\\x{262e}\\x{262f}\\x{2638}-\\x{263a}\\x{2648}-\\x{2653}\\x{2660}\\x{2663}\\x{2665}\\x{2666}\\x{2668}\\x{267b}\\x{267f}\\x{2692}-\\x{2694}\\x{2696}\\x{2697}\\x{2699}\\x{269b}\\x{269c}\\x{26a0}\\x{26a1}\\x{26aa}\\x{26ab}\\x{26b0}\\x{26b1}\\x{26bd}\\x{26be}\\x{26c4}\\x{26c5}\\x{26c8}\\x{26ce}\\x{26cf}\\x{26d1}\\x{26d3}\\x{26d4}\\x{26e9}\\x{26ea}\\x{26f0}-\\x{26f5}\\x{26f7}-\\x{26fa}\\x{26fd}\\x{2702}\\x{2705}\\x{2708}-\\x{270d}\\x{270f}]|\\x{23}\\x{20e3}|\\x{2a}\\x{20e3}|\\x{30}\\x{20e3}|\\x{31}\\x{20e3}|\\x{32}\\x{20e3}|\\x{33}\\x{20e3}|\\x{34}\\x{20e3}|\\x{35}\\x{20e3}|\\x{36}\\x{20e3}|\\x{37}\\x{20e3}|\\x{38}\\x{20e3}|\\x{39}\\x{20e3}|\\x{1f1e6}[\\x{1f1e8}-\\x{1f1ec}\\x{1f1ee}\\x{1f1f1}\\x{1f1f2}\\x{1f1f4}\\x{1f1f6}-\\x{1f1fa}\\x{1f1fc}\\x{1f1fd}\\x{1f1ff}]|\\x{1f1e7}[\\x{1f1e6}\\x{1f1e7}\\x{1f1e9}-\\x{1f1ef}\\x{1f1f1}-\\x{1f1f4}\\x{1f1f6}-\\x{1f1f9}\\x{1f1fb}\\x{1f1fc}\\x{1f1fe}\\x{1f1ff}]|\\x{1f1e8}[\\x{1f1e6}\\x{1f1e8}\\x{1f1e9}\\x{1f1eb}-\\x{1f1ee}\\x{1f1f0}-\\x{1f1f5}\\x{1f1f7}\\x{1f1fa}-\\x{1f1ff}]|\\x{1f1e9}[\\x{1f1ea}\\x{1f1ec}\\x{1f1ef}\\x{1f1f0}\\x{1f1f2}\\x{1f1f4}\\x{1f1ff}]|\\x{1f1ea}[\\x{1f1e6}\\x{1f1e8}\\x{1f1ea}\\x{1f1ec}\\x{1f1ed}\\x{1f1f7}-\\x{1f1fa}]|\\x{1f1eb}[\\x{1f1ee}-\\x{1f1f0}\\x{1f1f2}\\x{1f1f4}\\x{1f1f7}]|\\x{1f1ec}[\\x{1f1e6}\\x{1f1e7}\\x{1f1e9}-\\x{1f1ee}\\x{1f1f1}-\\x{1f1f3}\\x{1f1f5}-\\x{1f1fa}\\x{1f1fc}\\x{1f1fe}]|\\x{1f1ed}[\\x{1f1f0}\\x{1f1f2}\\x{1f1f3}\\x{1f1f7}\\x{1f1f9}\\x{1f1fa}]|\\x{1f1ee}[\\x{1f1e8}-\\x{1f1ea}\\x{1f1f1}-\\x{1f1f4}\\x{1f1f6}-\\x{1f1f9}]|\\x{1f1ef}[\\x{1f1ea}\\x{1f1f2}\\x{1f1f4}\\x{1f1f5}]|\\x{1f1f0}[\\x{1f1ea}\\x{1f1ec}-\\x{1f1ee}\\x{1f1f2}\\x{1f1f3}\\x{1f1f5}\\x{1f1f7}\\x{1f1fc}\\x{1f1fe}\\x{1f1ff}]|\\x{1f1f1}[\\x{1f1e6}-\\x{1f1e8}\\x{1f1ee}\\x{1f1f0}\\x{1f1f7}-\\x{1f1fb}\\x{1f1fe}]|\\x{1f1f2}[\\x{1f1e6}\\x{1f1e8}-\\x{1f1ed}\\x{1f1f0}-\\x{1f1ff}]|\\x{1f1f3}[\\x{1f1e6}\\x{1f1e8}\\x{1f1ea}-\\x{1f1ec}\\x{1f1ee}\\x{1f1f1}\\x{1f1f4}\\x{1f1f5}\\x{1f1f7}\\x{1f1fa}\\x{1f1ff}]|\\x{1f1f4}\\x{1f1f2}|\\x{1f1f5}[\\x{1f1e6}\\x{1f1ea}-\\x{1f1ed}\\x{1f1f0}-\\x{1f1f3}\\x{1f1f7}-\\x{1f1f9}\\x{1f1fc}\\x{1f1fe}]|\\x{1f1f6}\\x{1f1e6}|\\x{1f1f7}[\\x{1f1ea}\\x{1f1f4}\\x{1f1f8}\\x{1f1fa}\\x{1f1fc}]|\\x{1f1f8}[\\x{1f1e6}-\\x{1f1ea}\\x{1f1ec}-\\x{1f1f4}\\x{1f1f7}-\\x{1f1f9}\\x{1f1fb}\\x{1f1fd}-\\x{1f1ff}]|\\x{1f1f9}[\\x{1f1e6}\\x{1f1e8}\\x{1f1e9}\\x{1f1eb}-\\x{1f1ed}\\x{1f1ef}-\\x{1f1f4}\\x{1f1f7}\\x{1f1f9}\\x{1f1fb}\\x{1f1fc}\\x{1f1ff}]|\\x{1f1fa}[\\x{1f1e6}\\x{1f1ec}\\x{1f1f2}\\x{1f1f8}\\x{1f1fe}\\x{1f1ff}]|\\x{1f1fb}[\\x{1f1e6}\\x{1f1e8}\\x{1f1ea}\\x{1f1ec}\\x{1f1ee}\\x{1f1f3}\\x{1f1fa}]|\\x{1f1fc}[\\x{1f1eb}\\x{1f1f8}]|\\x{1f1fd}\\x{1f1f0}|\\x{1f1fe}[\\x{1f1ea}\\x{1f1f9}]|\\x{1f1ff}[\\x{1f1e6}\\x{1f1f2}\\x{1f1fc}]", //"unicodeEmoji"
		 "(https?:\\/\\/[[:word:]]+\\.[[:word:]]+\\/[[:graph:]]+)", //"urls"
		 "([a-zA-Z0-9_\\.]+\\@[a-zA-Z0-9_]+?\\.[a-zA-Z]{2,3})", //"email"
		 "(@[[:word:]]+)", // "mentions"
		 "(#[[:word:]]+[[:graph:]]*[[:word:]]+)", //"hashtags"
		 "([A-Za-z][A-Za-z'\\-_]+[A-Za-z])", //"smileys"
		 "('s)|(n't)", //"contraction_suffixes"
		 "([[:alpha:]]\\.)|(([[:alpha:]]\\.)+)|([A-Z][bcdfghjklmnpqrstvwxz]\\.)", //"abbriviations"
		 //"(\\$?-?0*(?:[[:digit:]]+(\\?\\!\\,)(?:\\.[[:digit:]]{1,2})?|(?:[[:digit:]]{1,3}(?:,[[:digit:]]{3})*(?:\\.[[:digit:]]{1,2})?)))",
		 //"[[:punct:]]+", //"punchuation"
		 "([[:word:]]+)", //"words"
		 //"numbers":"(%[0-9a-f][0-9a-f])",
	}
	var vals = make([]string, 0)
	for _,m := range regexps {
		vals = append(vals,m)
	}
	var reglist = strings.Join(vals, "|")
	//fmt.Printf(reglist)
	return reglist
}


/*
	ProcessPOSLine :-
	 	Takes the generated Regex expression, and matches tokens line by line to
	 	them. Each match is considered as a separate token.
*/
func ProcessPOSLine(line string, regularexp string) []string {
	//fmt.Printf(line)
	var matches = make([]string, 0, 2000)
	matches = append(matches,"<s>/starts")
	var search = regexp.MustCompile(regularexp)
    for _, m := range search.FindAllString(line, -1) {
        matches = append(matches, string(m))
        //fmt.Printf(string(m)+",")
    }
	//fmt.Printf("\n")
	//fmt.Printf("%#v",matches)
	return matches
}


func ProcessSentences(line string, regularexp string) []string {
	//fmt.Printf(line)
	var matches = make([]string, 0, 2000)
	matches = append(matches,"<s>")
	var search = regexp.MustCompile(regularexp)
    for _, m := range search.FindAllString(line, -1) {
        matches = append(matches, string(m))
        //fmt.Printf(string(m)+",")
    }
	//fmt.Printf("\n")
	//fmt.Printf("%#v",matches)
	return matches
}
/*
	ReadFile :- Reads from the Tweet.en.txt file, splits on each line since
	each tweet can be considered as a set of sentences and sentences seldom span multiple
	lines due to twitter's 140 character designation. It also addes the <s> and </s> tags.
*/
func ReadFile() [][]string {
	dat, err := ioutil.ReadFile("Corpora/tweets.en.txt")
    if err != nil {
        panic(err)
    }
    var regularexp = GetRegex()
    //fmt.Printf(string(dat))
    var s [] string = strings.Split(string(dat),"\n")
    var tokens = make([][]string,0)
    var matches = make([]string, 0, 2000)
    for _,i := range s {
    	matches = ProcessPOSLine(html.UnescapeString(i),regularexp)
    	matches = append(matches,"<\\s>")
    	tokens = append(tokens,matches)
    }
    return tokens
}

func ReadBrown() {
    corpus_location := "/home/sahil/nltk_data/corpora/brown"
    files, _ := ioutil.ReadDir(corpus_location)
    //tokens := make([][]string,0)
    for j, f := range files {
        if(f.Name() != "README" && f.Name() != "CONTENTS" && j < 400) {
            dat,_ := ioutil.ReadFile(corpus_location+"/"+f.Name())
            regularexp := GetRegex()
            //fmt.Printf(string(dat))
            s := strings.Split(string(dat),"\n")
            matches := make([]string, 0, 2000)
            for _,i := range s {
            	matches = ProcessPOSLine(html.UnescapeString(i),regularexp)
                if(len(matches) > 1) {
                    matches = append(matches,"<\\s>/ends")
                    //InsertTokens(matches)
                    InsertPOSNgram(matches,2)
                    //InsertWordPosgram(matches)
                    // fmt.Println(len(matches))
                    // fmt.Println("%v",matches)
                }
            	//tokens = append(tokens,matches)
            }
        } else {
            fmt.Printf("%s\n",f.Name())
        }
    }
}

func ReadCONLL() {
    corpus_location := "/home/sahil/nltk_data/corpora/conll2000/train.txt"
    dat,_ := ioutil.ReadFile(corpus_location)
    //fmt.Printf(string(dat))
    s := strings.Split(string(dat),"\n")
    //matches := make([]string, 0, 2000)
    for _,i := range s {
        line := strings.Split(string(i)," ")
        fmt.Printf("%#v\n",line)
        if(len(line) == 3 && line[0] != "." && line[1] != "." && line[2] != "O") {
            InsertChunkgram(line)
        }
    }
}

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
                returnTags := viterbi(strings.Join(sentence[1:len(sentence)-1]," "))
                fmt.Printf("%s\n%#v\n%#v\n=====================\n",strings.Join(sentence," "),returnTags,posTags)
            }
        }
    }
}

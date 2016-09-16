package main

import (
	"fmt"
    "strings"
    "time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Tokens struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Tokens      []string
	//count     int
	Timestamp time.Time
}

type Ngrams struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Ngram      string
	count     int
	Timestamp time.Time
}

func InsertTokens(tokens []string) {
    maxWait := time.Duration(5 * time.Second)
    session, err := mgo.DialWithTimeout("127.0.0.1",maxWait)
	if err != nil {
		panic(err)
	}
    defer session.Close()
	session.SetMode(mgo.Monotonic, true)
    c := session.DB("nlprokz").C("brownCorpus")
    err = c.Insert(&Tokens{Tokens: tokens, Timestamp: time.Now()})

	if err != nil {
		panic(err)
	}
}

func InsertWordPosgram(tokens []string) {
    maxWait := time.Duration(5 * time.Second)
    session, err := mgo.DialWithTimeout("127.0.0.1",maxWait)
	if err != nil {
		panic(err)
	}
    defer session.Close()
	session.SetMode(mgo.Monotonic, true)
    pos := session.DB("nlprokz").C("wordPosgram")
    //pos := session.DB("nlprokz").C("posUnigram")
	for j := 0; j < len(tokens); j++ {

        var temp = strings.Split(tokens[j],"/")
        if(len(temp) > 1) {
            posTag := temp[len(temp)-1]
            if(strings.Index(posTag,"+") != -1 && strings.Index(posTag,"+") != (len(posTag)-1)  ) {
                temp = strings.Split(posTag,"+")
                if(len(temp[1]) > 1 || temp[1] != "*") {
                    pos.Upsert(bson.M{"posTag":temp[1]},bson.M{"$inc": bson.M{"count": 1}})
                }
                posTag = temp[0]
            }
            if(strings.Contains(posTag,"-")) {
                posTag = strings.Split(posTag,"-")[0]
            }
            word := temp[0]
            if(len(temp) > 2) {
                word = strings.Join(temp[0:len(temp)-1],"/")
                fmt.Println("%s------------%s\n",tokens,word)
            }
            // if(strings.Contains(posTag,"-") ||strings.Contains(posTag,"+") || len(posTag) < 1 || posTag == "*") {
            //     fmt.Printf("%#v\n",tokens)
            // }

            pos.Upsert(bson.M{"posTag":posTag,"word":word},bson.M{"$inc": bson.M{"count": 1}})
        }
	}
}

func InsertPOSNgram(tokens []string, n int) {
    maxWait := time.Duration(5 * time.Second)
    session, err := mgo.DialWithTimeout("127.0.0.1",maxWait)
	if err != nil {
		panic(err)
	}
    defer session.Close()
	session.SetMode(mgo.Monotonic, true)
    pos := session.DB("nlprokz").C("posTags")
	for j := 0; j < (len(tokens) - n); j++ {
        posSeq := make([]string,0)
        for a := 0; a < n; a++ {
            var temp = strings.Split(tokens[j+a],"/")
            if(len(temp) > 1) {
                posSeq = append(posSeq,temp[1])
            }
        }
        pos.Upsert(&Ngrams{Ngram:strings.Join(posSeq," ")},bson.M{"$inc": bson.M{"count": 1}})
	}
}
// func main() {
//
//
//
//
// 	// // Drop Database
// 	// if IsDrop {
// 	// 	err = session.DB("nlprokz").DropDatabase()
// 	// 	if err != nil {
// 	// 		panic(err)
// 	// 	}
// 	// }
//
// 	// Collection People
//
//
// 	// Index
// 	// index := mgo.Index{
// 	// 	Key:        []string{"name", "phone"},
// 	// 	Unique:     true,
// 	// 	DropDups:   true,
// 	// 	Background: true,
// 	// 	Sparse:     true,
// 	// }
//
// 	// err = c.EnsureIndex(index)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
//
// 	// Insert Datas
//
//
// 	// Query One
// 	result := Person{}
// 	err = c.Find(bson.M{"name": "Ale"}).Select(bson.M{"phone": 0}).One(&result)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Phone", result)
//
// 	// Query All
// 	var results []Person
// 	err = c.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&results)
//
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Results All: ", results)
//
// 	// Update
// 	colQuerier := bson.M{"name": "Ale"}
// 	change := bson.M{"$set": bson.M{"phone": "+86 99 8888 7777", "timestamp": time.Now()}}
// 	err = c.Update(colQuerier, change)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	// Query All
// 	err = c.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&results)
//
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Results All: ", results)
// }

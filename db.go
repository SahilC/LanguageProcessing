package main

import (
	//"fmt"
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

func InsertNgram(tokens []string, n int) {
    maxWait := time.Duration(5 * time.Second)
    session, err := mgo.DialWithTimeout("127.0.0.1",maxWait)
	if err != nil {
		panic(err)
	}
    defer session.Close()
	session.SetMode(mgo.Monotonic, true)
    pos := session.DB("nlprokz").C("posTags")
    // wordTags := session.DB("nlprokz").C("wordTags")
	for j := 0; j < (len(tokens) - n); j++ {
		//ngram[strings.Join(tokens[k][j:j+n]," ")] += 1
        // var temp = strings.Split(tokens[j],"/")
        // word := temp[0]
        // posTag := temp[1]
        posSeq := make([]string,0)
        //wordTagSeq := make([]string,0)
        for a := 0; a < n; a++ {
            var temp = strings.Split(tokens[j+a],"/")
            //word := temp[0]
            //fmt.Printf("%s %#v\n",tokens[j],temp)
            if(len(temp) > 1) {
                posSeq = append(posSeq,temp[1])
            }
            //posSeq = append(posSeq,posTag)
        }
        pos.Upsert(&Ngrams{Ngram:strings.Join(posSeq," ")},bson.M{"$inc": bson.M{"count": 1}})
        //fmt.Printf("%s\n",)
        //fmt.Printf("===============\n")
        //posSeq := strings.Join(tokens[j:j+n]," ")
        //err = pos.Insert(&Ngram{Ngram: tokens,count:count Timestamp: time.Now()})

    	// if err != nil {
    	// 	panic(err)
    	// }
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

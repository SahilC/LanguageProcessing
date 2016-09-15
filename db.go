package main

import (
	//"fmt"
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

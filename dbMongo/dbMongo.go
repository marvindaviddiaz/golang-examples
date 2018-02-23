package dbMongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Graph struct {
	Id 	int `bson:"_id"`
	Name string `bson:"name"`
	Type string `bson:"type"`
}

func Main() {

	session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// get collection
	graph := session.DB("test").C("graph")

	// number of elements
	n,_ :=graph.Count()
	log.Println("Number of elements", n)


	//graph.Find(bson.M{"type":"BAR"}).One(&g)

	find := graph.Find(bson.M{"type":"BAR"})
	var g []Graph
	find.All(&g)
	for _, e := range g {
		log.Println(e)
	}

	// Projections
	var names []struct {
		// anonymous functions
		Name string `bson:"name"`
	}
	graph.Find(bson.M{"type":"PIE"}).Select(bson.M{"name" : 1}).All(&names)
	log.Println("")
	log.Println("Getting just names....")
	for _, e := range names {
		log.Println(e)
	}


}
package dbutil

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// GetCountOfCollection returns count of match param count
func GetCountOfCollection(collection *mgo.Collection, base *[]bson.M) int {
	c := 0
	r := bson.M{}
	p := append(*base, bson.M{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": 1}}})
	if err := collection.Pipe(p).One(&r); err == nil {
		c = r["count"].(int)
	}
	return c
}

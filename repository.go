package mongorepo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var dbName string = "testcooking" //database name

type Repository struct {
	TableName string
	DecodeTo  interface{}
}

func (r *Repository) Insert(t interface{}) error {

	session := OpenSession()
	defer session.Close()

	c := Table(session, r.TableName)
	err := c.Insert(&t)

	return err
}

func Table(s *mgo.Session, tName string) *mgo.Collection {
	return s.DB(dbName).C(tName)
}

//params will be bson structs
func (r *Repository) Update(filter interface{}, t interface{}) error {

	session := OpenSession()
	defer session.Close()

	c := Table(session, r.TableName)

	err := c.Update(filter, bson.M{"$set": &t})

	return err
}

func (r *Repository) Select(filter interface{}, take int) (interface{}, error) {

	session := OpenSession()
	defer session.Close()

	c := Table(session, r.TableName)

	result := make([]interface{}, 0)

	err := c.Find(filter).Limit(take).All(&result)

	return result, err
}

func OpenSession() *mgo.Session {
	session, err := mgo.Dial("localhost") //server

	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	return session
}

# mongorepo
Provide some generic functions to insert/update/select data from a mongo database.
<br>Depedencies:
<br>
https://gopkg.in/mgo.v2/bson

Quick example of how to use it (Note that Ingredients, User and Recipes are pseudo-classes):

``` go
package main

import (
	"github.com/mongorepo"
	"fmt"
	"log"
	"net/http"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	http.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
		var Ingredients = []Ingredient{}
		Ingredients = append(Ingredients, Ingredient{"Eggs"})
		Ingredients = append(Ingredients, Ingredient{"Water"})
		repo := mongorepo.Repository{TableName: "Recipes"}

		repo.Insert(Recipes{
			Name:        "Brazilian Cookie",
			Description: "Really nice to eat with friends",
			Ingredients: Ingredients,
			User:        User{"hmendes00", "hmendes00@gmail.com", "123456"}})
		w.Write([]byte("test"))

	})

	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		repo := mongorepo.Repository{TableName: "Recipes"}

		filter := bson.M{"name": "Brazilian Cookie"}
		recipe := bson.M{"name": "Brazilian Cookie Awesome"}

		repo.Update(filter, recipe)
		w.Write([]byte("test"))

	})

	http.HandleFunc("/many", func(w http.ResponseWriter, r *http.Request) {
		repo := mongorepo.Repository{TableName: "Recipes"}

		//filter := nil //bson.M{"name": "%"}

		var result, err = repo.Select(nil, 3)
		check(err)
		var recipes []Recipes
		err = mapstructure.Decode(result, &recipes)
		check(err)

		w.Write([]byte(fmt.Sprintf("%#v", recipes)))

	})

	log.Fatal(http.ListenAndServe(":3310", nil))
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

```

package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"todomvc/service/user/model"
)

func main() {
	session, err := mgo.Dial("127.0.0.1:27017")
	fmt.Println(err)
	var user model.User
	session.Login(&mgo.Credential{Username: "alomerry", Password: "120211"})
	err = session.DB("todo").C("user").Find(bson.M{"name": "alomerry"}).One(&user)
	fmt.Println(err)
	fmt.Println(user)
	fmt.Println(bson.ObjectIdHex("5f115b61eeea7f00310b63d1"))
	session.Close()
}

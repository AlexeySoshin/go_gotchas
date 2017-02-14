package app


import (
	"encoding/json"
	"log"
)


type Controller struct {}


func (c Controller) addUser(jsonInput string) {
	newUser := &UserModel{}
	err := json.Unmarshal([]byte(jsonInput), &newUser)

	if err != nil {
		return
	}

	userFromDB := db.getById(newUser.id)

	if userFromDB != nil && userFromDB.referral != newUser.referral {
		log.Fatalf("User with id %s already exists", userFromDB.id)
	}

	db.save(*newUser)
}

type DB struct {
	users map[string]UserModel
}

func (db DB) save(u UserModel) {

	db.users[u.id] = u
}

func (db DB) getById(id string) *UserModel {
	user, found := db.users[id]

	if found {
		return &user
	}

	return nil
}

type UserModel struct {
	id       string
	name     string
	referral string
}

var db = &DB{}
var controller = &Controller{}


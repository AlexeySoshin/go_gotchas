package app


import (
	"encoding/json"
	"errors"
)


type Controller struct {
	db *DB
}




type DB struct {
	users map[string]UserModel
}


func (db *DB) save(u UserModel) {

	db.users[u.Id] = u
}

func (db *DB) getById(id string) *UserModel {
	user, found := db.users[id]

	if found {
		return &user
	}

	return nil
}

type UserModel struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Referral string `json:"referral"`
}

func (c *Controller) addUser(jsonInput string) error {
	newUser := UserModel{}

	err := json.Unmarshal([]byte(jsonInput), &newUser)

	if err != nil || newUser.Id == "" {
		return errors.New("Parse error")
	}

	userFromDB := db.getById(newUser.Id)

	if userFromDB != nil && userFromDB.Referral != newUser.Referral {
		return errors.New("User already exists with different referral")
	}

	db.save(newUser)

	return nil
}

var db = &DB{users: make(map[string]UserModel)}
var controller = &Controller{db}


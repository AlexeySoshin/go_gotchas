package empty_strings

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

	db.users[u.ID] = u
}

func (db *DB) getByID(id string) *UserModel {
	user, found := db.users[id]

	if found {
		return &user
	}

	return nil
}

type UserModel struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Referral  *string `json:"referral"`
	Referral2 *string `json:"referral2"`
}

func (c *Controller) addUser(jsonInput string) error {
	newUser := UserModel{}

	err := json.Unmarshal([]byte(jsonInput), &newUser)

	if err != nil || newUser.ID == "" {
		return errors.New("Parse error")
	}

	userFromDB := db.getByID(newUser.ID)

	if userFromDB != nil {
		if referralsDiff(userFromDB.Referral, newUser.Referral) || referralsDiff(userFromDB.Referral2, newUser.Referral2) {
			return errors.New("User already exists with different referral")
		}
	}

	db.save(newUser)

	return nil
}

func referralsDiff(ref1 *string, ref2 *string) bool {
	if ref1 == nil && ref2 == nil {
		return false
	}

	if (ref1 == nil && ref2 != nil) || (ref1 != nil && ref2 == nil) {
		return true
	}

	return *ref1 != *ref2
}

var db = &DB{users: make(map[string]UserModel)}
var controller = &Controller{db}

package app

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUser(t *testing.T) {

	db = &DB{users: make(map[string]UserModel)}
	var uuid = "a78a2a52-1fc4-45ac-8d17-fb4967e54938"
	var jsonInput = fmt.Sprintf(`{"id":"%s",
	"name":"Alexey Soshin",
	"referral": "Linkedin"}`, uuid)
	controller.addUser(jsonInput)

	user := db.getByID(uuid)

	if user == nil || user.ID != uuid {
		t.Fail()
	}
}

func TestCreateTwoUsersDifferentReferrals(t *testing.T) {
	db = &DB{users: make(map[string]UserModel)}
	var uuid = "a78a2a52-1fc4-45ac-8d17-fb4967e54938"

	var jsonInput1 = fmt.Sprintf(`{"id":"%s","name":"Alexey Soshin","referral": "Linkedin"}`, uuid)
	var jsonInput2 = fmt.Sprintf(`{"id":"%s","name":"Alexey Soshin","referral": "Facebook"}`, uuid)

	controller.addUser(jsonInput1)
	controller.addUser(jsonInput2)

	user := db.getByID(uuid)

	assert.NotNil(t, user)
	assert.Equal(t, "Linkedin", *user.Referral)
}

func TestUpdateUser(t *testing.T) {
	db = &DB{users: make(map[string]UserModel)}
	var uuid = "a78a2a52-1fc4-45ac-8d17-fb4967e54938"

	var jsonInput1 = fmt.Sprintf(`{"id":"%s","name":"Alexey Soshin","referral": "Linkedin"}`, uuid)
	var jsonInput2 = fmt.Sprintf(`{"id":"%s","name":"Vladimir Soshin","referral": "Linkedin"}`, uuid)

	controller.addUser(jsonInput1)
	controller.addUser(jsonInput2)

	user := db.getByID(uuid)

	assert.NotNil(t, user)
	assert.Equal(t, "Vladimir Soshin", user.Name)
}

func TestUpdateUserNoFirstReferral(t *testing.T) {
	db = &DB{users: make(map[string]UserModel)}
	var uuid = "a78a2a52-1fc4-45ac-8d17-fb4967e54938"

	var jsonInput1 = fmt.Sprintf(`{"id":"%s","name":"Alexey Soshin"}`, uuid)
	var jsonInput2 = fmt.Sprintf(`{"id":"%s","name":"Vladimir Soshin","referral": "Linkedin"}`, uuid)

	controller.addUser(jsonInput1)
	controller.addUser(jsonInput2)

	user := db.getByID(uuid)

	assert.NotNil(t, user)
	assert.Equal(t, "Alexey Soshin", user.Name)
}

func TestUpdateUserNoSecondReferral(t *testing.T) {
	db = &DB{users: make(map[string]UserModel)}
	var uuid = "a78a2a52-1fc4-45ac-8d17-fb4967e54938"

	var jsonInput1 = fmt.Sprintf(`{"id":"%s","name":"Alexey Soshin","referral": "Linkedin"}`, uuid)
	var jsonInput2 = fmt.Sprintf(`{"id":"%s","name":"Vladimir Soshin"}`, uuid)

	controller.addUser(jsonInput1)
	controller.addUser(jsonInput2)

	user := db.getByID(uuid)

	assert.NotNil(t, user)
	assert.Equal(t, "Alexey Soshin", user.Name)
}

func TestUpdateUserNoReferrals(t *testing.T) {
	db = &DB{users: make(map[string]UserModel)}
	var uuid = "a78a2a52-1fc4-45ac-8d17-fb4967e54938"

	var jsonInput1 = fmt.Sprintf(`{"id":"%s","name":"Alexey Soshin"}`, uuid)
	var jsonInput2 = fmt.Sprintf(`{"id":"%s","name":"Vladimir Soshin"}`, uuid)

	controller.addUser(jsonInput1)
	controller.addUser(jsonInput2)

	user := db.getByID(uuid)

	assert.NotNil(t, user)
	assert.Equal(t, "Vladimir Soshin", user.Name)
}

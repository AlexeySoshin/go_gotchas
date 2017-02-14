package app

import (
	"fmt"
	"testing"
)

func TestCreateUser(t *testing.T) {

	var uuid = "a78a2a52-1fc4-45ac-8d17-fb4967e54938"
	var jsonInput = fmt.Sprintf(`{"id":"%s","name":"Alexey Soshin","referral": "Linkedin"}`, uuid)
	controller.addUser(jsonInput)

	user := db.getById(uuid)

	if user == nil || user.Id != uuid {
		t.Fail()
	}
}

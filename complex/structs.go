package complex

import (
	"encoding/json"
	"fmt"
	"strings"
)

// SelectJSONFields
// Use anonymous struct to select JSON string fields.
func SelectJSONFields() {
	type User struct {
		name, username, email string
	}

	leanne := User{
		name:     "Leanne Graham",
		username: "Bret",
		email:    "Sincere@april.biz",
	}

	var builder strings.Builder
	json.NewEncoder(&builder).Encode(struct {
		Name  string
		Email string
	}{
		Name:  leanne.name,
		Email: leanne.email,
	})

	fmt.Println(builder.String())
}

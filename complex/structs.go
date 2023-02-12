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

// Type assertions and interface{}
type typeA struct {
	data string
}

type typeB struct {
	data string
}

func newTypeA(data string) *typeA {
	a := &typeA{data: data}
	return a
}

func newTypeB(data string) *typeB {
	b := &typeB{data: data}
	return b
}

func (a *typeA) String() string {
	return fmt.Sprintf("typeA - (%v)", a.data)
}

func (b *typeB) String() string {
	return fmt.Sprintf("typeB - (%v)", b.data)
}

// TypeAssertion
func TypeAssertion() {
	vals := []interface{}{
		newTypeA("typeA"),
		newTypeB("typeB"),
	}

	for _, val := range vals {
		if i, ok := val.(*typeA); ok {
			fmt.Println(i)
		}

		if i, ok := val.(*typeB); ok {
			fmt.Println(i)
		}
	}
}

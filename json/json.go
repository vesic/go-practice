package json

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// When a struct defines an embedded struct, fields of the embedded struct
// are promoted as they are defined by the enclosing type.
func PromoteField() {
	type Geo struct {
		Lat  string `json:"lat"`
		Long string `json:"long"`
	}

	type User struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Geo
	}

	user := User{
		Name:     "Ervin",
		Username: "ervin",
		Geo: Geo{
			Lat:  "-37",
			Long: "81",
		},
	}

	var builder strings.Builder
	encoder := json.NewEncoder(&builder)
	encoder.Encode(&user)

	fmt.Println(builder.String())
}

//

type User struct {
	Name     string
	Username string
}

func (u *User) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"id":        rand.Int63(),
		"user_id":   u.Name,
		"user_name": u.Username,
	}

	jsn, err := json.Marshal(m)
	return jsn, err
}

// CustomEncoding
// Implement the Marshaler interface to create custom JSON encoding.
func CustomEncoding() {
	user := User{
		Name:     "Ervin",
		Username: "ervin",
	}

	var builder strings.Builder
	encoder := json.NewEncoder(&builder)
	encoder.Encode(&user)
	fmt.Println(builder.String())
	// {"id":914288788548139488,"user_id":"Ervin","user_name":"ervin"}
}

// DecodingFromFile
// Decode configuration from file
/*
{
	"id": "v1",
	"runtime": "python27",
	"threadsafe": true
}
*/
func DecodingFromFile() {
	type Config struct {
		Id         string
		Runtime    string
		threadsafe bool
	}
	var config Config

	data, err := os.ReadFile("config.json")
	if err != nil {
		panic("file read error")
	}

	decoder := json.NewDecoder(strings.NewReader(string(data)))
	err = decoder.Decode(&config)
	if err != nil {
		panic("json.Decode error")
	}

	fmt.Println(config.Runtime)
	// "python27"
}

// DecodingMultiConfigFromFile
// Decode multiple configuration from file.
/*
[
	{
		"id": "v1",
		"runtime": "python27",
		"threadsafe": true
	},
	{
		"id": "v2",
		"runtime": "python3",
		"threadsafe": true
	}
]
*/
func DecodingMultiConfigFromFile() {
	type Config struct {
		Id         string
		Runtime    string
		threadsafe bool
	}
	var config []Config

	data, err := os.ReadFile("config.json")
	if err != nil {
		panic("file read error")
	}

	decoder := json.NewDecoder(strings.NewReader(string(data)))
	err = decoder.Decode(&config)
	if err != nil {
		panic("json.Decode error")
	}

	for _, c := range config {
		fmt.Println(c.Id, c.Runtime, c.threadsafe)
	}
	// v1 python27 false
	// v2 python3 false
}

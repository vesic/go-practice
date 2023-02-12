package networking

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GetDataAndDecodeResponse
// Retrieve data and decode response.
func GetDataAndDecodeResponse() {
	type User struct {
		Id             int
		Name, Username string
	}

	r, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user User
	err = json.Unmarshal(data, &user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v, %v, %v\n", user.Id, user.Name, user.Username)
}

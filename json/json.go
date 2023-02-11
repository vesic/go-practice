package json

import (
	"encoding/json"
	"fmt"
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

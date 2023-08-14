package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal(body, x); err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
		}
	} else {
		fmt.Println("Error reading request body:", err)
	}
}

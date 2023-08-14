package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(x); err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
}

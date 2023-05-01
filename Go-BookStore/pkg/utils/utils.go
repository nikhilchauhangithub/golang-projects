package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) { //By using an empty interface, you can create functions and methods that can accept a wide variety of inputs, and then inspect and manipulate those inputs dynamically at runtime.
if body, err := ioutil.ReadAll(r.Body); err ==nil {
	if err := json.Unmarshal([]byte(body), x); err!=nil {
		return
	}
}
}
package app

import (
	"encoding/json"
	"net/http"
)

type Robot struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

var RoboArray = []Robot{}


func FetchRobots() {
	c := http.Client{}
	res, err := c.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		println(err.Error())
	}else {
		decoder := json.NewDecoder(res.Body)
		decoder.Decode(&RoboArray)
		println(RoboArray)
	}



}
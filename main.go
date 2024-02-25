package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
)

type Activity struct {
	Activity      string  `json:"activity"`
	Accessibility float32 `json:"accessibility"`
	Type          string  `json:"type"`
	Participants  uint8   `json:"participants"`
	Price         float32 `json:"price"`
	Link          string  `json:"link"`
	Key           string  `json:"key"`
}

func main() {
	response, err := http.Get("http://www.boredapi.com/api/activity/")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var activity Activity
	json.Unmarshal(data, &activity)

	values := reflect.ValueOf(activity)
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		fmt.Println(types.Field(i).Name, values.Field(i))
	}
}

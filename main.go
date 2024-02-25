package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Activity struct {
	Activity      string  `json:"activity"`
	Type          string  `json:"type"`
	Participants  uint8   `json:"participants"`
	Accessibility float32 `json:"accessibility"`
	Price         float32 `json:"price"`
	Link          string  `json:"link"`
	// Key           string  `json:"key"`
}

func (a *Activity) print() {
	values := reflect.ValueOf(*a)
	types := values.Type()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendRow([]interface{}{types.Field(0).Name, values.Field(0)})
	t.AppendSeparator()

	for i := 1; i < values.NumField(); i++ {
		t.AppendRow([]interface{}{types.Field(i).Name, values.Field(i)})
	}

	t.Render()
}

func main() {
	const BORED_API_URL string = "http://www.boredapi.com/api/activity/"

	response, err := http.Get(BORED_API_URL)
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
	activity.print()
}

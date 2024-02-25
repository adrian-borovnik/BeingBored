package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

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

func printActivities(activities []Activity) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Activity", "Type", "Participants", "Accessibility", "Price", "Link"})
	for i, activity := range activities {
		t.AppendRow(table.Row{i + 1, activity.Activity, activity.Type, activity.Participants, activity.Accessibility, activity.Price, activity.Link})
	}
	t.Render()
}

func fetchActivity(wg *sync.WaitGroup, ch chan Activity) {

	defer func() {
		wg.Done()
	}()

	const BORED_API_URL string = "http://www.boredapi.com/api/activity/"

	response, err := http.Get(BORED_API_URL)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var activity Activity
	json.Unmarshal(data, &activity)
	ch <- activity
}

func main() {
	const ACTIVITY_COUNT int = 10

	var wg sync.WaitGroup
	ch := make(chan Activity, ACTIVITY_COUNT)

	startTime := time.Now()

	for i := 0; i < ACTIVITY_COUNT; i++ {
		wg.Add(1)
		go fetchActivity(&wg, ch)
	}

	wg.Wait()
	close(ch)

	var activities []Activity
	for activity := range ch {
		activities = append(activities, activity)
	}

	duration := time.Since(startTime)
	fmt.Println("Fetch time:", duration)

	// for _, a := range activities {
	// 	a.print()
	// }

	printActivities(activities)

}

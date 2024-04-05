package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Event struct {
	Time time.Time
}

func main() {
	t := time.Now()
	event1 := Event{t}

	b, err := json.Marshal(event1)
	if err != nil {
		panic(err)
	}

	var event2 Event
	err = json.Unmarshal(b, &event2)
	if err != nil {
		panic(err)
	}

	fmt.Println(event1 == event2)
	fmt.Println(event1.Time)
	fmt.Println(event2.Time)
}

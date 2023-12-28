package main

import (
	"fmt"
	"encoding/json"
	"time"
)

type Learning struct {
	Category    string `json:"category"`
	Description string `json:"description"`
	Created time.Time `json:"created"`
}

func (l *Learning) setLearning(category, description string) {
	l.Category = category
	l.Description = description
	l.Created = time.Now()
}

func (l *Learning) setLearningWithTime(category, description string, created time.Time) {
	l.Category = category
	l.Description = description
	l.Created = created
}


func (l *Learning) setCategory(category string) {
	l.Category = category
}


func (l *Learning) setDescription(description string) {
	l.Description = description
}

func main() {
	l1 := Learning{}
	l1.setLearning("garden", "plant tomatoes")
	res1A, _ := json.Marshal(l1)
	fmt.Println(string(res1A))

}

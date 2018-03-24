package taskDatabase

import (
	"fmt"
)

type Task struct {
	Index int    `json:"index" bson:"index"`
	Title string `json:"title" bson:"title"`
	Time  string `json:"time" bson:"time"`
	Body  string `json:"body" bson:"body"`
}

func (t Task) String() string {
	formatString := "Index: %d\nTitle: %s\nTime: %s\nBody: %s\n"
	return fmt.Sprintf(formatString, t.Index, t.Title, t.Time, t.Body)
}

type TaskPosition struct {
	StartIndex int `json:"startIndex" bson:"startIndex"`
	EndIndex   int `json:"endIndex" bson:"endIndex"`
}

func (t TaskPosition) String() string {
	formatString := "Start Index: %d\nEnd Index: %d\n"
	return fmt.Sprintf(formatString, t.StartIndex, t.EndIndex)
}

package tasks

import (
	"fmt"
)

type Task struct {
	Index int    `json:"index"`
	Title string `json:"title"`
	Time  string `json:"time"`
	Body  string `json:"body"`
}

func (t Task) String() string {
	formatString := "Index: %d\nTitle: %s\nTime: %s\nBody: %s\n"
	return fmt.Sprintf(formatString, t.Index, t.Title, t.Time, t.Body)
}

type TaskPosition struct {
	StartIndex int `json:"startIndex"`
	EndIndex   int `json:"endIndex"`
}

func (t TaskPosition) String() string {
	formatString := "Start Index: %d\nEnd Index: %d\n"
	return fmt.Sprintf(formatString, t.StartIndex, t.EndIndex)
}

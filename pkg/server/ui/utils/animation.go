package utils

import (
	"encoding/json"
	"fmt"
)

type Animation struct {
	Name     string `json:"name"`
	Duration int    `json:"duration"`
	Moment   string `json:"moment"`
}

func NewStartAnimation(name string, duration int) string {
	return animation(name, duration, "start")
}

func NewEndAnimation(name string, duration int) string {
	return animation(name, duration, "end")
}

func animation(name string, duration int, moment string) string {
	obj := Animation{
		Name: fmt.Sprintf("animate__animated animate__%v animate__delay-%vms", name, duration),
		Duration: duration,
		Moment: moment,
	}
	str, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(str)
}
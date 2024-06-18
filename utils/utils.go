package utils

import (
	"fmt"
	"time"
)

func convertDateString(str string) *time.Time {
	layout := "2006-01-02 15:04:05"
	date, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	}
	return &date
}

package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	date := "2022-12-30 18:00:00+8000"
	parse, _ := time.Parse("2006-01-02 15:04:05", date)
	fmt.Println(parse)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the time study of gloang")
	presentTime := time.Now()

	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))

	createDate := time.Date(2020, time.September, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}

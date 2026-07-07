package main

import (
	"fmt"
	"time"
)

// func main() {
// 	mainTime()
// }

func mainTime() {
	now := time.Now()

	fmt.Println("now       :", now)
	fmt.Println("Local     :", now.Local())
	zoneName, zoneOffset := now.Zone()
	fmt.Println("Zone      :", zoneName, zoneOffset)
	fmt.Println("UTC       :", now.UTC())
	fmt.Println("Unix      :", now.Unix())
	fmt.Println("UnixNano  :", now.UnixNano())

	utc := time.Date(2022, time.July, 22, 0, 0, 0, 0, time.UTC)
	wib := time.Date(2024, 07, 7, 7, 0, 0, 0, time.Local)

	fmt.Println(utc.UTC())
	fmt.Println(wib.UTC())

	parseRFC3339, _ := time.Parse(time.RFC3339, "2022-07-22T07:32:22Z")
	fmt.Println("parseRFC3339:", parseRFC3339)

	parseMyTime, _ := time.Parse("2006-01-02 15:04:05", "2024-07-07 19:33:00")
	fmt.Println("parseMyTime:", parseMyTime)

	fmt.Println(now.Format("2006-01-02 15:04:05"))
}

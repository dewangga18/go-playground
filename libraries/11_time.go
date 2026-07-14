package main

import (
	"fmt"
	"time"
)

// func main() {
// 	mainTime()
// 	mainDuration()
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

	utc := time.Date(2022, time.July, 22, 5, 0, 0, 0, time.UTC)
	wib := time.Date(2024, 07, 7, 7, 0, 0, 0, time.Local)

	fmt.Println(utc.UTC())
	fmt.Println(wib.UTC())

	parseRFC3339, _ := time.Parse(time.RFC3339, "2022-07-22T07:32:22Z")
	fmt.Println("parseRFC3339:", parseRFC3339)

	parseMyTime, _ := time.Parse("2006-01-02 15:04:05", "2024-07-07 19:33:00")
	fmt.Println("parseMyTime:", parseMyTime)

	fmt.Println(now.Format("2006-01-02 15:04:05"))

	fmt.Println("Year:", utc.Year())
	fmt.Println("Month:", utc.Month())
	fmt.Println("Day:", utc.Day())
	fmt.Println("Hour:", utc.Hour())
}

func mainDuration() {
	duration1 := time.Second * 100         // 100 seconds
	duration2 := time.Minute * 10          // 10 minutes
	duration3 := time.Hour * 1             // 1 hour

	fmt.Println("Seconds", duration1.Seconds())
	fmt.Println("Minutes", duration2.Minutes())
	fmt.Println("Hours", duration3.Hours())

	diff := duration3 - duration2 - duration1
	fmt.Println("Duration", diff)

	parseDuration, _ := time.ParseDuration("2h30m")
	fmt.Println("ParseDuration", parseDuration)
	fmt.Println("ParseDuration hour", parseDuration.Hours())
	fmt.Println("ParseDuration min", parseDuration.Minutes())
}

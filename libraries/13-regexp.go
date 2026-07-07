package main

import (
	"fmt"
	"regexp"
)

// func main() {
// 	mainRegexp()
// }

func mainRegexp() {
	text := "golang regexp is fun and golang is awesome"
	emailText := "user@example.com, admin@test.org, invalid-email"
	csvLine := "a,b,c,d,e"

	fmt.Println("=== Compile — compile pattern (returns error if invalid) ===")
	re, err := regexp.Compile(`golang`)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Compiled:", re) // golang

	fmt.Println("\n=== MustCompile — compile or panic (use when pattern is certain) ===")
	re2 := regexp.MustCompile(`golang`)
	fmt.Println("MustCompiled:", re2) // golang

	fmt.Println("\n=== MatchString — check if pattern matches anywhere ===")
	fmt.Println(regexp.MustCompile(`golang`).MatchString(text))       // true
	fmt.Println(regexp.MustCompile(`java`).MatchString(text))         // false

	fmt.Println("\n=== FindString — first match ===")
	fmt.Println(regexp.MustCompile(`golang`).FindString(text))        // golang

	reDigit := regexp.MustCompile(`\d+`)
	fmt.Println(reDigit.FindString("order 99 price 500"))             // 99

	fmt.Println("\n=== FindAllString — all matches (n = -1 for all) ===")
	all := regexp.MustCompile(`golang`).FindAllString(text, -1)
	fmt.Println(all)                                                  // [golang golang]
	fmt.Println("Count:", len(all))                                   // 2

	// limit results
	limited := regexp.MustCompile(`golang`).FindAllString(text, 1)
	fmt.Println(limited)                                              // [golang]

	fmt.Println("\n=== ReplaceAllString — replace matches with new string ===")
	replaced := regexp.MustCompile(`golang`).ReplaceAllString(text, "Go")
	fmt.Println(replaced)                                             // Go regexp is fun and Go is awesome

	// replace digits with placeholder
	replacedDigit := regexp.MustCompile(`\d+`).ReplaceAllString("phone 123, zip 456", "***")
	fmt.Println(replacedDigit)                                        // phone ***, zip ***

	// capture groups in replacement ($1, $2, ...)
	reEmail := regexp.MustCompile(`(\w+)@(\w+\.\w+)`)
	masked := reEmail.ReplaceAllString(emailText, "$1 at $2")
	fmt.Println(masked)                                               // user at example.com, admin at test.org, invalid-email

	fmt.Println("\n=== Split — split string by pattern ===")
	parts := regexp.MustCompile(`,`).Split(csvLine, -1)
	fmt.Println(parts)                                                // [a b c d e]

	// split with limit
	limitedParts := regexp.MustCompile(`,`).Split(csvLine, 3)
	fmt.Println(limitedParts)                                         // [a b c,d,e]

	// split on whitespace
	words := regexp.MustCompile(`\s+`).Split("hello   world  foo", -1)
	fmt.Println(words)                                                // [hello world foo]

	fmt.Println("\n=== FindStringSubmatch — match with capture groups ===")
	logLine := "ERROR 2024-07-07 15:30:00 Connection timeout"
	reLog := regexp.MustCompile(`(\w+) (\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}) (.+)`)
	matches := reLog.FindStringSubmatch(logLine)
	fmt.Println(matches)                                              // [ERROR 2024-07-07 15:30:00 Connection timeout ERROR 2024-07-07 15:30:00 Connection timeout]
	fmt.Println("Level:", matches[1])                                 // ERROR
	fmt.Println("Time:", matches[2])                                  // 2024-07-07 15:30:00
	fmt.Println("Message:", matches[3])                               // Connection timeout

	// extract emails with named groups
	reNamed := regexp.MustCompile(`(?P<name>\w+)@(?P<domain>\w+\.\w+)`)
	emailMatch := reNamed.FindStringSubmatch("user@example.com")
	fmt.Println("Name:", emailMatch[reNamed.SubexpIndex("name")])      // user
	fmt.Println("Domain:", emailMatch[reNamed.SubexpIndex("domain")])  // example.com
}

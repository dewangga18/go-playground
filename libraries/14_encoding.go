package main

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// func main() {
// 	mainEncoding()
// }

func isValidHex(s string) bool {
	_, err := hex.DecodeString(s)
	return err == nil
}

func mainEncoding() {
	fmt.Println("=== encoding/json — Marshal & Unmarshal ===")

	type User struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email,omitempty"`
	}

	// Marshal — struct/slice/map → JSON bytes
	users := []User{
		{"Budi", 25, "budi@example.com"},
		{"Sari", 30, ""},
	}

	jsonBytes, _ := json.Marshal(users)
	fmt.Println("Marshal:", string(jsonBytes))
	// [{"name":"Budi","age":25,"email":"budi@example.com"},{"name":"Sari","age":30}]

	jsonBytesIndent, _ := json.MarshalIndent(users, "", "  ")
	fmt.Println("MarshalIndent:")
	fmt.Println(string(jsonBytesIndent))
	// [
	//   {
	//     "name": "Budi",
	//     "age": 25,
	//     "email": "budi@example.com"
	//   },
	//   {
	//     "name": "Sari",
	//     "age": 30
	//   }
	// ]

	// Unmarshal — JSON bytes → struct/slice/map
	var decoded []User
	jsonStr := `[{"name":"Agus","age":28,"email":"agus@test.com"},{"name":"Dewi","age":22}]`
	err := json.Unmarshal([]byte(jsonStr), &decoded)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Unmarshal:", decoded)
	// [{Agus 28 agus@test.com} {Dewi 22 }]

	// Unmarshal to map (when struct is unknown)
	var raw []map[string]any
	json.Unmarshal([]byte(jsonStr), &raw)
	fmt.Println("Raw map:", raw)
	// [map[age:28 email:agus@test.com name:Agus] map[age:22 name:Dewi]]

	// Marshal map
	mapData := map[string]any{"name": "Test", "count": 42, "active": true}
	mapJSON, _ := json.Marshal(mapData)
	fmt.Println("Map to JSON:", string(mapJSON))
	// {"active":true,"count":42,"name":"Test"}

	fmt.Println("\n=== encoding/csv — Read & Write ===")

	// Write CSV
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	writer.Write([]string{"Name", "Age", "City"})
	writer.Write([]string{"Budi", "25", "Jakarta"})
	writer.Write([]string{"Sari", "30", "Bandung"})
	writer.Flush()

	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("CSV written:")
	fmt.Println(buf.String())
	// Name,Age,City
	// Budi,25,Jakarta
	// Sari,30,Bandung

	// Read CSV
	reader := csv.NewReader(strings.NewReader(buf.String()))
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CSV read:")
	for _, record := range records {
		fmt.Println(" ", record)
	}
	// [Name Age City]
	// [Budi 25 Jakarta]
	// [Sari 30 Bandung]

	fmt.Println("\n=== encoding/base64 — Encode & Decode ===")

	data := []byte("Hello, Golang!")

	// Standard base64 (with + and /)
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println("Base64 Std:", encoded)
	// SGVsbG8sIEdvbGFuZyE=

	decodedB64, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Println("Decoded:", string(decodedB64))
	// Hello, Golang!

	// URL-safe base64 (with - and _ instead of + and /)
	encodedURL := base64.URLEncoding.EncodeToString(data)
	fmt.Println("Base64 URL:", encodedURL)
	// SGVsbG8sIEdvbGFuZyE=

	decodedURL, _ := base64.URLEncoding.DecodeString(encodedURL)
	fmt.Println("Decoded URL:", string(decodedURL))
	// Hello, Golang!

	fmt.Println("\n=== encoding/hex — Encode & Decode ===")

	dataHex := []byte("Hello, Golang!")

	// Encode to hex string
	encodedHex := hex.EncodeToString(dataHex)
	fmt.Println("Hex Encode:", encodedHex)
	// 48656c6c6f2c20476f6c616e6721

	// Decode from hex string
	decodedHex, _ := hex.DecodeString(encodedHex)
	fmt.Println("Hex Decode:", string(decodedHex))
	// Hello, Golang!

	// Check if string is valid hex
	fmt.Println("Valid Hex:", isValidHex(encodedHex))
	// Valid Hex: true

	fmt.Println("Invalid Hex:", isValidHex("xyz"))
	// Invalid Hex: false
}

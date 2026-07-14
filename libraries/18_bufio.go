package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

// func main() {
// 	mainBufio()
// }

func mainBufio() {
	fmt.Println("=== bufio.NewReader — create a buffered reader ===")
	fmt.Println("bufio wraps an io.Reader with an internal buffer for efficient reads")
	fmt.Println("Default buffer size is 4096 bytes")

	reader := bufio.NewReader(strings.NewReader("Hello, buffered world!"))
	// Read one byte at a time from buffer (more efficient than unbuffered)
	b, _ := reader.ReadByte()
	fmt.Printf("First byte: %c\n", b) // H

	// Read next 5 bytes
	peekBuf, _ := reader.Peek(5)
	fmt.Printf("Peek next 5: %q\n", string(peekBuf)) // "ello,"

	// Read until delimiter
	line1, _ := reader.ReadString(' ')
	fmt.Printf("ReadString until ' ': %q\n", line1) // "ello, "

	fmt.Println("\n=== bufio.ReadString — read until delimiter (returns string) ===")
	data := "name=Budi\nage=25\ncity=Jakarta\n"
	sc := bufio.NewReader(strings.NewReader(data))

	for {
		line, err := sc.ReadString('\n')
		fmt.Printf("  Line: %q", line)
		if err != nil {
			fmt.Println(" (EOF)")
			break
		}
		fmt.Println()
	}
	// Line: "name=Budi\n"
	// Line: "age=25\n"
	// Line: "city=Jakarta\n"
	// Line: "" (EOF)

	fmt.Println("\n=== bufio.ReadBytes — read until delimiter (returns []byte) ===")
	data2 := "a,b,c,d,e"
	br := bufio.NewReader(strings.NewReader(data2))

	for i := 0; i < 5; i++ {
		chunk, err := br.ReadBytes(',')
		if err != nil {
			// Last chunk (no trailing delimiter) — still has data
			fmt.Printf("  Chunk (last): %q\n", string(chunk))
			break
		}
		fmt.Printf("  Chunk: %q\n", string(chunk))
	}
	// Chunk: "a,"
	// Chunk: "b,"
	// Chunk: "c,"
	// Chunk: "d,"
	// Chunk (last): "e"

	fmt.Println("\n=== bufio.Scanner — read lines/tokens efficiently ===")
	text := "line one\nline two\nline three\n"

	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		fmt.Printf("  Line: %s\n", scanner.Text()) // line one, line two, line three
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Scanner with file
	fmt.Println("\nScanning a file:")
	fileData := "header\ncontent\nfooter\n"
	fileScanner := bufio.NewScanner(strings.NewReader(fileData))
	fileScanner.Split(bufio.ScanLines) // default, but explicit for clarity
	lineNum := 1
	for fileScanner.Scan() {
		fmt.Printf("  Line %d: %s\n", lineNum, fileScanner.Text())
		lineNum++
	}

	fmt.Println("\n=== bufio Scanner — scan by words (ScanWords) ===")
	sentence := "Go is awesome and fast!"
	wordScanner := bufio.NewScanner(strings.NewReader(sentence))
	wordScanner.Split(bufio.ScanWords)

	wordCount := 0
	for wordScanner.Scan() {
		fmt.Printf("  Word %d: %s\n", wordCount+1, wordScanner.Text())
		wordCount++
	}
	fmt.Printf("  Total words: %d\n", wordCount)
	// Word 1: Go
	// Word 2: is
	// Word 3: awesome
	// Word 4: and
	// Word 5: fast!
	// Total words: 5

	// ScanBytes — individual bytes
	fmt.Println("\nScanBytes:")
	byteScanner := bufio.NewScanner(strings.NewReader("ABC"))
	byteScanner.Split(bufio.ScanBytes)
	for byteScanner.Scan() {
		fmt.Printf("  Byte: %s\n", byteScanner.Text())
	}
	// Byte: A, Byte: B, Byte: C

	// ScanRunes — individual runes (handles multi-byte UTF-8)
	fmt.Println("\nScanRunes:")
	runeScanner := bufio.NewScanner(strings.NewReader("Halo, dunia!"))
	runeScanner.Split(bufio.ScanRunes)
	for runeScanner.Scan() {
		fmt.Printf("  Rune: %s\n", runeScanner.Text())
	}
	// H, a, l, o, ,,  , d, u, n, i, a, !

	fmt.Println("\n=== bufio.NewWriter — create a buffered writer ===")
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)

	n, _ := writer.Write([]byte("Hello "))
	fmt.Printf("  Wrote %d bytes to buffer (not yet flushed)\\n", n) // 6

	writer.Write([]byte("World"))
	writer.Write([]byte("!"))
	fmt.Printf("  Buffer is empty until Flush: %q\n", buf.String()) // "" (empty!)

	writer.Flush()
	fmt.Printf("  After Flush: %q\n", buf.String()) // "Hello World!"

	fmt.Println("\n=== bufio.WriteString — write a string directly ===")
	var buf2 bytes.Buffer
	w := bufio.NewWriter(&buf2)

	// WriteString is more efficient than Write([]byte(s))
	n2, _ := w.WriteString("This is a string!")
	fmt.Printf("  Wrote %d bytes via WriteString\\n", n2) // 18

	// Multiple writes
	w.WriteString(" Another string.")
	w.Flush()
	fmt.Printf("  After Flush: %q\n", buf2.String()) // "This is a string! Another string."

	fmt.Println("\n=== bufio.Flush — flush buffered data to underlying writer ===")
	fmt.Println("Always call Flush() after writing — data is buffered until flushed!")

	var buf3 bytes.Buffer
	bw := bufio.NewWriter(&buf3)

	bw.WriteString("data1 ")
	bw.WriteString("data2 ")
	bw.WriteString("data3")
	// NOT flushed yet — buf3 is still empty
	fmt.Printf("  Before Flush: %q\n", buf3.String()) // ""

	bw.Flush()
	fmt.Printf("  After Flush:  %q\n", buf3.String()) // "data1 data2 data3"

	// Demonstrating buffer size (default 4096)
	fmt.Println("\nDefault buffer size is 4096 bytes")
	fmt.Println("bufio.NewWriterSize(w, size) lets you customize")
	fmt.Println("bufio.NewWriterSize(os.Stdout, 8192) — 8KB buffer")

	fmt.Println("\n=== Practical: read file line by line ===")
	fmt.Println("Opening and reading libraries/18-bufio.go headers:")
	file, err := os.Open("libraries/18-bufio.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScan := bufio.NewScanner(file)
	for i := 0; i < 5 && fileScan.Scan(); i++ {
		fmt.Printf("  %s\n", fileScan.Text())
	}
	if err := fileScan.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n=== Practical: write to file with buffered writer ===")
	tmpFile := "temp_bufio_test.txt"
	f, err := os.Create(tmpFile)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFile)

	fw := bufio.NewWriter(f)
	fw.WriteString("line 1\n")
	fw.WriteString("line 2\n")
	fw.WriteString("line 3\n")
	fw.Flush()
	f.Close()

	// Verify by reading it back
	verifyData, _ := os.ReadFile(tmpFile)
	fmt.Printf("  Written to %s:\\n%s", tmpFile, string(verifyData))

	fmt.Println("\n=== bufio split functions summary ===")
	fmt.Println("  bufio.ScanLines   — split by newlines (default)")
	fmt.Println("  bufio.ScanWords   — split by whitespace")
	fmt.Println("  bufio.ScanBytes   — split by individual bytes")
	fmt.Println("  bufio.ScanRunes   — split by individual runes (UTF-8 safe)")
	fmt.Println("  Custom: scanner.Split(yourSplitFunc) — for custom splitting")
}

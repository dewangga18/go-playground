package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// func main() {
// 	mainIO()
// }

func mainIO() {
	fmt.Println("=== io.Reader — read from any data source ===")
	fmt.Println("Reader is an interface: Read(p []byte) (n int, err error)")

	// strings.NewReader — Reader from string
	reader := strings.NewReader("Hello, Golang io!")
	buf := make([]byte, 8) // read in 8-byte chunks

	fmt.Println("\nReading in chunks:")
	for {
		n, err := reader.Read(buf)
		fmt.Printf("  read %d bytes: %q\n", n, buf[:n])
		if err == io.EOF {
			break
		}
	}

	// bytes.NewReader — Reader from []byte
	fmt.Println("\nbytes.NewReader:")
	byteReader := bytes.NewReader([]byte{65, 66, 67, 68, 69}) // A, B, C, D, E
	full, _ := io.ReadAll(byteReader)
	fmt.Printf("  Bytes: %v\n", full)          // [65 66 67 68 69]
	fmt.Printf("  String: %s\n", string(full)) // ABCDE

	fmt.Println("\n=== io.Writer — write to any data destination ===")
	fmt.Println("Writer is an interface: Write(p []byte) (n int, err error)")

	// bytes.Buffer implements both Reader and Writer
	var bufWriter bytes.Buffer
	n, _ := bufWriter.Write([]byte("Hello, Writer!"))
	fmt.Printf("  Wrote %d bytes: %q\n", n, bufWriter.String()) // 15, "Hello, Writer!"

	// Multiple writes
	bufWriter.Write([]byte(" More data."))
	fmt.Printf("  After append: %q\n", bufWriter.String()) // "Hello, Writer! More data."

	fmt.Println("\n=== Custom Reader/Writer — implementing the interfaces ===")
	upper := &UpperCaseReader{data: "golang is fun", pos: 0}
	result, _ := io.ReadAll(upper)
	fmt.Printf("  UpperCaseReader: %s\n", string(result)) // GOLANG IS FUN

	var upperWriter UpperCaseWriter
	upperWriter.Write([]byte("hello world"))
	fmt.Printf("  UpperCaseWriter: %s\n", upperWriter.String()) // HELLO WORLD

	fmt.Println("\n=== io.ReadAll — read everything from a reader ===")
	r := strings.NewReader("Read me entirely!")
	data, _ := io.ReadAll(r)
	fmt.Printf("  ReadAll: %q (%d bytes)\n", string(data), len(data)) // "Read me entirely!" (18)

	// Empty reader
	empty := strings.NewReader("")
	emptyData, _ := io.ReadAll(empty)
	fmt.Printf("  Empty: %q\n", string(emptyData)) // ""

	fmt.Println("\n=== io.Copy — copy from Reader to Writer ===")
	src := strings.NewReader("Data to copy")
	var dst bytes.Buffer
	written, _ := io.Copy(&dst, src)
	fmt.Printf("  Copied %d bytes: %q\n", written, dst.String()) // 13, "Data to copy"

	// Copy to Discard (discard data)
	discardSrc := strings.NewReader("This will be discarded")
	discarded, _ := io.Copy(io.Discard, discardSrc)
	fmt.Printf("  Discarded %d bytes\n", discarded) // 23

	fmt.Println("\n=== io.CopyN — copy exactly N bytes ===")
	src2 := strings.NewReader("Long data string to copy partially")
	var dst2 bytes.Buffer
	writtenN, _ := io.CopyN(&dst2, src2, 10)
	fmt.Printf("  CopiedN %d bytes: %q\n", writtenN, dst2.String()) // 10, "Long data "

	// CopyN with remaining
	var dst3 bytes.Buffer
	remaining, _ := io.CopyN(&dst3, src2, 5)
	fmt.Printf("  Next %d bytes: %q\n", remaining, dst3.String()) // 5, "strin"

	fmt.Println("\n=== io.MultiWriter — write to multiple writers simultaneously ===")
	var buf1, buf2 bytes.Buffer
	mw := io.MultiWriter(&buf1, &buf2)

	mw.Write([]byte("Multi-writer test"))
	fmt.Printf("  buf1: %q\n", buf1.String()) // "Multi-writer test"
	fmt.Printf("  buf2: %q\n", buf2.String()) // "Multi-writer test"

	// MultiWriter + Copy — tee data to multiple destinations
	var log1, log2 bytes.Buffer
	logWriter := io.MultiWriter(&log1, &log2)
	io.Copy(logWriter, strings.NewReader("Log entry: OK"))
	fmt.Printf("  log1: %q\n", log1.String()) // "Log entry: OK"
	fmt.Printf("  log2: %q\n", log2.String()) // "Log entry: OK"

	// MultiWriter with Discard — write to one logger while also discarding
	var logger bytes.Buffer
	mwDiscard := io.MultiWriter(&logger, io.Discard)
	mwDiscard.Write([]byte("Only logger sees this"))
	fmt.Printf("  logger: %q\n", logger.String()) // "Only logger sees this"

	fmt.Println("\n=== io.TeeReader — read and write simultaneously (like 'tee') ===")
	var teeBuf bytes.Buffer
	teeSrc := strings.NewReader("Tee reader test data")
	teeReader := io.TeeReader(teeSrc, &teeBuf)

	// Read from TeeReader — data also gets written to teeBuf
	teeResult, _ := io.ReadAll(teeReader)
	fmt.Printf("  Read: %q\n", string(teeResult))    // "Tee reader test data"
	fmt.Printf("  Teed to buf: %q\n", teeBuf.String()) // "Tee reader test data"

	// Practical: TeeReader + MultiWriter — log while processing
	var processLog, auditLog bytes.Buffer
	auditWriter := io.MultiWriter(&processLog, &auditLog)
	auditSrc := strings.NewReader("audit trail entry")
	auditReader := io.TeeReader(auditSrc, auditWriter)

	processed, _ := io.ReadAll(auditReader)
	fmt.Printf("  Processed: %q\n", string(processed))     // "audit trail entry"
	fmt.Printf("  Process log: %q\n", processLog.String())  // "audit trail entry"
	fmt.Printf("  Audit log: %q\n", auditLog.String())      // "audit trail entry"

	fmt.Println("\n=== io.Discard — /dev/null for Go ===")
	// Discard is io.Writer that discards everything
	nDiscard, _ := io.Discard.Write([]byte("gone forever"))
	fmt.Printf("  Wrote %d bytes to Discard — all gone\n", nDiscard) // 13

	// Discard in MultiWriter (already shown above)
	// Discard with Copy (already shown above)

	fmt.Println("\n=== Practical: io.Copy with file (os.File implements Reader+Writer) ===")
	fmt.Println("  os.File, bytes.Buffer, strings.NewReader, http.Response.Body")
	fmt.Println("  ... all implement io.Reader and/or io.Writer")
	fmt.Println("  This means you can pipe data between them seamlessly!")
}

// UpperCaseReader — custom Reader that uppercases everything
type UpperCaseReader struct {
	data string
	pos  int
}

func (u *UpperCaseReader) Read(p []byte) (int, error) {
	if u.pos >= len(u.data) {
		return 0, io.EOF
	}

	// Uppercase on first call, then read in chunks
	if u.pos == 0 {
		u.data = strings.ToUpper(u.data)
	}

	n := copy(p, u.data[u.pos:])
	u.pos += n

	if u.pos >= len(u.data) {
		return n, io.EOF
	}
	return n, nil
}

// UpperCaseWriter — custom Writer that uppercases before writing
type UpperCaseWriter struct {
	buf bytes.Buffer
}

func (w *UpperCaseWriter) Write(p []byte) (int, error) {
	return w.buf.Write([]byte(strings.ToUpper(string(p))))
}

func (w *UpperCaseWriter) String() string {
	return w.buf.String()
}

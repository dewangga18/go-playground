package main

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

// func main() {
// 	mainPath()
// }

func mainPath() {
	fmt.Println("=== path — URL-style path operations (forward slashes) ===")
	fmt.Println("path package works with /-separated paths (like URLs)")

	fmt.Println("\n--- Base — get the last element of a path ---")
	fmt.Println(path.Base("/a/b/c/file.txt"))       // file.txt
	fmt.Println(path.Base("/a/b/c/"))                // c (trailing slash removed)
	fmt.Println(path.Base(""))                       // .

	fmt.Println("\n--- Dir — get the directory part (everything except Base) ---")
	fmt.Println(path.Dir("/a/b/c/file.txt"))         // /a/b/c
	fmt.Println(path.Dir("/a/b/c/"))                 // /a/b/c
	fmt.Println(path.Dir("file.txt"))                // .
	fmt.Println(path.Dir(""))                        // .

	fmt.Println("\n--- Ext — get the file extension (including dot) ---")
	fmt.Println(path.Ext("/a/b/c/file.txt"))         // .txt
	fmt.Println(path.Ext("/a/b/c/archive.tar.gz"))   // .gz (only last extension!)
	fmt.Println(path.Ext("/a/b/c/"))                 // (empty — dirs have no ext)
	fmt.Println(path.Ext("file"))                    // (empty — no extension)

	fmt.Println("\n--- Join — join path elements with / ---")
	fmt.Println(path.Join("a", "b", "c"))            // a/b/c
	fmt.Println(path.Join("a/", "/b/", "/c/"))      // a/b/c (cleaned)
	fmt.Println(path.Join("/a", "b", "c"))           // /a/b/c
	fmt.Println(path.Join("a", "..", "b", "c"))      // b/c (resolve ..)

	fmt.Println("\n--- Split — split path into Dir + Base ---")
	dir, file := path.Split("/a/b/c/file.txt")
	fmt.Printf("Dir: %q, File: %q\n", dir, file)    // Dir: "/a/b/c/", File: "file.txt"

	dir2, file2 := path.Split("file.txt")
	fmt.Printf("Dir: %q, File: %q\n", dir2, file2)  // Dir: "", File: "file.txt"

	fmt.Println("\n--- Clean — clean up a path (remove ., .., double slashes) ---")
	fmt.Println(path.Clean("a/b/../c"))              // a/c
	fmt.Println(path.Clean("a/./b/./c"))             // a/b/c
	fmt.Println(path.Clean("/a/b//c///d"))           // /a/b/c/d
	fmt.Println(path.Clean("/../a/b"))               // /a/b (cannot go above root)
	fmt.Println(path.Clean(""))                       // .

	fmt.Println("\n\n=== filepath — OS-aware path operations (uses / on Linux/macOS, \\ on Windows) ===")
	fmt.Println("filepath adapts to OS path separator")

	fmt.Println("\n--- filepath.Join — join using OS separator ---")
	fmt.Println(filepath.Join("a", "b", "c"))           // a/b/c
	fmt.Println(filepath.Join("/a", "b", "c"))          // /a/b/c
	fmt.Println(filepath.Join("a", "..", "b", ".", "c")) // b/c

	fmt.Println("\n--- filepath.Base ---")
	fmt.Println(filepath.Base("/a/b/c/file.txt"))       // file.txt
	fmt.Println(filepath.Base("/a/b/c/"))               // c
	fmt.Println(filepath.Base(""))                       // .

	fmt.Println("\n--- filepath.Dir ---")
	fmt.Println(filepath.Dir("/a/b/c/file.txt"))        // /a/b/c
	fmt.Println(filepath.Dir("/a/b/c/"))                // /a/b/c
	fmt.Println(filepath.Dir("file.txt"))               // .

	fmt.Println("\n--- filepath.Ext ---")
	fmt.Println(filepath.Ext("/a/b/c/file.txt"))        // .txt
	fmt.Println(filepath.Ext("archive.tar.gz"))         // .gz
	fmt.Println(filepath.Ext("noext"))                  // (empty)

	fmt.Println("\n--- filepath.WalkDir — walk directory tree ---")
	// WalkDir visits each file/directory recursively
	// callback func(path string, d fs.DirEntry, err error) error
	fmt.Println("Walking src/ directory:")
	filepath.WalkDir("src", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			fmt.Printf("  [DIR]  %s\n", p)
		} else {
			fmt.Printf("  [FILE] %s (%d bytes)\n", p, mustFileSize(p))
		}
		return nil
	})

	fmt.Println("\n--- filepath.Glob — match files by pattern ---")
	matches, _ := filepath.Glob("src/**/*.go")
	fmt.Println("All .go files in src/:")
	for _, m := range matches {
		fmt.Println(" ", m)
	}

	// Match a specific pattern
	mdMatches, _ := filepath.Glob("**/*.md")
	fmt.Println("\nAll .md files:")
	for _, m := range mdMatches {
		fmt.Println(" ", m)
	}

	fmt.Println("\n--- filepath.Abs — get absolute path ---")
	rel := "src/basics/hello-world.go"
	abs, _ := filepath.Abs(rel)
	fmt.Println("Relative:", rel)
	fmt.Println("Absolute:", abs)

	// Also works with current dir
	cwd, _ := filepath.Abs(".")
	fmt.Println("Current dir:", cwd)
}

// mustFileSize returns file size, or 0 on error (helper for WalkDir demo)
func mustFileSize(p string) int64 {
	info, err := os.Stat(p)
	if err != nil {
		return 0
	}
	return info.Size()
}

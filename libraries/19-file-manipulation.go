package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// func main() {
// 	mainFileManipulation()
// }

func mainFileManipulation() {
	err := createNewFile("logs/hi.txt", "sample hi!")
	if err != nil {
		fmt.Println("failed to create file", err.Error())
	} else {
		fmt.Println("success")
	}

	_ = createDirectory("logs")
	fmt.Println("created logs directory")

	err = readDirectory("fake-logs")
	if err != nil {
		fmt.Println("read fake directory", err.Error())
	}

	_ = readDirectory("logs")
	fmt.Println("success read logs directory")

	_ = createNewFile("logs/hi.txt", "hi everyone")
	fmt.Println("success create new file")

	content, _ := readFile("logs/hi.txt")
	fmt.Println("content:", content)

	_ = updateFile("logs/hi.txt", "sample rewrite all file!")
	fmt.Println("success update file")
	
	content, _ = readFile("logs/hi.txt")
	fmt.Println("content:", content)

	content, _ = readFile("logs/hi.txt")
	fmt.Println("content:", content)

	_ = appendFile("logs/hi.txt", "\nit's line 2", "\nhi it's line 3")
	fmt.Println("success append")

	content, _ = readFile("logs/hi.txt")
	fmt.Println("content:", content)

	content, _ = readFile2("logs/hi.txt")
	fmt.Println("content:", content)
	
	_ = copyFile("logs/hi.txt", "logs/hi-new.txt")
	fmt.Println("copy new file")

	_ = renameFile("logs/hi.txt", "logs/hi-old.txt")
	fmt.Println("rename file")

	_ = deleteFile("logs/hi.txt")
	fmt.Println("delete file")

	_ = deleteFileAll("logs")
	fmt.Println("delete directory all")

	err = readDirectory("logs")
	if err != nil {
		fmt.Println("failed to read directory", err.Error())
	}
}

func createNewFile(name string, msg string) error {
	file, err := os.OpenFile(
		name, 
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC, // use os.O_TRUNC for truncate new file
		// os.O_CREATE | os.O_EXCL, // for return error when file existed
		0666,
	)
	if err != nil {
		return err
	}
	
	defer file.Close()
	
	file.WriteString(msg)
	return nil
}

func createDirectory(pathname string) error {
	return os.Mkdir(pathname, 0755)
}

func readDirectory(pathname string) error {
	files, err := os.ReadDir(pathname)
	if err != nil {
		return err
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
	return nil
}

func readFile(name string) (string, error) {
	content, err := os.ReadFile(name)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func readFile2(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()
	
	reader := bufio.NewReader(file)
	var msg string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		msg += string(line)
	}

	return msg, nil
}

func updateFile(name string, content string) error {
	return os.WriteFile(name, []byte(content), 0644)
}

func appendFile(name string, content ...string) error {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, msg := range content {
		_, err = file.WriteString(msg)
		if err != nil {
			return err
		}
	}
	
	return err
}

func copyFile(src string, name string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	dst, err := os.Create(name)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, source)
	return err
}

func renameFile(src string, name string) error {
	err := os.Rename(src, name)
	return err
}

// delete one file or empty directory
func deleteFile(name string) error {
	return os.Remove(name)
}

// delete file or directory and all its contents
func deleteFileAll(dirpath string) error {
	return os.RemoveAll(dirpath)
}
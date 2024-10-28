package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type UserPrompt int
type Person struct {
	fname string
	lname string
}

func validateFile(path string) (bool, *os.File, error) {
	fileptr, _er := os.Open(path)
	if _er != nil {
		return false, nil, _er
	}
	return true, fileptr, nil
}

func getFileAsString(file *os.File) (string, int, error) {
	fileinfo, err := file.Stat()
	if err != nil {
		return "", 0, err
	}
	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	byteCount, err := file.Read(buffer)
	if err != nil {
		return "", 0, err
	}

	return string(buffer), byteCount, nil
}

func getPersonSlice(fileptr *os.File) []Person {
	reader := bufio.NewReader(fileptr)
	var personSliceofname = make([]Person, 0)
	for {
		lineofname, _, _err := reader.ReadLine()
		if _err == io.EOF {
			break
		}
		fields := strings.Split(string(lineofname), " ")
		person := Person{fname: fields[0], lname: fields[1]}
		personSliceofname = append(personSliceofname, person)
	}
	return personSliceofname
}

func main() {
	var filePath string
	fmt.Println("Enter Absolute File path : ")
	fmt.Scan(&filePath)

	isValid, fileptr, _ := validateFile(filePath)
	if !isValid {
		fmt.Printf("%s: Invalid file\n", filePath)
		os.Exit(0)
	}

	names := getPersonSlice(fileptr)
	for _, name := range names {
		fmt.Printf("fname: [%-20s]; lname: [%-20s]\n", name.fname, name.lname)
	}
}

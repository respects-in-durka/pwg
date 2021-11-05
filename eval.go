package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func generateFileName() string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyz1234567890")
	result := make([]rune, 8)

	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return "/tmp/" + string(result) + ".go"
}

func createFile(content string) string {
	name := generateFileName()
	err := ioutil.WriteFile(name, []byte(content), 0755)

	if err == nil {
		return name
	}
	panic(err)
}

func deleteFile(filename string) {
	err := os.Remove(filename)

	if err != nil {
		panic(err)
	}
}

func run(code string) (error, string) {
	file := createFile(code)
	cmd := exec.Command("go", "run", file)
	out, err := cmd.Output()
	deleteFile(file)

	if err == nil {
		return nil, string(out)
	}
	return err, ""
}
package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"testing"
)

func TestAnswer(t *testing.T) {
	inbuf := readFile("./stdin.txt")
	stubStdin(inbuf, func() {
		main()
	})
}

func stubStdin(inbuf string, fn func()) {
	// stub os.Stdin with os.Pipe
	inr, inw, _ := os.Pipe()
	inw.Write([]byte(inbuf))
	inw.Close()
	os.Stdin = inr
	// overwrite scanner
	scanner = bufio.NewScanner(os.Stdin)
	fn()
}

func readFile(fileName string) string {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	text, err := readfile("tet.txt") // if we don't know about pre-defined func just click that func and enter f12
	if err != nil {
		panic(err) // eof is a func
	} else {
		fmt.Println(text)
	}

}

func readfile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("unable to open %s", filename)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

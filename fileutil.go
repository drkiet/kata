package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readAll(fileName string) []byte {
	fmt.Println("filename: " + fileName)
	data, err := ioutil.ReadFile(fileName)
	check(err)
	return data
}

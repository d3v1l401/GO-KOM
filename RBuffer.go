package main

import (
	"fmt"
	"io/ioutil"
)

var keyList []byte

func importKeyList() []byte {
	keyList, err := ioutil.ReadFile("keyList.d3v")
	if err != nil {
		fmt.Printf("%s (keyList importation)", err)
	}
	return keyList
}

func processBuffer(buffer []byte, size uint32) []byte {

	output := make([]byte, size)

	var index uint32
	for index = 0; index < size; index++ {
		output[index] = buffer[index] ^ keyList[index]
	}

	return output
}

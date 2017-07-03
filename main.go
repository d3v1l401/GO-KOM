package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

var dataName string

func main() {
	counter := 0
	importKeyList()
fromstart:
	counter++
	dataName = fmt.Sprintf("data%03d.kom", counter)
	data, err := os.Open(dataName)

	if err != nil {
		panic(err)
	}

	arrayed := make([]byte, 27)
	data.Read(arrayed)
	_magic := string(arrayed)

	arrayed = make([]byte, 25)
	data.Read(arrayed)

	arrayed = make([]byte, 4)
	data.Read(arrayed)
	_filesNum := binary.LittleEndian.Uint32(arrayed)

	data.Read(arrayed)
	_crc := binary.LittleEndian.Uint32(arrayed)

	data.Read(arrayed)
	data.Read(arrayed)

	data.Read(arrayed)
	_xmlsize := binary.LittleEndian.Uint32(arrayed)

	arrayed = make([]byte, _xmlsize)
	data.Read(arrayed)
	_xmlBuffer := string(arrayed)

	fmt.Printf("Files number: %d\n", _filesNum)
	fmt.Printf("Is Compressed: %d\n", _crc)
	fmt.Printf("XML Size: %d\n", _xmlsize)
	fmt.Printf("Header: %s\n", _magic)

	var _headerInfo = KOMHeader{_magic, _crc, _filesNum, _xmlsize, _crc, _xmlBuffer}

	//fileList := interpretXML(_headerInfo)
	interpretXML(_headerInfo, data)
	goto fromstart
}

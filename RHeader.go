package main

import (
	"bytes"
	"compress/zlib"
	"encoding/xml"
	"fmt"
	"os"
)

type KOMHeader struct {
	_header       string
	_crc          uint32
	_filesNumber  uint32
	_xmlSize      uint32
	_isCompressed uint32
	_xmlBuffer    string
}

type FileList struct {
	XMLName xml.Name   `xml:"Files"`
	Files   []FileInfo `xml:"File"`
}

type FileInfo struct {
	Name           string `xml:"Name,attr"`
	Size           uint32 `xml:"Size,attr"`
	CompressedSize uint32 `xml:"CompressedSize,attr"`
	Checksum       string `xml:"Checksum,attr"`
	FileTime       string `xml:"FileTime,attr"`
	Algorithm      uint32 `xml:"Algorithm,attr"`
	Buffer         []byte
}

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func fileSave(name string, buffer []byte) {
	CreateDirIfNotExist(dataName)
	f, err := os.Create(dataName + "_out\\" + name)

	if err != nil {
		panic(err)
	}

	defer f.Close()
	f.Write(buffer)
	f.Sync()
	f.Close()
}

func interpretXML(header KOMHeader, data *os.File) map[string]FileInfo {
	var fileList = make(map[string]FileInfo)

	var slave FileList

	err := xml.Unmarshal([]byte(header._xmlBuffer), &slave)
	if err != nil {
		panic(err)
	}

	var index uint32
	for index = 0; index < header._filesNumber; index++ {
		slave.Files[index].Buffer = make([]byte, slave.Files[index].CompressedSize)

		taken, err := data.Read(slave.Files[index].Buffer)

		if err != nil {
			fmt.Printf("[!] %s for %s.\n", err, slave.Files[index].Name)
		}

		if uint32(taken) == slave.Files[index].CompressedSize {
			fileList[slave.Files[index].Name] = slave.Files[index]
			fmt.Printf("[+] Reading \"%s\" (%d -> %d) [%d].\n", slave.Files[index].Name, taken, slave.Files[index].Size, slave.Files[index].Algorithm)

			switch slave.Files[index].Algorithm {

			case 0, 2:

				var loled = make([]byte, slave.Files[index].Size)
				input := bytes.NewReader(slave.Files[index].Buffer)
				r, err := zlib.NewReader(input)

				if err != nil {
					fmt.Printf("[!] %s for %s.\n", err, slave.Files[index].Name)
					fileSave(slave.Files[index].Name, slave.Files[index].Buffer)
				} else {
					r.Read(loled)
					fileSave(slave.Files[index].Name, loled)
				}
			case 1, 3:

				var loled = make([]byte, slave.Files[index].Size)
				input := bytes.NewReader(slave.Files[index].Buffer)
				r, err := zlib.NewReader(input)

				if err != nil {
					//fmt.Printf("Cannot decompress: %s, saving raw buffer.\n", err)
					fileSave(slave.Files[index].Name, slave.Files[index].Buffer)
				} else {
					r.Read(loled)

					//loled = processBuffer(loled, 84)

					fileSave(slave.Files[index].Name, loled)
				}

			default:
				fileSave(slave.Files[index].Name, slave.Files[index].Buffer)
			}

		} else {
			fmt.Printf("Could not read %s.\n", slave.Files[index].Name)
		}
	}

	return fileList
}

package main

import "fmt"

type Shape interface {
	Area()
	Perimeter()
}

type FileOperation interface {
	Read() string
	Write(data string)
}

type TextFile struct {
	filename string
}

type BinaryFile struct {
	filename string
}

func (tf TextFile) Read() string {
	return "Hello"
}

func (tf TextFile) Write(data string) {

}

func (tf BinaryFile) Read() string {
	return "Hello"
}

func (tf BinaryFile) Write(data string) {

}

func main() {
	binFile := BinaryFile{
		filename: "example.bin",
	}

	textFile := TextFile{
		filename: "example.txt",
	}

	fileOps := []FileOperation{textFile, binFile}

	fmt.Printf("fileOps: %v\n", fileOps)

	PrintInterface(binFile)
	PrintInterface(textFile)
	PrintInterface(5)
	PrintInterface("Dave")

}

func PrintInterface(iface interface{}) {
	fmt.Println(iface)
}

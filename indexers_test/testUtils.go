package indexers

import (
	"io/ioutil"
	"os"
)

var inputsPath string = "/tmp/"

var inputFilePath string = inputsPath + "inputs.txt"
var subfile1Path string = inputsPath + "foo1.txt"
var subfile2Path string = inputsPath + "foo2.txt"

var inputFileContent string = "romeo_I_II.txt" + "\n" + "romeo_III_IV.txt" + "\n" + "romeo_V.txt"
var subfile1Content string = "this is a test. cool."
var subfile2Content string = "this is also a test.\nboring."

// InitializeTests helps initialize testing data
func InitializeTests() *os.File {

	f := writeTestFile(inputFilePath, inputFileContent)
	f1 := writeTestFile(subfile1Path, subfile1Content)
	f2 := writeTestFile(subfile2Path, subfile2Content)

	f1.Close()
	f2.Close()
	return f
}

func writeTestFile(path, data string) *os.File {
	dat := []byte(data)
	err := ioutil.WriteFile(path, dat, 0644)
	check(err)
	f, err := os.Open(path)
	check(err)

	return f
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

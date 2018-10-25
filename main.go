package main

import (
	"flag"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	outputFile := flag.String("out", "out.yaml", "output file")
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatalln("Usage is: FinYAML <inputfile> ")
	}
	templateFile := flag.Args()[0]

	input, err := ioutil.ReadFile(templateFile)
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}
	inputString := string(input)
	start := strings.Index(inputString, "{{")
	end := strings.Index(inputString, "}}")
	runes := []rune(inputString)

	fileName := string(runes[start+3 : end-1])

	inputFileData, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}
	fileContents := string(inputFileData)
	fileContents = "\"" + strings.Replace(fileContents, "\n", "\\n", -1) + "\""
	result := string(runes[0:start]) + fileContents

	err = ioutil.WriteFile(*outputFile, []byte(result), 0644)
	if err != nil {
		log.Fatalf("Could not write file: %v", err)
	}

}

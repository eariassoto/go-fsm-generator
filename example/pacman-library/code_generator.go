// MIT License
//
// Copyright (c) 2018 Emmanuel Arias
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//
// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates the code for the SCXMl files It can be invoked by
// running go generate
package main

import (
	"log"
	"os"

	"github.com/eariassoto/scxml_fsm_generator/pkg/scxml_fsm_generator"
)

func main() {
	inputFiles := []string{"scxml/ghost.xml"}
	outputFiles := []string{"ghost.go"}

	for i, scxmlFilename := range inputFiles {
		scxmlFile, err := os.Open(scxmlFilename)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer scxmlFile.Close()

		outputFile, err := os.Create(outputFiles[i])
		if err != nil {
			log.Fatal(err)
			return
		}
		defer outputFile.Close()

		scxml_fsm_generator.GenerateFSMCodeForSCXML(scxmlFile, outputFile)
	}
}

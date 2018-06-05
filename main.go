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
package main

import (
	"flag"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/eariassoto/scxml_fsm_generator/pkg/fsmgen"
)

func main() {
	outputDir := flag.String("output_dir", "", "Output directory to write the generated code")
	flag.Parse()

	for _, inputFile := range flag.Args() {

		scxmlFile, err := os.Open(inputFile)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer scxmlFile.Close()

		_, scxmlFileName := filepath.Split(scxmlFile.Name())
		scxmlFileNameSplit := strings.Split(scxmlFileName, ".")
		scxmlFileName = scxmlFileNameSplit[0] + ".go"
		outputFile, err := os.Create(path.Join(*outputDir, scxmlFileName))
		if err != nil {
			log.Fatal(err)
			return
		}
		defer outputFile.Close()

		fsmgen.GenerateFSMCodeForSCXML(scxmlFile, outputFile)
	}
}

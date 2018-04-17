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
/* This program generates finite states machines from XML files. It parses
   the description of a finite state machine and generates the proper Go
   code for the state machine.

   The generator will not provide you boilerplates for the callbacks described
   in the SCXML. The user ought to implement those callbacks under the same
   package.

   The generator creates constants for all the FSM's states and transitions.
   It provides this function:
       NewStateMachine(data interface{}) *StateMachine
   to create a new FSM instance. The data paramater will be stored and passed
   to all of the FSM's callbacks. The StateMachine struct provides two functions:
   		GetNextState(stimulus Stimulus)  (State, error)
   		MoveNextState(stimulus Stimulus) (State, error)
   the former to look ahead the next movement, and the latter to move the FSM.

   You can install this program in your machine and use the go:generate command
   in your program:
   		//go:generate go-fsm-generator -input_files=my_fsm1.xml,my_fsm2.xml
*/
package main

import (
	"flag"
	"log"
	"os"
	"path"
	"strings"
)

func main() {
	inputDir := flag.String("input_dir", "scxml", "Directory where the XML files are located")
	inputFiles := flag.String("input_files", "fsm.xml", "Comma-separated list of XML files")
	fsmPackage := flag.String("package", "main", "Package for the generated code")
	outputDir := flag.String("output_dir", "", "Output directory to write the generated code")

	flag.Parse()

	for _, inputFile := range strings.Split(*inputFiles, ",") {

		xmlFile, err := os.Open(path.Join(*inputDir, inputFile))
		if err != nil {
			log.Fatal(err)
			return
		}
		defer xmlFile.Close()

		fsm, err := ParseXMLFile(xmlFile)
		if err != nil {
			log.Fatal(err)
			return
		}

		fsm.Package = *fsmPackage

		outputFileName := fsm.Name + ".go"
		outputFile, err := os.Create(path.Join(*outputDir, outputFileName))
		if err != nil {
			log.Fatal(err)
			return
		}
		defer outputFile.Close()

		err = GenerateCodeForFSM(fsm, outputFile)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

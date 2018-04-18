# scxml_fsm_generator
This library generates finite states machines from XML files. It parses the description of a finite state machine and generates the proper Go code for the state machine.

The generator will not provide you boilerplates for the callbacks described in the SCXML. The user ought to implements those callbacks under the same package.

The generator creates constants for all the FSM's states and transitions. It provides this function:
```go
NewStateMachine(data interface{}) *StateMachine
```
to create a new FSM instance. The data paramater will be stored and passed to all of the FSM's callbacks. The StateMachine struct provides two functions:
```go
GetNextState(stimulus Stimulus)  (State, error)
MoveNextState(stimulus Stimulus) (State, error)
```
the former to look ahead the next movement, and the latter to move the FSM.

You can use this library in your project creating a small program similar to this one:

```go
// The following directive is necessary to make the package coherent:
// +build ignore
// This program generates the code for the SCXMl files It can be invoked by
// running go generate
package main

import (
	"github.com/eariassoto/scxml_fsm_generator"
	"log"
	"os"
)

func main() {
	inputFiles := []string{"scxml/my_fsm.xml"}
	outputFiles := []string{"ghost_fsm.go"}

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
```
to invoke the generator using the go:generate command:
```go
//go:generate go run fsm_code_generator.go
```

# go-fsm-generator
This program generates finite states machines from XML files. It parses the description of a finite state machine and generates the proper Go code for the state machine.

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

You can install this program in your machine and use the go:generate command in your program:
```go
//go:generate go-fsm-generator -input_files=my_fsm1.xml,my_fsm2.xml
```

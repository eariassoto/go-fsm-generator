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
package fsm_code_generator

import (
	"io"
	"text/template"

	"github.com/eariassoto/scxml_fsm_generator/internal/scxml_to_fsm_parser"
)

const FSM_CODE_TEMPLATE_FILE = `{{- $FSMName := .Name -}}
// Code generated by go generate; DO NOT EDIT
// This file was generated by robots at {{ .Timestamp }}
package {{.Package}}

type {{$FSMName}}State int

type {{$FSMName}}Stimulus int

const (
{{- range $index, $element := .States}}
	{{.Id}} {{$FSMName}}State = {{$index}}
{{- end}}
)

{{$FSMStimulus := .GetTransitions -}}
const (
{{- range $index, $element := $FSMStimulus}}
	{{$element}} {{$FSMName}}Stimulus = {{$index}}
{{- end}}
)

type {{$FSMName}}StateMachine struct {
	CurrentState {{$FSMName}}State
	Data         interface{}
}

type fsmError struct {
	s string
}

func (e *fsmError) Error() string {
	return e.s
}

func New{{$FSMName}}StateMachine(data interface{}) *{{$FSMName}}StateMachine {
	return &{{$FSMName}}StateMachine{ {{- .InitialState}}, data}
}

func Get{{$FSMName}}StimulusString(stimulus {{$FSMName}}Stimulus) string {
	switch stimulus {
{{- range $element := $FSMStimulus}}
	case {{$element}}:
		return "{{$element}}"
{{- end}}
	default:
		return ""
	}
}

func Get{{$FSMName}}StateString(state {{$FSMName}}State) string {
	switch state {
{{- range .States}}
	case {{.Id}}:
		return "{{.Id}}"
{{- end}}
	default:
		return ""
	}
}

func (fsm *{{$FSMName}}StateMachine) MoveNextState(stimulus {{$FSMName}}Stimulus) ({{$FSMName}}State, error) {
	nextState, err := fsm.GetNextState(stimulus)
	if err != nil {
		return fsm.CurrentState, err
	}
	cbs := GetOnExitCallback(fsm.CurrentState)
	for _, cb := range cbs {
		cb(fsm.Data)
	}
	if fsm.CurrentState == nextState {
		cbs = GetOnLoopCallback(fsm.CurrentState)
		for _, cb := range cbs {
			cb(fsm.Data)
		}
	} else {
		cbs = GetOnEntryCallback(nextState)
		for _, cb := range cbs {
			cb(fsm.Data)
		}
	}
	fsm.CurrentState = nextState
	return nextState, nil
}

func (fsm *{{$FSMName}}StateMachine) GetNextState(stimulus {{$FSMName}}Stimulus) ({{$FSMName}}State, error) {
	switch fsm.CurrentState {
{{- range $state := .States}}
		case {{.Id}}:
			switch stimulus {
{{- range $transition := $state.Transitions}}
			case {{.Event}}:
				{{- if ne .Cond ""}}
				if {{.Cond}}(fsm.Data) == true {
					return {{.Target}}, nil
				}
				return fsm.CurrentState, &fsmError{"Conditional callback returned false"}
				{{- else}}
				return {{.Target}}, nil
				{{- end -}}
{{- end}}
			default:
				return fsm.CurrentState, &fsmError{"Invalid stimulus"}
			}
{{- end}}
	default:
		return fsm.CurrentState, &fsmError{"Invalid stimulus"}
	}
}

func GetOnEntryCallback(state {{$FSMName}}State) []func(interface{}) {
	switch state {
{{- range .States}}
{{- if ne .OnEntryAction.Name ""}}
	case {{.Id}}:
		return []func(interface{}){ {{- .OnEntryAction.Name}}}
{{- end}}
{{- end}}
	default:
		return []func(interface{}){}
	}
}

func GetOnExitCallback(state {{$FSMName}}State) []func(interface{}) {
	switch state {
{{- range .States}}
{{- if ne .OnExitAction.Name ""}}
	case {{.Id}}:
		return []func(interface{}){ {{- .OnExitAction.Name}}}
{{- end}}
{{- end}}
	default:
		return []func(interface{}){}
	}
}

func GetOnLoopCallback(state {{$FSMName}}State) []func(interface{}) {
	switch state {
{{- range .States}}
{{- if ne .OnLoopAction.Name ""}}
	case {{.Id}}:
		return []func(interface{}){ {{- .OnLoopAction.Name}}}
{{- end}}
{{- end}}
	default:
		return []func(interface{}){}
	}
}

`

// GenerateCodeForFSM takes a parsed SCXML file and generated the code
// from a pre-defined template. All of the code for a FSM will be written
// in a single Writer.
func GenerateCodeForFSM(fsm *scxml_to_fsm_parser.ScxmlStateMachine, outputFile io.Writer) error {
	fsmTemplate := template.Must(template.New(fsm.Name).Parse(FSM_CODE_TEMPLATE_FILE))
	return fsmTemplate.Execute(outputFile, fsm)
}

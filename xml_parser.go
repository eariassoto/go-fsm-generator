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
	"encoding/xml"
	"io"
	"io/ioutil"
	"time"
)

// XMLStateMachine represents the entire FSM described on a SCXML document.
type XMLStateMachine struct {
	Name         string `xml:"name,attr"`
	Package      string
	States       []XMLState `xml:"state"`
	InitialState string     `xml:"initial,attr"`
	Timestamp    time.Time
}

// GetTransitions returns an unordered list of all the different transitions
// appearing in the FSM. This list will be used to define the FSM's stimulus.
func (fsm *XMLStateMachine) GetTransitions() []string {
	set := make(map[string]struct{})
	for _, state := range fsm.States {
		for _, transition := range state.Transitions {
			set[transition.Event] = struct{}{}
		}
	}
	transitions := make([]string, len(set))
	idx := 0
	for key, _ := range set {
		transitions[idx] = key
		idx++
	}
	return transitions
}

// XMLState represents a FSM state. It can have actions for when the FSM
// exits/enters/re-enters that particular state.
type XMLState struct {
	Id            string          `xml:"id,attr"`
	OnEntryAction XMLAction       `xml:"onEntry"`
	OnLoopAction  XMLAction       `xml:"onLoop"`
	OnExitAction  XMLAction       `xml:"onExit"`
	Transitions   []XMLTransition `xml:"transition"`
}

// XMLTransition represents a transition in the FSM. A transitions defines
// the name of the event and which state the FSM should move. A transition's
// event's name will be formerly called a Stimulus.
// TODO[eariassoto]: Add the conditional callback
type XMLTransition struct {
	Event  string `xml:"event,attr"`
	Target string `xml:"target,attr"`
}

// XMLAction represents one or more state's callbacks. The user can register
// one callback function or a comma-separated list of functions.
type XMLAction struct {
	Name string `xml:"action,attr"`
}

// ParseXMLFile takes a Reader interface and tries to parse a SCXML document
// into a XMLStateMachine struct.
func ParseXMLFile(xmlFile io.Reader) (*XMLStateMachine, error) {
	fsm := XMLStateMachine{Timestamp: time.Now().UTC()}

	buffer, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(buffer, &fsm)

	return &fsm, nil
}

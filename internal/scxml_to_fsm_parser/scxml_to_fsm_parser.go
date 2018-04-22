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
package scxml_to_fsm_parser

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"time"
)

// ScxmlStateMachine represents the entire FSM described on a SCXML document.
type ScxmlStateMachine struct {
	Timestamp    time.Time
	Name         string       `xml:"name,attr"`
	Package      string       `xml:"package,attr"`
	States       []ScxmlState `xml:"state"`
	InitialState string       `xml:"initial,attr"`
}

// GetTransitions returns an unordered list of all the different transitions
// appearing in the FSM. This list will be used to define the FSM's stimulus.
func (fsm *ScxmlStateMachine) GetTransitions() []string {
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

// ScxmlState represents a FSM state. It can have actions for when the FSM
// exits/enters/re-enters that particular state.
type ScxmlState struct {
	Id            string            `xml:"id,attr"`
	OnEntryAction ScxmlAction       `xml:"onEntry"`
	OnLoopAction  ScxmlAction       `xml:"onLoop"`
	OnExitAction  ScxmlAction       `xml:"onExit"`
	Transitions   []ScxmlTransition `xml:"transition"`
}

// ScxmlTransition represents a transition in the FSM. A transitions defines
// the name of the event and which state the FSM should move. A transition's
// event's name will be formerly called a Stimulus.
// TODO[eariassoto]: Add the conditional callback
type ScxmlTransition struct {
	Event  string `xml:"event,attr"`
	Target string `xml:"target,attr"`
}

// ScxmlAction represents one or more state's callbacks. The user can register
// one callback function or a comma-separated list of functions.
type ScxmlAction struct {
	Name string `xml:"action,attr"`
}

// ParseScxmlFile takes a Reader interface and tries to parse a SCXML document
// into a ScxmlStateMachine struct.
func ParseScxmlFile(xmlFile io.Reader) (*ScxmlStateMachine, error) {
	fsm := ScxmlStateMachine{Timestamp: time.Now().UTC()}

	buffer, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(buffer, &fsm)

	return &fsm, nil
}

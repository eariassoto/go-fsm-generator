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
package tomlparser

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"time"

	toml "github.com/pelletier/go-toml"
)

// TomlState represents a FSM state. It can have actions for when the FSM
// exits/enters/re-enters that particular state.
type TomlState struct {
	Name          string `toml:"name"`
	OnEntryAction string `toml:"onEntry"`
}

// TomlTransition represents a transition in the FSM. A transitions defines
// the name of the event and which state the FSM should move. If the transition
// has a condition, the transition will be performed if the condition
// evaluates true.
type TomlTransition struct {
	Event     string `toml:"event"`
	StateFrom string `toml:"stateFrom"`
	StateTo   string `toml:"stateTo"`
	Condition string `toml:"condition,omitempty"`
}

// TomlFSM represents the FSM described on a TOML document.
type TomlFSM struct {
	Name         string `toml:"name"`
	Description  string `toml:"description"`
	Package      string `toml:"package"`
	InitialState string `toml:"initialState"`
	Timestamp    time.Time
	States       []TomlState      `toml:"states"`
	Transitions  []TomlTransition `toml:"transitions"`
}

// ParseFsmFile takes a Reader interface and tries to parse a TOML document
// into a TomlFSM struct.
func ParseFsmFile(fsmFile io.Reader) (*TomlFSM, error) {
	buffer, err := ioutil.ReadAll(fsmFile)
	if err != nil {
		return nil, errors.New("Could not read input file")
	}

	fsm := TomlFSM{Timestamp: time.Now().UTC()}

	err = toml.Unmarshal(buffer, &fsm)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("Could not parse file")
	}

	return &fsm, nil
}

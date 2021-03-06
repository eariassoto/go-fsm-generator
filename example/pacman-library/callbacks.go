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
	"fmt"
)

// Conditionals

func pacmanClose(data interface{}) bool {
	fmt.Println("Conditional pacmanClose called")
	return true
}

// State callbacks

func chase(data interface{}) {
	ghostStruct, ok := data.(*Ghost)
	if ok {
		ghostStruct.CurrentActivity = "Chasing Pacman"
		fmt.Println(ghostStruct.CurrentActivity)
	}
}

func run(data interface{}) {
	ghostStruct, ok := data.(*Ghost)
	if ok {
		ghostStruct.CurrentActivity = "Running from Pacman"
		fmt.Println(ghostStruct.CurrentActivity)
	}
}

func walkAround(data interface{}) {
	ghostStruct, ok := data.(*Ghost)
	if ok {
		ghostStruct.CurrentActivity = "Walking around the maze"
		fmt.Println(ghostStruct.CurrentActivity)
	}
}

func reset(data interface{}) {
	ghostStruct, ok := data.(*Ghost)
	if ok {
		ghostStruct.CurrentActivity = "Eaten by Pacman"
		fmt.Println(ghostStruct.CurrentActivity)
	}
}

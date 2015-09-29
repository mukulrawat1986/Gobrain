// This is a simple implementation of a Brainfuck interpreter in Golang

package main

import (
    _"bufio"
    "fmt"
    _"log"
    "os"
)

// Here we define our tape data structure
// Tape is a slice of 8-bit unsigned integers. 
// We will be getting ASCII values from the user as input, uint8 
// are perfect to store them.
// We also have a pointer that points to the current cell in the tape.
// We can increment or decrement this pointer.
type data struct{
    tape []uint8
    ptr int
}

// Now here we will define some methods on our data strucutre.

// Initialize our tape data structure
func (d *data) initialize(){
    d.tape = make([]uint8, 30000)
    d.ptr = 0
}

// Increment the pointer
func (d *data) increment_ptr(){
    d.ptr += 1
}

// Decrement the pointer
func (d *data) decrement_ptr(){
    d.ptr -= 1
}

// Increment the value in current cell
func (d *data) increment_value(){
    d.tape[d.ptr] += 1
}

// Decrement the value in current cell
func (d *data) decrement_value(){
    d.tape[d.ptr] -= 1
}

// Read input from stdin
// The tape reads only one character at a time
func (d *data) read () {
    // we will read only one ASCII character from the stdin
    // we will create a separate bye to store the character in
    r := make([]uint8, 1)
     _, _ = os.Stdin.Read(r)
    d.tape[d.ptr] = r[0]
}


// Print out the ascii character to screen
func (d *data) print(){
    fmt.Printf("%s",string(uint8(d.tape[d.ptr])))
}


func main(){
    var  d data
    d.initialize()
    fmt.Printf("%d %d\n", d.tape[0], d.ptr)
}


// This is a simple implementation of a Brainfuck interpreter in Golang

package main

import (
    _"bufio"
    "fmt"
    "io/ioutil"
    "log"
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
    d.tape = make([]byte, 30000)
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
    // we will read only one ASCII character from the file
    // we will create a separate bye to store the character in
    r := make([]byte, 1)
     _, _ = os.Stdin.Read(r)
    d.tape[d.ptr] = r[0]
}


// Print out the ascii character to screen
func (d *data) print(){
    //fmt.Printf("%s",string(uint8(d.tape[d.ptr])))
    fmt.Printf("%c\n",d.tape[d.ptr])
}


func main(){

    if len(os.Args) != 2{
        log.Printf("Please enter the file name")
        os.Exit(1)
    }

    filename := os.Args[1]

    fmt.Printf("Filename is %v\n", filename)

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	source := string(buf)

    // set our type for our machine
    var d data

    // initialize it
    d.initialize()

    for i := 1; i < len(source); i++{
        
        fmt.Printf("%v\n", d.tape[d.ptr])

        switch source[i]{
        
        case '<':
            d.decrement_ptr()
        
        case '>':
            d.increment_ptr()
       
        case '+':
            d.increment_value()
      
        case '-':
            d.decrement_value()
        
        case '.':
            d.print()

        case ',':
            d.read()
        
        case '[':
            if d.tape[d.ptr] == 0{
                for j := i; j < len(source); j++{
                    if source[j]  == ']'{
                        i = j
                        break
                    }
                    if j == len(source) - 1 {
                        log.Printf("Braces missing in the loop\n")
                        os.Exit(1)
                    }
                }
            }

        case ']':
            if d.tape[d.ptr] != 0{
                for j := i; j > 0; j-- {
                    if source[j] == '['{
                        i = j
                        break
                    }
                    if j == 0{
                        log.Printf("Braces missing in the loop\n")
                    }
                }
            }
        }
    }
}

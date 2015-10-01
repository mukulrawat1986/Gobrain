// This is a simple implementation of a Brainfuck interpreter in Golang

package main

import (
    "bufio"
    "fmt"
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

    if len(os.Args) != 2{
        log.Printf("Please enter the file name")
        os.Exit(1)
    }

    filename := os.Args[1]

    fmt.Printf("Filename is %v\n", filename)

    // open file
    f, err := os.Open(filename)

    if err != nil{
        log.Printf("Error opening file")
        os.Exit(1)
    }

    // close the file
    defer f.Close()

    // scan the file
    scan := bufio.NewScanner(f)

    // set our type for our machine
    var d data

    // initialize it
    d.initialize()

    for scan.Scan(){

        // we scan one line from the file
        l := scan.Text()

        // Now we loop over every character in the file
        for _, x := range l{

            switch l{

            case '>':
                d.increment_ptr()

            case '<':
                d.decrement_ptr()

            case '+':
                d.increment_value()

            case '-':
                d.decrement_value()

            case '.':
                d.print()

            case ',':
                d.read()

            case '[':


            case ']':
            }

        }
    }

    if scan.Err() != nil{
        log.Printf("Error while scanning ", scan.Err())
        os.Exit(1)
    }

}


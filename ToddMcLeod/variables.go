package main

import "fmt"

var s string

func main() {

	a := 4

	fmt.Println(a)

	b, c, _, d := 2, "hello", true, 50.2

	fmt.Println(b, c, d)

	fmt.Println(s)
	s = "Hey"

	fmt.Println(s)

	fmt.Printf("%v \t %T \t %x \n ", s, s, s)
}

//Zero Values for variables

//var <var name> <var type>
//eg. var a string

/*
boolean - false
int     - 0
float   - 0.0
string  - ""
pointers   - nil
functions  - nil
interfaces - nil
slices     - nil
channels   - nil
maps       - nil

*/

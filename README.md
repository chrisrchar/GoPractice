# Go Class Code
This repo holds code from a Udemy class on learning the syntax and structures of GoLang. The class can be found [here](https://www.udemy.com/course/go-the-complete-developers-guide/).

# Projects
## Cards
The first project in the class meant to familiarize with Go's syntax and how functions work inside packages. Created a custom **deck** type with functions to simulate a card deck. Also used Go's io package to read and write the deck slice to local disk

## Even and Odd
The first assignment in the class. Simple code that iterates through an array to determine if the number in the array is even or odd then prints the result to the console.

## Structs
Code introducing structs and also writing functions with pointers. Includes embedded structs and a function to update a struct.

## Map
Code introducing maps and how to use them. Also showing how they differ from structs based on strict typing and being pass by reference.

## Interface & http
Introduction to interfaces in Go. Used as a simple form of abstraction to make sure specific types implement specific functions. Interfaces was a simple example of two chat bot types that are both a bot type. The http file utilized the http package with more depth into the Reader and Writer interfaces.

## Assignment 2 & 3
Assignments 2 and 3 were more practice with interfaces. Assignment 2 was practice in writing an interface with shapes that implement a `getArea()` function and then the `printArea()` function takes the more generic `shape` type interface as an argument.

Assignment 3 was more so practice with the Go documentation and using packages to read in an external file, then print it out to the terminal.

## Channels
Coded a website status checker to introduce go routines/asynchronous functions in Go. Also used a function literal to delay routines. 
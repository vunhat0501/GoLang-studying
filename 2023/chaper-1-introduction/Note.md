# Chapter 1

## Lesson 3: Go is fast and simple

Go code generally runs faster than interpreted languages and compiles faster than other compiled languages like C and Rust.

## Lesson 4: Comparing Go's speed

Go is _generally_ faster and more lightweight than interpreted or VM-powered languages like:

- Python
- JavaScript
- PHP
- Ruby
- Java

However, in terms of execution speed, Go does lag behind some other compiled languages like:

- C
- C++
- Java

Go is a bit slower mostly due to its automated memory management, also known as "Go runtime". Slightly slower speed is the price we pay for memory safety and simple syntax!

## Lesson 5: The compilation process

Computers need machine code, they don't understand English or even Go. We need to convert our high-level (Go) code into machine code, which is really just a set of instructions that some specific hardware can understand. In your case, your CPU.

The Go compiler's job is to take Go code and produce machine code. On Windows, that would be a **_.exe_** file. On Mac and Linux, it would be any executable file.

### A note on the structure of a Go program

1. **_package main_** lets the Go compiler know that we want this code to compile and run as a standalone program, as opposed to being a library that's imported by other programs.
2. **_import fmt_** imports the **_fmt_** (formatting) package. The formatting package exists in Go's standard library and let's us do things like print text to the console.
3. **_func main()_** defines the **_main_** function. **_main_** is the name of the function that acts as the entry point for a Go program.

## Lesson 6: Compiling explained

Computer don't know how to do anything unless we as programmers tell them what to do.

Unfortunately computers don't understand human language. In fact, they don't even understand un-compiled computer programs.

For example, the code:

    package main

    import "fmt"

    func main() {
        fmt.Println("Hello, World!")
    }

means _nothing_ to a computer.

### Computer need machine code

A computer's CPU only understands its own _instruction set_, which we can call "machine code".

Instructions are basic math operations like addition, subtraction, multiplication, and the ability to save data temporarily.

For example, an ARM processor uses the _ADD_ instruction when supplied with the number 0100 in binary.

### Go, C, Rust, etc.

Go, C, and Rust are all languages where the code first converted to machine code by the compiler before executed.

## Lesson 7-8-9: Compiled vs Interpreted

Compiled programs can be run without access to the original source code, and without access to a compiler.

Note how this is different than interpreted language like Python and JavaScript. With Python and JavaScript the code is interpreted at runtime by a separate program known as the "interpreter". Distributing code for users to run can be a pain because they need to have an interpreter installed, and they need access to the original source code.

## Lesson 10: Go is strongly typed

Go enforces strong and **static** typing, meaning variable can only have a single type. A **string** variable like "hello world" can not be changed to an **int**, such as number 3.

One of the biggest benefits of strong typing is that errors can be caught at "compile time". In other words, bugs are more easily caught ahead of time because they are detected when the code is compiled before even runs.

Contrast with most interpreted languages, where the variable types are dynamic. Dynamic typing can lead to subtle bugs that are hard to detect. With interpreted languages, code _must_ be run (sometimes in production if unlucky) to catch syntax and type errors.

### Concatenating strings

Two strings can be [concatenated](https://www.w3schools.com/python/python_strings_concatenate.asp) with the **+** operator. Because Go is strongly typed, it won't allow you to concatenate a string variable with a numeric variable.

### Assignment

Using simple basic authentication for the Textio API (see in chapter-1-introduction/lesson10). This is how our users will communicate to us who they are and how many features they are paying for with each request to our API.

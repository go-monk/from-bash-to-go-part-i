This is the first part of a series introducing Bash programmers to Go. This part goes through the language building blocks that will be used in later parts.

# Ok but why?

> A language that doesn't affect the way you think about programming is not worth knowing. - Alan J. Perlis

You might be wondering along these lines - I already kind of know Bash (or a similar language) and can do all I need in it. It's easy and fast. Why should I learn Go? That's a good question. Always ask why because the answer to a why question provides a reason and thus understanding and motivation. In our case I think the answer has to do with the difference between programming and software engineering.

Programming means writing a program that works and does something useful. Software engineering is programming plus time and other people. It's the initial writing of a program and its being modified by you or other people over time. Programming alone is hard enough. First you need to understand the domain and the concrete problem to solve within the domain. Then you design a solution and implement it in a programming language whose syntax and idioms you should know well. This process can, and *should*, take multiple iterations. When you are done you go work on other stuff. Then you might be asked to modify something in the program (to fix a bug or add new functionality) or to hand over the program to someone else (people come and go).

The most important thing to do when doing the software engineering is to reduce the cognitive load; to reduce the system's complexity. This requires hard work, attention to detail and using good tools. I think Go is a good tool for software engineering because it includes "a cultural agenda of radical simplicity". See https://github.com/go-monk/from-bash-to-go for a practical example of how and why migrate a script from Bash to Go.

# Building blocks

In this section I swiftly introduce some of the language building blocks that I hope will help you start understanding the Go syntax, semantics and idioms. I recommend actually writing (copying) the code below in your favorite editor. And then running it. And maybe changing it a bit and running again. If you break the code be happy, that's a way to learn :-).

## Writing and running Go code

Packages are Go's way of organizing and (re)using code.

Bash is organized mostly via files - each program usually lives in a file:

```
+------------+
| script1.sh |
+------------+
+------------+
| script2.sh |
+------------+
+------------+
| script3.sh |
+------------+
```

Go code lives in one or more *packages* that are contained in one or more .go files within a single directory. Packages can be grouped into *modules* for versioning and sharing.

It can be visualized like this:

```
+--------------------------+
| module example.net/hello |
|                          |
|  +-------------------+   |
|  |   package main    |   |
|  |                   |   |
|  |  +-------------+  |   |
|  |  | greeting.go |  |   |
|  |  +-------------+  |   |
|  |  +-------------+  |   |
|  |  | hello.go    |  |   |
|  |  +-------------+  |   |
|  |                   |   |
|  +-------------------+   |
|                          |
+--------------------------+
```

Go identifiers - constants, variables, types and functions - are visible (exported) outside of a package when their name starts with an uppercase letter. Otherwise they are confined to the current package.

Let's create our first package. In case you want to run your code (as opposed to using it as an importable library) you need at least the `main` package.

First create a directory and change to it:

```
$ mkdir hello
$ cd hello
```

Then create `hello.go` file with the following content:

```go
package main

import "fmt"

func main() {
	fmt.Println("hello")
}
```

The `main` function is where the program's execution starts.

The easiest way to run a Go program is:

```sh
$ go run hello.go # build the binary and run it
hello
```

As mentioned above, you can spread package code into multiple files within the same directory:

```go
// hello.go
package main

import "fmt"

func main() {
        fmt.Println(greeting)
}
```

```go
// greeting.go
package main

const greeting = "hello"
```

Now we need to include both package files:

```sh
$ go run hello.go greeting.go
hello
```

Module is a group of packages that is versioned as a unit. To create a module:

```sh
$ go mod init github.com/jsmith/hello
$ go mod tidy # download dependencies
```

To build for a different OS and/or CPU architecture than the one you are running:

```sh
macOS$ GOOS=linux GOARCH=amd64 go build
```

To see the list of all supported OS/ARCH combinations:

```sh
$ go tool dist list
```

See https://go.dev/doc/tutorial/getting-started for more.

## Variables and types

In Bash all simple variables are strings:

```sh
name=Jack
age=40
active=true

# this is not a problem in Bash, since there are no types
age=forty
```

Go is a statically typed language. It means that every variable has a type and the type cannot change during program's run:

```go
// the := operator infers the type from the value
name := "Jack"  // string
age := 40       // int
active := true  // bool

// compile-time error: cannot use "forty" 
// (untyped string constant) as int value in assignment
age = "forty"
```

See https://go.dev/tour/basics/11 for all basic types.

Variables declared without an explicit value are given their *zero* value:

* `""` (the empty string) for strings
* `0` for numeric types
* `false` for boolean types

```go
var i int
var f float64
var b bool
var s string
fmt.Printf("%v %v %v %q\n", i, f, b, s) // 0 0 false ""
```

Sometimes you might need to convert a type:

```go
name := "Jack"
age := "40"
// Convert string to slice of bytes or runes
nameRunes := []rune(name) // when you care about UTF-8 encoded characters
nameBytes := []byte(name) // when you care about raw data (I/O, network, crypto, performance)
// Convert string to an int
ageInt, _ := strconv.Atoi(age) // NOTE: ignoring error for brevity
```

A scope is a part of the program in which a variable can be seen. While in Bash variables are often global, in Go we have following scopes:

* package scope - when declared outside a function a variable can be seen by the entire package
* function scope - when declared within a function it can be seen only within function's `{}`
* statement scope - can be seen within `{}` of a statement (for loop, if/else)

Go's type system and scoping rules prevent many common errors that can occur in Bash scripts, especially as they grow larger and more complex and are maintained over time.

## Slices and maps

Together with basic (data) types - like numbers, strings and booleans above - there are other types that can hold multiple pieces of data: arrays, structs, slices and maps. Arrays (underlying the slices) are seldom used directly. We'll cover structs later.

A slice is a dynamically-sized (array is a fixed-sized) group of elements of certain type. Here are some common [slice operations](https://go.dev/play/p/c_n7e76AGue):

```go
// Create (declare and initialize) a slice of integers
primes := []int{2, 3, 5, 7, 11}

// Append an element
primes = append(primes, 13)

// Print first and last element
fmt.Println(primes[0], primes[len(primes)-1]) // 2 13

// Slice a slice
s := primes[1:4] // s == [3 5 7]

// Loop over all elements (ignoring indices by using _)
for _, p := range primes {
	fmt.Printf("%d ", p)
}
```

To learn more about slices see https://go.dev/blog/slices-intro.

A map (called associative array, dictionary or hash in other languages) maps keys to values. Some common [operations on maps](https://go.dev/play/p/7cSaQIDlhXV):

```go
// Create a map of strings to integers
m := map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
}

// Insert or update an element
m["four"] = 4

// Retrieve an element
elem := m["one"] // elem == 1

// Delete an element
delete(m, "four")

// Test that a key is present with a two-value assignment
elem, ok := m["two"] // elem == 2, ok == true
elem, ok = m["four"] // elem == 0, ok == false

// Loop over all elements getting key and value for each
for k, v := range m {
	fmt.Printf("%s -> %d\n", k, v)
}
```

To learn more about maps visit https://go.dev/blog/maps.

## Loops

Most languages have multiple statements for looping (for, while, do while). In Go there's only a `for` loop but it can implement the functionality of all the loop types. Here are some examples:

```go
// Traditional C-style loop.
for i := 0; i < 10; i++ {
    fmt.Print(i) // 0123456789
}

// While loop.
var i int
for i < 10 {
    fmt.Printf("i is less than 10 (%d)\n", i)
    i++
}
fmt.Printf("i is 10\n")

// Infinite loop.
for {
    fmt.Println("To Infinity and Beyond.")
    time.Sleep(time.Second)
}
```

You might occasionally need the `break` or `continue` statement to control the loop:

```go
for {
    if err := doSomething(); err != nil {
        break // break out of the loop
    }
    fmt.Println("keep doing something")
}

for i := 0; i < 10; i++ {
    if i % 2 == 0 {
        continue // continue with the next loop
    }
    fmt.Printf("that's an odd number: %d\n", i)
}
```

## Functions

Functions in Bash and Go serve similar purposes but work quite differently. Bash functions are more like commands that often operate on global state:

```sh
# Function definition
greet() {
    echo "Hello, $1!"
}

# Function call
greet "World"  # Prints: Hello, World!
``` 

Go functions are first-class values (can be stored in variables, passed as function arguments or return values) with explicit parameters and return values:

```go
package main

import "fmt"

// Function with parameter named who of type string and 
// a return value (without name) of type string.
func greet(who string) string {
	return fmt.Sprintf("Hello, %s!", who)
}

func main() {
    message := greet("World")
    fmt.Println(message)  // Hello, World!
}
```

Go functions can return multiple values (need parenthesis). The second value is often an error:

```go
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	a, b := 10, 2
	res, err := divide(a, b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while dividing: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%d : %d = %d\n", a, b, res)
}
```

See https://github.com/go-monk/error-handling to learn more about errors and how to handle them in Go.

## Pointers

Variables are just names (or aliases) for memory addresses. Memory addresses are large numbers usually displayed in hexadecimal notation. So for instance `var i int = 1` means that we store value `1` of type `int` in a place named `i` that's somewhere in the memory. To find out the memory address we take a pointer to the variable using the `&` operator:

```go
var i int = 1 // or just i := 1 within a function
println(&i)   // large number like 0x14000058730
```

We can also store the pointer in another variable instead of just printing it:

```go
var i int = 1
var ip *int = &i
```

But how do we access the original value not just the memory address? We dereference it using the `*` operator:

```go
fmt.Println(ip)  // 0x14000058730 (memory address)
fmt.Println(*ip) // 1 (value stored at the memory address)
```

We can see that `*` is used in two contexts. First we use it to indicate that ip holds a pointer to int (`var ip *int`). Then we use it to dereference the pointer to get to the original value (`*ip`).

Before dereferencing a pointer you must make sure it's not nil:

```go
var p *int
fmt.Println(p == nil) // true
fmt.Println(*p)       // panics
```

Pointers are generally useful for two things:
* If we pass some data to a function, the parameter (or a receiver - see methods below) that we use inside the function holds a copy of the data. If we want to modify the data we need to pass a pointer.
* If data is large and/or we need very high efficiency we don't have to copy the data we can just point to it (this is similar to filesystem's symlinks).

If you haven't worked with pointers before and/or my explanation makes no sense, please have a look at one or more of these:
* https://tour.golang.org/moretypes/1
* https://yourbasic.org/golang/pointers-explained
* https://dave.cheney.net/2017/04/26/understand-go-pointers-in-less-than-800-words-or-your-money-back

## Structs and methods

Bash does not have native support for structs or methods. The closest you can get is using associative arrays (called maps, hashes, or dictionaries in other languages) to group related data, but there is no way to attach behavior (methods) to them:

```bash
# Define a "person" using an associative array
declare -A person
person[name]="Alice"
person[age]=30

# Access fields
echo "Name: ${person[name]}"
echo "Age: ${person[age]}"
```

To group data in Go we create our own type called `Person` that is based on a `struct`. Struct is an aggregate data type that has zero or more fields. The fields don't have to have the same type:

```go
package main

import "fmt"

// Define our type
type Person struct {
    Name string
    Age  int
}

func main() {
    // Create and use a variable of our type
    p := Person{Name: "Alice", Age: 30}
    fmt.Printf("Name: %s\n", p.Name)
    fmt.Printf("Age: %d\n", p.Age)
}
```

To attach behavior to data we take a normal function and prefix its name with the type (usually a struct) we want to attach it to. This "prefix" is called a receiver. Receiver, in the same way as a function parameter, holds copy of the data. In case you want to modify the data you need to use a pointer receiver. To avoid confusion it's best to have all receivers either pointers or non-pointers (values), instead of mixing them as in the example below:

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Method is just a function with a receiver (named p and of type Person).
// The receiver binds the function to some data.
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// This method has a pointer receiver (of type *Person - pointer to Person).
func (p *Person) IncreaseAge(years int) {
	p.Age += years
}

func main() {
	p := Person{Name: "Bob", Age: 40}
	p.Greet() // Hello, my name is Bob and I am 40 years old.
	p.IncreaseAge(5)
	p.Greet() // Hello, my name is Bob and I am 45 years old.
}
```

## Interfaces

Bash does not have interfaces. In Go, an interface is an abstract type that describes some behavior by using a set of function signatures. An interface is implemented by any type that has the whole set of methods (remember, method is just a function attached to a type via receiver).

What does it mean? Let's look at an interface example called `Stringer`. It's defined in the `fmt` package that's part of the Go standard library:
```
$ go doc fmt.Stringer
package fmt // import "fmt"

type Stringer interface {
        String() string
}
    Stringer is implemented by any value that has a String method, which defines
    the “native” format for that value. The String method is used to print
    values passed as an operand to any format that accepts a string or to an
    unformatted printer such as Print.
```

Imagine we have a program with the following type:

```go
type Backup struct {
	Desc     string
	Size     int64 // bytes
	LastDone time.Time
}
```

And we want to print out this type nicely. Here's a first attempt:

```go
bak := Backup{
    Desc:     "personal code",
    Size:     1024,
    LastDone: time.Now(),
}
fmt.Println(bak) // {personal code 1024 2025-07-09 16:42:36.189568 +0200 CEST m=+0.000201918}
```

That's not too bad but we can do better. We read in the `fmt.Stringer` documentation that `fmt.Print` accepts a Stringer. We can make our `Backup` type a `Stringer` by attaching a method with `String() string` signature to it:

```go
func (b Backup) String() string {
	return fmt.Sprintf(
		"backup of %s (%d bytes) was last done on %s",
		b.Desc, b.Size, b.LastDone.Format(time.DateTime))
}
```

Now when we pass it as an argument to `fmt.Print` its `String` method is called and its output is printed:

```
backup of personal code (1024 bytes) was last done on 2025-07-09 16:47:49
```

Another standard library interface is the [io.Writer](https://pkg.go.dev/io#Writer). Any (concrete) type that has the method:

```go
Write(p []byte) (n int, err error)
```

implements `io.Writer` (or we say is a Writer) and thus can be used as the first argument of the `fmt.Fprint` function:

```go
package fmt // import "fmt"

func Fprint(w io.Writer, a ...any) (n int, err error)
```

There are several types in the standard library that implement the `io.Writer` interface, like `os.Stdout`, `os.Stderr` or `net.Conn`:

```go
fmt.Fprint(os.Stdout, "hello") // prints to STDOUT
fmt.Fprint(os.Stderr, "error") // prints to STDERR

conn, _ := net.Dial("tcp", "example.com:80") // NOTE: ignoring error for brevity
fmt.Fprint(conn, "GET HTTP/1.0")             // prints to a network connection
```

As an exercise try to implement a simple netcat-like read-only TCP client using [io.Copy](https://pkg.go.dev/io#Copy). You can have a look at https://github.com/gokatas/netcat for inspiration.

To learn more about reading data in Go see https://github.com/go-monk/reading-data.

## Testing

Let's have some fun now, since we are getting tired ... We craft a package `word` that has a function telling us whether a word is a palindrome:

```go
// ./word/1/word.go
// IsPalindrome reports whether s reads the same forward and backward.
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
```

We range over a string comparing its edge elements. We start with `i == 0` thus comparing the first element (at index 0) with the last element (at index that is one less than the string size). In the second iteration we compare the second element with the second to last. And so on. If all are the same, we have a palindrome! Nice and easy.

But since we know now that we should be doing software engineering instead of programming, or as John Osterhout writes in "A Philosophy of Software Design" strategic programming instead of tactical programming, we take the effort of writing a test for our function:

```go
// ./word/1/word_test.go
func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") == false`)
	}
}
```

Let's see:

```sh
$ go test
PASS
ok      word    0.381s
```

Sweet, satisfied we go for a coffee ... When we come back, we find a Slack message from our Slovak colleague complaining about our new package. He says that it doesn't recognize the word ťahať as a palindrome. Really? We turn this complaint into a test case:

```go
// ./word/2/word_test.go
func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("ťahať") {
		t.Error(`IsPalindrome("ťahať") == false`)
	}
}
```

When we run the test it fails, so our colleague is right. Now, we can go for the easy option and comment on the function that the input must be an ASCII sequence. But since we have already decided for strategic programming we must take the difficult path. After some research we find out [how Go strings really work](https://go.dev/blog/strings). So strings are just (read-only) sequences of bytes. *Any* bytes. A string is not required to hold Unicode text, UTF-8 text, or any other predefined format. Therefore let's first convert the string to a rune slice:

```go
// ./word/2/word.go
func IsPalindrome(s string) bool {
	runes := []rune(s)
	for i := range runes {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}
```

Now the tests pass again. Phew...

It didn't take long and we've got another bug report. Some smartass came up with this cool sentence: `A man, a plan, a canal: Panama`. Our current implementation of `IsPalindrome` thinks it's not a palindrome. Ok, first let's improve our tests, we'll use something called table-driven testing:

```go
// ./word/3/word_test.go
func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"kayak", true},
		{"ťahať", true},
		{"A man, a plan, a canal: Panama", true},
		{"", true},
		{"ab", false},
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}
```

Now, after some head scratching we realize the problem is we don't ignore whitespace, punctuation and letter case. Let's fix that:

```go
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
```

## Goroutines and channels

Concurrency is a way of structuring a program such that multiple functions can execute non-sequentially. If there are multiple CPUs on the machine (which is very likely today) they get executed in parallel. Goroutines and channels are for some people one of the most interesting features of Go since they make concurrent programming relatively easy. Nevertheless concurrent programming is still inherently more complex for our brains than sequential. Use concurrency only when it really makes sense.

Here's a function that calculates `n`th number from Fibonacci sequence. It's a recursive function, i.e. a function that keeps calling itself until `n < 2` (you always need such a condition in a recursive function otherwise the function will keep calling itself forever and the program will crash with "stack overflow" error):

```go
// ./fib/1/main.go
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-2) + fib(n-1)
}
```

It's a slow algorithm; to get the 45th Fibonacci number it takes some time:

```sh
$ time go run ./fib/1/main.go 
fib(45) = 1134903170

real    0m4.121s
user    0m3.614s
sys     0m0.130s
```

Wouldn't it be nice to know the program is still running while waiting for it to finish? Here's a little spinner with the rotation speed defined by the `delay` parameter:

```go
// ./fib/2/main.go
func spinner(delay time.Duration) {
	for {
		for _, r := range `\|/-` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
```

But here's the problem: if we call the `spinner` function before the `fib` function the `fib` function will be called only after the `spinner` function returned (finished). Of course, the `spinner` function will never return; the outer for loop is infinite.

Enters `go` keyword. When you prefix a function with `go` it causes the function to be called in a newly created goroutine and the `go` statement completes immediately. It's similar to appending `&` to a Bash command thus sending it to the background and freeing the terminal for use. 

```go
f()		// call f(); wait for it to return
go f()	// create a new goroutine that calls f(); don't wait
```

So this is our solution of the problem:

```go
// ./fib/2/main.go 
func main() {
	go spinner(time.Millisecond * 100)
	fmt.Printf("\rfib(45) = %d\n", fib(45))
}
```

That's nice. But notice that once we run the `spinner` in a new goroutine we have no way to communicate with it or stop it. It works in this case because the spinner writes stuff on the terminal and gets stopped when the whole program terminates after calculating and printing the Fibonacci number.

To solve these synchronization and communication problems between goroutines Go provides channels. They are similar to the shell pipes but are typed. This is how one works with a channel:

```go
c := make(chan int) // declaring and initializing
c <- 1 				// sending on a channel
value := <-c 		// receiving from a channel
```

The "arrow" indicates the direction of data flow.

Now, let's modify our `fib` function. It will not calculate Nth number from the Fibonacci sequence but it will keep producing the numbers from the sequence:

```go
// ./fib/3/main.go
func fib() <-chan int {
	c := make(chan int)
	go func() {
		a, b := 0, 1
		for {
			c <- a
			a, b = b, a+b
		}
	}()
	return c
}
```

First we create a channel of integers called `c`. Next we start a goroutine which executes an (anonymous) function that keeps sending the Fibonacci numbers on the channel. And we return the channel as a receive-only (`<-`).

In the main function we take out the first 46 numbers from the channel and print them in sequence:

```go
// ./fib/3/main.go
c := fib()
for i := 0; i <= 45; i++ {
	fmt.Printf("fib(%d) = %d\n", i, <-c)
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}
```

The communication between the `fib` and `main` function (`main` is also being executing on a goroutine) happens by sending and receiving on a channel. What's more subtle is the synchronization part. Since un-buffered channels, like the one we used, can hold only one element, the `fib` function blocks (stops executing) on the `c <- a` line until the main function picks up (`<- c`) the value from the channel. I added the `time.Sleep` to emphasize the fact that it's the `main` function that blocks and unblocks the goroutine launched by the `fib` function. 

To learn more about concurrency directly from one of the Go creators watch [this video](https://youtu.be/f6kdp27TYZs).

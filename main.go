package main

// This calls a JS function from Go.
func main() {
	println("adding two numbers using JavaScript add():", add(2, 3)) // expecting 5
}

// This function is imported from JavaScript, as it doesn't define a body.
// You should define a function named 'main.add' in the WebAssembly 'env'
// module from JavaScript.
func add(x, y int) int

// This function is exported to JavaScript, so can be called using
// exports.multiply() in JavaScript.
//export multiply
func multiply(x, y int) int {
	return x * y
}

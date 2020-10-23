package main

func main() {
}

// This function is exported to JavaScript, so can be called using
// exports.add() in JavaScript.
//export add
func add(x, y int) int {
	return x + y
}

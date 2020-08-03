package logger

import "fmt"

func ExampleLevel_String() {
	level := DebugLevel

	fmt.Println(level.String())
	// Output:
	// debug
}
func ExampleLevel_CapitalString() {

	level := DebugLevel

	fmt.Println(level.CapitalString())
	// Output:
	// DEBUG
}

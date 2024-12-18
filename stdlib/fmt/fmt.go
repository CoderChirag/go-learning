package fmt

import (
	"fmt"
	"os"
)

// Prints to stdout without formatting, still formats escapes characters
// func Print(a ...any) (int, error)
// func Println(a ...any) (int, error)
// Starting with 'S': Returns string instead of printing
// func Sprint(a ...any) string
// func Sprintln(a ...any) string
// Starting with 'F': Takes a writer to print
// func Fprint(w io.Writer, a ...any) (int, error)
// func Fprintln(w io.writer, a ...any) (int, error)
// Ending with 'f': formats and prints
// func Printf(format string, a ...any) (int, error)
// func Sprintf(format string, a ...any) string
// func Fprintf(w io.writer, format string, a ...any) (int, error)
func FmtExamples() {
	// built in functions
	name := "Chirag"
	print(name + " ")
	print(name + "\n")
	println(name)
	fmt.Print("Rock ")
	fmt.Print("Tree\n")
	fmt.Println("Sky")
	fmt.Println(44)
	fmt.Println(false)
	pet1, pet2, pet3 := "dog", "cat", "fish"
	fmt.Print("pets: ", pet1, ", ", pet2, ", ", pet3, "\n")
	n, err := fmt.Println("Soil")
	fmt.Println("n:", n, "err:", err)
	a := fmt.Sprint("car")
	fmt.Print(a, "\n")
	fmt.Fprint(os.Stdout, "Hello World\n")
}
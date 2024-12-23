package errors

import (
	"errors"
	"fmt"
	"os"
)
var _ error = (*CustomError)(nil)
type CustomError struct {
	msg string
	err error
	name string
}
func (e *CustomError) Error() string {
	return  e.name + ": " + e.msg + ": " + fmt.Sprint(e.err)
}
func (e *CustomError) Unwrap() error {
	return e.err
}
func readFile() (string, error) {
	content, err := os.ReadFile("a.go")
	if err != nil {
		e := &CustomError{err: err, msg: "File Not Found", name: "FileNotFoundError"}
		return "", e
	}
	return string(content), nil
}
func ErrorsExamples(){
	content, err := readFile()
	
	if err != nil {
		fmt.Println(err)
		if(errors.Is(err, os.ErrNotExist)){
			fmt.Println("os.ErrorNotExist occurred.")
		}
		return
	}
	fmt.Println(content)
}
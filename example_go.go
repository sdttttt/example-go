package example

import "github.com/codecov/example-go/awesome"

var result string

func Setup() {

	// Comment

	result = awesome.Smile()

}

func pp() {

	// Comment
	println(1)
}

func GetResult() string {

	/*
	   Comment
	*/

	return result

}

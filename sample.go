package workflowdemo

import "fmt"

type Foo int

const (
	FooFirst Foo = iota
	FooSecond
	FooThird
	FooFourth
)

func Try(f Foo) {
	switch f {
	case FooFirst:
		fmt.Println("this is the first case")
		fmt.Println("this is the first case")
		fmt.Println("this is the first case")
		fmt.Println("this is the first case")
		fmt.Println("this is the first case")
		fmt.Println("this is the first case")
		fmt.Println("this is the first case")
		fmt.Println("this is the first case")
	case FooSecond:
		fmt.Println("this is the second case")
		fmt.Println("this is the second case")
		fmt.Println("this is the second case")
		fmt.Println("this is the second case")
		fmt.Println("this is the second case")
		fmt.Println("this is the second case")
		fmt.Println("this is the second case")
		fmt.Println("this is the second case")
	case FooThird:
		fmt.Println("this is the third case")
		fmt.Println("this is the third case")
		fmt.Println("this is the third case")
		fmt.Println("this is the third case")
		fmt.Println("this is the third case")
		fmt.Println("this is the third case")
		fmt.Println("this is the third case")
		fmt.Println("this is the third case")
	}
}

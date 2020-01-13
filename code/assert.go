package main

import "fmt"

type Animal interface {
	print()
	// climb()
}

type Dog struct {
	name int
}

func (dog *Dog) print() {
	fmt.Println(dog.name)
}

type stringer string

func main() {
	var s interface{}
	s = stringer("abc")
	if _, ok := s.(string); ok {
		fmt.Println("s is string")
	} else {
		fmt.Println("s is not string")
	}

	// str := "123"
	// var _ int = int(str)	//用于test文件
	var dog interface{}
	dog = &Dog{name: 1}
	fmt.Printf("%#v\n", dog)
	var dog2 Dog
	dog2 = Dog{name: 2}
	dog2.print()
	if u, ok := dog.(Animal); ok {
		fmt.Printf("%#v\n", u)
		u.print()
	} else {
		fmt.Println("Dog is not satisfize Animal")
	}

	var dog3 Animal
	dog3 = &Dog{name: 3}
	dog3.print()

}

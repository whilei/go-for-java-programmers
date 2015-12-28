package oo
import "fmt"

type Animal interface{
	Move()
}

type Animal2 interface{
	Move()
	Speak()
}

type Dog struct {
	//Speak func()
}

func (this *Dog) Move()  {
	fmt.Println("dog move")
}

package oo

import "fmt"

type Animal interface {
	Move()
}

type Animal2 interface {
	Move()
	Speak()
}

type Dog struct {
	//Speak func()
	p string
}

func NewDog() *Dog {
	return &Dog{p: "1"}
}

func (this *Dog) Move() {
	fmt.Println("dog move " + this.p)
}

type BlankDog struct {
	*Dog
}

func (this *BlankDog) Move() {
	this.Dog.Move()
	fmt.Println("blank dog move " + this.p)
}

func NewBlankDog() *BlankDog {
	return &BlankDog{NewDog()}
}

type WhiteDog Dog

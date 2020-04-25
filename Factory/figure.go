package main

import "fmt"

const (
	CIRCLE = "circle"
	SQUARE = "square"
)

type Figure interface {
	Draw()
}

type Square struct{
}

func (s Square) Draw(){
	fmt.Println("heyyy am a square")
}

type Circle struct{
}

func (c Circle) Draw(){
	fmt.Println("heyyy am a circle")
}
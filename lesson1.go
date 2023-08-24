package main

import (
	"fmt"
	"time"
)

func runGoroutineLesson1() {
	fmt.Println("Go Routine Lesson 1: Goroutines")
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("terminated from main")
}

func sayHello() {
	fmt.Println("Hello world")
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

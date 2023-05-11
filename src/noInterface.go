package main

import "fmt"

type Dog struct{}

type Fish struct{}

type Bird struct{}

func (dog) walk() string {
	return "I am a dog and I am walking"
}

func (fish) swim() string {
	return "I am a fish and I am swimming"
}

func (bird) fly() string {
	return "I am a bird and I am flying"
}

func moveDog(d dog) {
	fmt.Println(d.walk())
}

func moveFish(f fish) {
	fmt.Println(f.swim())
}

func moveBird(b bird) {
	fmt.Println(b.fly())
}

func main() {
	d := dog{}
	moveDog(d)
	f := fish{}
	moveFish(f)
	b := bird{}
	moveBird(b)
}

package main

import "fmt"

// interfaceは型宣言のようなもので、Sayというメソッドを持ってないとダメですよという指定
type Human interface {
	Say() string
}

type Person struct {
	Name string
}

type Dog struct {
	Name string
}

func (p Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

// human.Say()を担保するために、Humanにinterfaceをつけている
func DriveCar(human Human) {
	if human.Say() == "MR.mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get out")
	}
}

func main() {
	var mike Human = &Person{"Mike"}
	var x Human = &Person{"x"}
	var dog Dog = Dog{"dog"}
	DriveCar(mike)
	DriveCar(x)
	// DogはSayというメソッドを持ってないのでエラーになる
	// cannot use dog (variable of type Dog) as Human value in argument to DriveCar: Dog does not implement Human (missing method Say)
	DriveCar(dog)
}

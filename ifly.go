package main

import(
	"fmt"
	"./bird"
	"./fly"
	"./animal"
)

func main() {

	afly := new(bird.Bird)
	var isName = "eagle"
	var isSex = "Male"
	afly.SetName(isName)
	afly.SetSex(isSex)
	var whoa string = afly.Fly()

	fmt.Printf("Struct from bird => %s\n",whoa)
        
	var bfly fly.IFly = new(animal.Animal)
        var isNameA = "bird"
        bfly.SetName(isNameA)
        var whob string = bfly.Fly()

	fmt.Printf("Struct from animal=> %s\n",whob)

	fmt.Printf("Nice to see you,hello world!\n")
}

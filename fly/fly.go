package fly  

//import "fmt"

type IFly interface {
	SetName(Name string)
	SetSex(Sex string)
	Fly() string
}

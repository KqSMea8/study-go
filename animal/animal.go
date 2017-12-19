package animal  

type Animal struct {
	Name string
	Sex string
}
func (a *Animal) Fly() string {
	return "The animal name " + a.Name + ",it can fly"
}
func (a *Animal) SetName(Name string) {
	a.Name = Name
}
func (a *Animal) SetSex(Sex string) {
	a.Sex = Sex
}

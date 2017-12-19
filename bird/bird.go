package bird 

type Bird struct{
        Name string
	Sex string
}
func (this *Bird) Fly() string {
	return "The bird name " + this.Name + ",it is " + this.Sex
}
func (b *Bird) SetName(Name string) {
	b.Name = Name
}
func (b *Bird) SetSex(Sex string) {
	b.Sex = Sex
}
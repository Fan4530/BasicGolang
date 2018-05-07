package _struct

import "fmt"

type Dog struct {
	name string
	weight int
	Sound string
}
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}
//func (d Dog) SpeakThreeTimes() {
//	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
//	fmt.Println(d.Sound)
//}

//问题： 指针和非指针是怎么操作的？？
//注意*Dog和Dog的定义的区别
func (d *Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}


func main() {
	poodle := Dog{"Poodle", 37, "Woof"}
	fmt.Println(poodle)
	poodle.Speak()

	poodle.Sound = "Arf"
	poodle.Speak()

	poodle.SpeakThreeTimes()
	poodle.SpeakThreeTimes()


}

//no overloading in golang
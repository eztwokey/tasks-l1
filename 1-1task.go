package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) Speak() {
	fmt.Printf("Привет, меня зовут %s, мне %d года.\n", h.Name, h.Age)
}

func (h *Human) Walk() {
	fmt.Printf("%s идет.\n", h.Name)
}

type Action struct {
	Human
	Hobby string
}

func (a *Action) DoAction() {
	fmt.Printf("%s занимается своим хобби: %s.\n", a.Name, a.Hobby)
}

func main() {

	action := &Action{
		Human: Human{
			Name: "Кирилл",
			Age:  23,
		},
		Hobby: "Программирование  для ВБ",
	}

	action.Speak()

	action.Walk()

	action.DoAction()

	fmt.Printf("Имя: %s\n", action.Name)
}

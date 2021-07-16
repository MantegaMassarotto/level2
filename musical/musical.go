package main

import "fmt"

type Trumpeter struct {
	Name string
}

type Violinist struct {
	Name string
}

type MusicalPlayer interface {
	Greetings() string
}

func (t *Trumpeter) Greetings() string {
	return fmt.Sprintf("Hello, my name is %v and I'm a trumpeter", t.Name)
}

func (t *Violinist) Greetings() string {
	return fmt.Sprintf("Hello, my name is %v and I'm a Violinist", t.Name)
}

func main() {
	v := &Violinist{"Murillo"}
	t1 := &Trumpeter{"Richard"}
	t2 := &Trumpeter{"Michael"}
	t3 := &Trumpeter{"Angel"}
	t4 := &Trumpeter{"Marco"}

	musicians := []MusicalPlayer{v, t1, t2, t3, t4}

	for _, musician := range musicians {
		fmt.Println(musician.Greetings())
	}
}

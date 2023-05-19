package main

import "fmt"

type F struct {
	Name string
	Age  int
	Gs   []*G
}

func (f *F) String() string {
	return fmt.Sprintf("NAME=%q, AGE=%d, G=%v", f.Name, f.Age, f.Gs)
}

type G struct {
	Type string
}

func (g *G) String() string {
	return fmt.Sprintf("Type=%q", g.Type)
}

func main() {
	f := &F{
		Name: "John",
		Age:  20,
		Gs: []*G{
			&G{
				Type: "A-type",
			},
			&G{
				Type: "B-type",
			},
		},
	}

	fmt.Printf("%v\n", f)
	fmt.Printf("%#v\n", f)
	fmt.Printf("%T\n", f)
}

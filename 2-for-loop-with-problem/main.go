package main

type Gopher struct {
	Name string
}

func main() {
	oldGophers := []Gopher{
		{"Phil"},
		{"Nir"},
		{"Josh"},
	}
	newGophers := make([]*Gopher, 0)
	for _, gopher := range oldGophers {
		newGophers = append(newGophers, &gopher)
	}

	for _, gopher := range newGophers {
		println(gopher.Name)
	}
}

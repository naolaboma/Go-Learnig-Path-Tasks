package main

import (
	"fmt"
	"os"
)

type Grade struct {
	Name string
	Subj map[string]float64
}

func newStGrade(name string) Grade {
	g := Grade{
		Name: name,
		Subj: map[string]float64{},
	}
	return g
}

// add a name
func (g *Grade) addName(subject string, grade float64) {
	g.Subj[subject] = grade
}

func (g Grade) format() string {
	fs := fmt.Sprintf("Grade report:\n%-25v %v\n", "Subject", "Grade(numeric)")
	var total float64 = 0
	length := 0
	// list the grades
	for k, v := range g.Subj {
		fs += fmt.Sprintf("%-25v     %v \n", k+":", v)
		total += v
		length += 1
	}
	var avg float64 = total / float64(length)
	fs += fmt.Sprintf("\n\n%-25v     %0.2f\n", "total:", total)
	fs += fmt.Sprintf("%-25v     %0.2f\n", "avg:", avg)
	return fs
}
func (g *Grade) save() {
	data := []byte(g.format())
	err := os.WriteFile("grades/"+g.Name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Grade was saved to file")
}

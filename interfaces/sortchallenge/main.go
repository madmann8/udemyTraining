package main

import (
	"sort"
	"fmt"
)

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {
studyGroup := people{"John", "Zeno", "Al", "Jenny"}
sort.Sort(people(studyGroup))
fmt.Print(studyGroup)
sort.Sort(sort.Reverse(studyGroup))
fmt.Print(studyGroup)
}
package main

import "sort"

type personSort struct {
	persons []Person
	by      By
}

func (by By) Sort(persons []Person) {
	ps := &personSort{}
	sort.Sort(ps) // need 3 function， Len, Swap, Less implement sort;
}

func (s *personSort) Len() int {
	return len(s.persons)
}

func (s *personSort) Swap(i, j int) {
	s.persons[i], s.persons[j] = s.persons[j], s.persons[i]
}

func (s *personSort) Less(i, j int) bool {
	return s.by(&s.persons[i], &s.persons[j])
}

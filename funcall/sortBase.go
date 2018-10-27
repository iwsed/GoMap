package main

type ByAge []Person

func (b ByAge) Len() int { return len(b) }

func (b ByAge) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

func (b ByAge) Less(i, j int) bool { return b[i].age < b[j].age }

package main

import "fmt"

type deck [] int

func (d deck) print(){
	for i,card := range d{
		j:=i*2
		fmt.Println(j,card)
	}
}

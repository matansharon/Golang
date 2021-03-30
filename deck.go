package main

import "fmt"

type deck [] string

func (d deck) print(){
	for i,card := range d{
		
		fmt.Println(i,card)
	}
}

func new_deck() deck{
	cards:= deck{}
	suiets := []string{"Spades","Diamonds","Hearts","Clubs"}
	values:= []string{"Ace","Two","Three","Four","Five","Six","Seven","Eight","Nine","Ten","Jack","Queen","King"}
	for _,suit := range suiets{
		for _,value := range values{
			cards=append(cards,suit+" of "+value)
		}
	}
	return cards
}

func deal (d deck, hand int)(deck,deck){
	return d[:hand],d[hand:]
}

func (d deck)to_byte_slice([]byte){
	st:=""
	for _,card := range d{
		st=append(st,card+" ")
	}
	return []byte(st)
}
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	suit  string
	value int
}

type Deck struct {
	cards [52]Card
}

var Suits []string = []string{"♤", "♥", "♢", "♧"}
var Values []string = []string{
	"Ace", "Two", "Three", "Four", "Five",
	"Six", "Seven", "Eight", "Nine", "Ten",
	"Jack", "Queen", "King"}

func New() Deck {
	d := Deck{}
	var i int = 0
	for k, _ := range d.cards {
		var c = (i * 13) - k + 1
		if k%13 == 0 {
			i++
		}
		d.cards[k] = Card{Suits[i-1], c}
	}
	return d
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := range d.cards {
		j := rand.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

func (d *Deck) Print() {
	for _, c := range d.cards {
		fmt.Printf("%s\n", c.Name())
	}
}

func (c *Card) Name() string {
	return fmt.Sprintf("%s (%d) %s", Values[c.value-1], c.value, c.suit)
}

func main() {
	d := New()
	d.Shuffle()
	d.Print()
}

package Deck

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck []Card

type Card struct {
	Number string
	Suit   string
	Rank   float64
}

func NewDeck(deckNumber int) Deck {
	cards := Deck{}
	for v := 0; v < deckNumber; v++ {
		numbers := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
		suits := []string{"♣", "♦", "♥", "♠"}
		ranks := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 0.5, 0.5, 0.5}

		for i, number := range numbers {
			for index := range suits {
				cards = append(cards, Card{Number: number, Suit: suits[index], Rank: ranks[i]})
			}
		}
	}
	return cards
}

func (d Deck) Print() {
	for i := range d {
		fmt.Println(d[i])
	}
}

func (d Deck) Shuffle() {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

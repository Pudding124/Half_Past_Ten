package main

import (
	"Poker/Deck"
	"Poker/Rule"
	"fmt"
)

func main() {
	deck := Deck.NewDeck(2)
	deck.Shuffle()
	peopleList := []*Rule.People{}
	banker := Rule.People{Owner: "D", Money: 100, Banker: true}
	people_1 := Rule.People{Owner: "A", Money: 100}
	people_2 := Rule.People{Owner: "B", Money: 100}
	people_3 := Rule.People{Owner: "C", Money: 100}
	peopleList = append(peopleList, &people_1)
	peopleList = append(peopleList, &people_2)
	peopleList = append(peopleList, &people_3)

	Rule.StartGame(peopleList, &banker, deck)
	Rule.JudgeGame(peopleList, &banker)

	fmt.Println()
	fmt.Printf("最終結果: 莊家 %v: ", banker)
	fmt.Println()
	for _, p := range peopleList {
		fmt.Printf("最終結果: 玩家 %v: ", p)
		fmt.Println()
	}
}

package Rule

import (
	"Poker/Deck"
	"fmt"
)

type People struct {
	Banker    bool
	Owner     string
	Money     int
	Rank      float64
	TotalCard int
}

func isRankExceeded(rank float64) bool {
	if rank > 10.5 {
		return true
	}
	return false
}

func StartGame(gamers []People, banker *People, decks Deck.Deck) {
	// 先均發一張牌
	deckIndex := 0
	(*banker).Rank = decks[deckIndex].Rank
	(*banker).TotalCard += 1
	deckIndex += 1
	for index := range gamers {
		(&gamers[index]).Rank = decks[deckIndex].Rank
		(&gamers[index]).TotalCard += 1
		deckIndex += 1
	}

	choose := ""
	for index := range gamers {
		user := &gamers[index]
		fmt.Println("Gamer: ", user.Owner)

		for true {
			fmt.Println("分數: ", user.Rank)
			fmt.Print("還要牌嗎? ")
			fmt.Scanln(&choose)

			if choose == "Y" {
				user.Rank += decks[deckIndex].Rank
				(&gamers[index]).TotalCard += 1
				deckIndex += 1

				if isRankExceeded(user.Rank) {
					fmt.Println("分數: ", user.Rank)
					fmt.Println("分數超過了!!")
					break
				}
			} else if choose == "N" {
				break
			}
		}
	}

	for true {
		fmt.Println("莊家分數: ", (*banker).Rank)
		fmt.Print("還要牌嗎? ")
		fmt.Scanln(&choose)

		if choose == "Y" {
			(*banker).Rank += decks[deckIndex].Rank
			(*banker).TotalCard += 1
			deckIndex += 1

			if isRankExceeded((*banker).Rank) {
				fmt.Println("莊家分數超過了!!")
				break
			}
		} else if (*banker).Rank < 6 {
			fmt.Println("莊家分數小於 6，一定需要牌")
			(*banker).Rank += decks[deckIndex].Rank
			(*banker).TotalCard += 1
			deckIndex += 1

			if isRankExceeded((*banker).Rank) {
				fmt.Println("莊家分數超過了!!")
				break
			}
		} else {
			break
		}
	}
}

func JudgeGame(gamers []People, banker *People) {
	// Rule 1 -> Banker Fail
	if isRankExceeded((*banker).Rank) {
		for index := range gamers {
			if isRankExceeded((&gamers[index]).Rank) == false {
				(*banker).Money -= 10 * handType(gamers[index])
				(&gamers[index]).Money += 10 * handType(gamers[index])
			}
		}
	} else { // Rule 2 -> Banker Success
		for index := range gamers {
			if isRankExceeded((&gamers[index]).Rank) { // gamer fail
				(*banker).Money += 10 * handType(*banker)
				(&gamers[index]).Money -= 10 * handType(*banker)
			} else if (*banker).Rank > (&gamers[index]).Rank { // 莊家大
				if handType(*banker) == 1 {
					if handType(gamers[index]) == 1 {
						(*banker).Money += 10 * handType(*banker)
						(&gamers[index]).Money -= 10 * handType(*banker)
					} else {
						(*banker).Money -= 10 * handType(gamers[index])
						(&gamers[index]).Money += 10 * handType(gamers[index])
					}
				} else if handType(*banker) == 2 {
					if handType(gamers[index]) == 1 {
						(*banker).Money += 10 * handType(*banker)
						(&gamers[index]).Money -= 10 * handType(*banker)
					} else if handType(gamers[index]) == 2 {
						continue
					} else if handType(gamers[index]) == 5 {
						(*banker).Money -= 10 * (handType(gamers[index]) - handType(*banker))
						(&gamers[index]).Money += 10 * (handType(gamers[index]) - handType(*banker))
					}
				} else if handType(*banker) == 5 {
					if handType(gamers[index]) == 1 {
						(*banker).Money += 10 * handType(*banker)
						(&gamers[index]).Money -= 10 * handType(*banker)
					} else if handType(gamers[index]) == 2 {
						(*banker).Money += 10 * (handType(*banker) - handType(gamers[index]))
						(&gamers[index]).Money -= 10 * (handType(*banker) - handType(gamers[index]))
					} else if handType(gamers[index]) == 5 {
						continue
					}
				}
			} else if (*banker).Rank < (&gamers[index]).Rank { // 莊家小
				if handType(*banker) == 1 {
					if handType(gamers[index]) == 1 {
						(*banker).Money -= 10 * handType(*banker)
						(&gamers[index]).Money += 10 * handType(*banker)
					} else {
						(*banker).Money -= 10 * handType(gamers[index])
						(&gamers[index]).Money += 10 * handType(gamers[index])
					}
				} else if handType(*banker) == 2 {
					if handType(gamers[index]) == 1 {
						(*banker).Money += 10 * handType(*banker)
						(&gamers[index]).Money -= 10 * handType(*banker)
					} else if handType(gamers[index]) == 2 {
						continue
					} else if handType(gamers[index]) == 5 {
						(*banker).Money -= 10 * (handType(gamers[index]) - handType(*banker))
						(&gamers[index]).Money += 10 * (handType(gamers[index]) - handType(*banker))
					}
				} else if handType(*banker) == 5 {
					if handType(gamers[index]) == 1 {
						(*banker).Money += 10 * handType(*banker)
						(&gamers[index]).Money -= 10 * handType(*banker)
					} else if handType(gamers[index]) == 2 {
						(*banker).Money += 10 * (handType(*banker) - handType(gamers[index]))
						(&gamers[index]).Money -= 10 * (handType(*banker) - handType(gamers[index]))
					} else if handType(gamers[index]) == 5 {
						continue
					}
				}
			} else if (*banker).Rank == (&gamers[index]).Rank { // 莊家小
				if handType(*banker) == 1 {
					if handType(gamers[index]) == 1 {
						continue
					} else {
						(*banker).Money -= 10 * handType(gamers[index])
						(&gamers[index]).Money += 10 * handType(gamers[index])
					}
				} else if handType(*banker) == 2 {
					if handType(gamers[index]) == 1 {
						(*banker).Money += 10 * handType(*banker)
						(&gamers[index]).Money -= 10 * handType(*banker)
					} else if handType(gamers[index]) == 2 {
						continue
					} else if handType(gamers[index]) == 5 {
						(*banker).Money -= 10 * (handType(gamers[index]) - handType(*banker))
						(&gamers[index]).Money += 10 * (handType(gamers[index]) - handType(*banker))
					}
				} else if handType(*banker) == 5 {
					if handType(gamers[index]) == 1 {
						(*banker).Money += 10 * handType(*banker)
						(&gamers[index]).Money -= 10 * handType(*banker)
					} else if handType(gamers[index]) == 2 {
						(*banker).Money += 10 * (handType(*banker) - handType(gamers[index]))
						(&gamers[index]).Money -= 10 * (handType(*banker) - handType(gamers[index]))
					} else if handType(gamers[index]) == 5 {
						continue
					}
				}
			}
		}
	}
}

func handType(people People) int {
	multiplier := 1
	if people.Rank == 10.5 && people.TotalCard == 2 {
		return 2
	} else if people.TotalCard == 5 {
		return 3
	} else if people.Rank == 10.5 && people.TotalCard == 5 {
		return 5
	}
	return multiplier
}

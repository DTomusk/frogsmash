package game

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

const kFactor = 32.0

type Item struct {
	Id    int
	Score float64
}

func Run(items []Item) {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	for {
		item1, item2 := pickTwoItems(items)
		fmt.Printf("Comparing Item %d (Score: %.2f) with Item %d (Score: %.2f)\n", item1.Id, item1.Score, item2.Id, item2.Score)
		fmt.Print("Which item is better? (0 or 1): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		var choice int
		_, err = fmt.Sscanf(input, "%d", &choice)
		if err != nil || (choice != 0 && choice != 1) {
			fmt.Println("Invalid input. Please enter 0 or 1.")
			continue
		}
		compare(item1, item2, choice)
	}
}

func GenerateItems() []Item {
	items := make([]Item, 0)
	for i := 1; i <= 10; i++ {
		item := Item{
			Id:    i,
			Score: 1000,
		}
		items = append(items, item)
	}
	return items
}

func pickTwoItems(items []Item) (*Item, *Item) {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if len(items) < 2 {
		return nil, nil // not enough items
	}

	i := r.Intn(len(items))
	j := r.Intn(len(items) - 1)
	if j >= i {
		j++
	}

	return &items[i], &items[j]
}

func compare(item1, item2 *Item, choice int) {
	if choice == 0 {
		updateEloScores(item1, item2)
	} else {
		updateEloScores(item2, item1)
	}
}

func updateEloScores(winner, loser *Item) {
	expectedWinner := 1 / (1 + math.Pow(10, (loser.Score-winner.Score)/400))
	expectedLoser := 1 / (1 + math.Pow(10, (winner.Score-loser.Score)/400))
	winner.Score += kFactor * (1 - expectedWinner)
	loser.Score += kFactor * (0 - expectedLoser)
}

package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l *League) MatchResult(first, second string, firstScore, secondScore int) {
	if firstScore > secondScore {
		l.Wins[first]++
	} else if secondScore > firstScore {
		l.Wins[second]++
	}
}

func (l League) Ranking() []string {
	sort.Slice(l.Teams, func(i, j int) bool {
		return l.Wins[l.Teams[i].Name] > l.Wins[l.Teams[j].Name]
	})

	var orderedNames []string
	for _, team := range l.Teams {
		orderedNames = append(orderedNames, team.Name)
	}

	return orderedNames
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	results := r.Ranking()
	for _, v := range results {
		io.WriteString(w, v)
		io.WriteString(w, "\n")
	}
}

func main() {
	team1 := Team{"FC Original", []string{"Player1", "Player2", "Player3"}}
	team2 := Team{"FC Oriji", []string{"Player4", "Player5", "Player6"}}
	team3 := Team{"FC O", []string{"Player7", "Player8", "Player9"}}

	league1 := League{[]Team{team1, team2, team3}, map[string]int{}}
	league1.MatchResult("FC Original", "FC Oriji", 3, 0)
	league1.MatchResult("FC Oriji", "FC O", 2, 1)
	league1.MatchResult("FC Original", "FC O", 1, 0)
	league1.MatchResult("FC Oriji", "FC O", 1, 3)
	league1.MatchResult("FC Original", "FC O", 1, 4)
	league1.MatchResult("FC Original", "FC O", 1, 5)
	league1.MatchResult("FC Oriji", "FC O", 4, 3)
	league1.MatchResult("FC Oriji", "FC O", 4, 3)
	league1.MatchResult("FC Oriji", "FC O", 4, 3)

	fmt.Println(league1.Wins)

	fmt.Println(league1.Ranking())

	RankPrinter(league1, os.Stdout)
}

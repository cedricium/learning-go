package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name   string
	Teams  map[string]Team
	Wins   map[string]int
	Points map[string]int
}

func (l *League) MatchResult(a string, as int, h string, hs int) {
	if _, ok := l.Teams[a]; !ok {
		return
	}
	if _, ok := l.Teams[h]; !ok {
		return
	}
	if as < 0 || hs < 0 {
		return
	}
	if as > hs {
		l.Wins[a]++
	} else if hs > as {
		l.Wins[h]++
	}
	l.Points[a] += as
	l.Points[h] += hs
}

func (l *League) Ranking() []string {
	// initialize slice for team names (empty for now)
	teams := make([]string, 0, len(l.Teams))

	// populate teams using key of l.Teams
	for k := range l.Teams {
		teams = append(teams, k)
	}

	// sort teams using same team keys, comparing against l.Wins
	// sort.Slice(teams, func(i, j int) bool {
	// 	return l.Wins[teams[i]] > l.Wins[teams[j]]
	// })

	// improved sorted ranking: wins first, team points second, and last by name
	slices.SortFunc(teams, func(i, j string) int {
		return cmp.Or(
			cmp.Compare(l.Wins[j], l.Wins[i]),
			cmp.Compare(l.Points[j], l.Points[i]),
			cmp.Compare(l.Teams[j].Name, l.Teams[i].Name),
		)
	})
	return teams
}

func main() {
	nfl := League{
		Name: "National Football League",
		Teams: map[string]Team{
			"49ers":     {Name: "SF 49ers", Players: []string{"Deebo Samuel"}},
			"Cardinals": {Name: "AZ Cardinals", Players: []string{"Marvin Harrison Jr."}},
			"Jets":      {Name: "NY Jets", Players: []string{"Mike Williams"}},
			"Rams":      {Name: "LA Rams", Players: []string{"Puka Nacua"}},
			"Seahawks":  {Name: "SEA Seahawks", Players: []string{"DK Metcalf"}},
		},
		Wins:   map[string]int{},
		Points: map[string]int{},
	}

	nfl.MatchResult("49ers", 7, "Seahawks", 21)
	nfl.MatchResult("Rams", 28, "Cardinals", 17)
	nfl.MatchResult("Seahawks", 10, "Cardinals", 7)
	nfl.MatchResult("Rams", 10, "49ers", 14)
	nfl.MatchResult("Cardinals", 10, "Rams", 7)
	nfl.MatchResult("Seahawks", 35, "49ers", 21)
	nfl.MatchResult("Cardinals", 10, "Jets", 24)
	nfl.MatchResult("49ers", 14, "Rams", 3)
	nfl.MatchResult("Rams", 14, "Jets", 3)

	fmt.Println(nfl.Points)    // map[49ers:56 Cardinals:44 Jets:27 Rams:62 Seahawks:66]
	fmt.Println(nfl.Wins)      // map[49ers:2 Cardinals:1 Jets:1 Rams:2 Seahawks:3]
	fmt.Println(nfl.Ranking()) // [Seahawks Rams 49ers Cardinals Jets] - Wins > Points > Name
	// fmt.Println(nfl.Ranking()) // [Seahawks 49ers Rams Cardinals Jets] - ties (SF, LA) are random
}

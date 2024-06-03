/*
CLI version of the popular puzzle game of the same name, Battleship.

Grids are arranged as follows: columns are denoted using numbers, 1-10.
Rows are denoted using letters, A-J.

Terminology:
  - fleet: ships to be positioned and attacked. for ship types, see below
  - ocean grid: player's own grid where player's fleet is positioned
  - target grid: opponent's ocean grid, where player's shots are targeted
  - shot: letter-number (e.g. D-4) coordinate combo denoting location on target grid

Fleet (lengths denoted in parenthesis):
  - Carrier 		(5)
  - Battleship 	(4)
  - Destroyer 	(3)
  - Submarine 	(3)
  - Patrol Boat (2)

Roughly following the Hasbro "Battleship" (2002) rules: https://www.hasbro.com/common/instruct/BattleShip_(2002).PDF
*/
package main

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	shotHit  string = "x"
	shotMiss string = "."
)

var grid [][]string = [][]string{
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
}

type Coordinate struct {
	Row int
	Col int
}

// Coordinates decodes a shot string in the expected alphanumeric form
// `<char><num>` (<char> A-J row, <num> 0-9 column) and returns the
// corresponding row and column indices for the given shot.
func coordinates(shot string) (*Coordinate, error) {
	if len(shot) != 2 {
		return nil, errors.New("invalid shot format")
	}

	runes := bytes.Runes([]byte(strings.ToUpper(shot)))
	if !unicode.IsLetter(runes[0]) {
		return nil, errors.New("invalid shot row, must be letter (a-z, A-Z)")
	}
	row := runes[0] - 65

	if !unicode.IsNumber(runes[1]) {
		return nil, errors.New("invalid shot col, must be number (0-9)")
	}
	col, _ := strconv.Atoi(string(shot[1]))

	return &Coordinate{int(row), col}, nil
}

func main() {
	// === simulating placing fleet ===
	grid[9][3], grid[9][4] = "▪️", "▪️"                                                       // patrol boat
	grid[1][1], grid[2][1], grid[3][1] = "▪️", "▪️", "▪️"                                     // submarine
	grid[4][7], grid[4][8], grid[4][9] = "▪️", "▪️", "▪️"                                     // destroyer
	grid[5][2], grid[6][2], grid[7][2], grid[8][2] = "▪️", "▪️", "▪️", "▪️"                   // battleship
	grid[0][5], grid[1][5], grid[2][5], grid[3][5], grid[4][5] = "▪️", "▪️", "▪️", "▪️", "▪️" // carrier

	// === simulating attacks ===
	coord, _ := coordinates("H8")
	grid[coord.Row][coord.Col] = shotMiss // shot (miss): H:8 —> grid[7][8]
	coord, _ = coordinates("J4")
	grid[coord.Row][coord.Col] = shotHit // shot (hit): J:5 —> grid[9][4]

	// === displaying grid with fleet and shots ===
	fmt.Println("   0 1 2 3 4 5 6 7 8 9")
	// fmt.Println("   0 1 2 3 4 5 6 7 8 9      0 1 2 3 4 5 6 7 8 9")
	for row := range grid {
		// fmt.Println(row, grid[row])                  // DEBUG - uses index for row
		fmt.Println(string(rune(row)+65), grid[row]) // uses characters for row
		// fmt.Printf("%c %v   %c %v\n", row+65, grid[row], row+65, grid[row]) // uses characters for row
	}
}

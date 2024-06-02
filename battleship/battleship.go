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
	"errors"
	"fmt"
	"strconv"
	"strings"
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

// Coordinates decodes a shot string in the expected alphanumeric form
// `<char>:<num>` (<char> A-J row, <num> 0-9 column), and returns the
// corresponding row and column indices for the given shot. Errors are
// returned if the shot format is not the correct length, is an invalid
// format, or the shot falls outside of the specified row/col ranges.
func coordinates(shot string) (int8, int8, error) {
	if len(shot) != 3 {
		return -1, -1, errors.New("invalid shot length")
	}

	if !strings.Contains(shot, ":") {
		return -1, -1, errors.New("invalid shot format")
	}
	coords := strings.Split(shot, ":")
	if len(coords) != 2 || len(coords[0]) == 0 || len(coords[1]) == 0 {
		return -1, -1, errors.New("invalid shot format")
	}

	char := strings.ToUpper(coords[0])
	if char[0] < 65 || char[0] > 74 {
		return -1, -1, errors.New("shot row out of range")
	}
	row := rune(char[0]) - 65

	col, err := strconv.Atoi(coords[1])
	if err != nil {
		return -1, -1, err
	}

	return int8(row), int8(col), nil
}

func main() {
	// === simulating placing fleet ===
	grid[9][3], grid[9][4] = "▪️", "▪️"                                                       // patrol boat
	grid[1][1], grid[2][1], grid[3][1] = "▪️", "▪️", "▪️"                                     // submarine
	grid[4][7], grid[4][8], grid[4][9] = "▪️", "▪️", "▪️"                                     // destroyer
	grid[5][2], grid[6][2], grid[7][2], grid[8][2] = "▪️", "▪️", "▪️", "▪️"                   // battleship
	grid[0][5], grid[1][5], grid[2][5], grid[3][5], grid[4][5] = "▪️", "▪️", "▪️", "▪️", "▪️" // carrier

	// === simulating attacks ===
	r, c, _ := coordinates("H:8")
	grid[r][c] = shotMiss // shot (miss): H:8 —> grid[7][8]
	r, c, _ = coordinates("J:4")
	grid[r][c] = shotHit // shot (hit): J:5 —> grid[9][4]

	// === displaying grid with fleet and shots ===
	fmt.Println("   0 1 2 3 4 5 6 7 8 9")
	// fmt.Println("   0 1 2 3 4 5 6 7 8 9      0 1 2 3 4 5 6 7 8 9")
	for row := range grid {
		// fmt.Println(row, grid[row])                  // DEBUG - uses index for row
		fmt.Println(string(rune(row)+65), grid[row]) // uses characters for row
		// fmt.Printf("%c %v   %c %v\n", row+65, grid[row], row+65, grid[row]) // uses characters for row
	}

	// fmt.Println(coordinates("A:0"))  // 0, 0, nil
	// fmt.Println(coordinates("J:9"))  // 9, 9, nil
	// fmt.Println(coordinates("C:40")) // -1, -1, invalid shot length
	// fmt.Println(coordinates("C-4"))  // -1, -1, invalid shot format
	// fmt.Println(coordinates(":12"))  // -1, -1, invalid shot format
	// fmt.Println(coordinates("AA:"))  // -1, -1, invalid shot format
	// fmt.Println(coordinates("Z:4"))  // -1, -1, show row out of range
	// fmt.Println(coordinates("C:Z"))  // -1, -1, strconv.Atoi: parsing "Z": invalid syntax
}

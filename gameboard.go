package main

import (
	"bytes"
	"errors"
)


type SnakeSegment struct {
	Index int
	Parent *Battlesnake
	ActiveBattlesnake bool
}

func (s SnakeSegment) IsHead() bool {
	return s.Parent != nil && s.Parent.Length > 0 && s.Index == 0
}

func (s SnakeSegment) IsTail() bool {
	return s.Parent != nil && s.Parent.Length > 0 && s.Index == s.Parent.Length - 1
}

func (s SnakeSegment) IsBody() bool {
	return s.Parent != nil && s.Parent.Length > 0 && !(s.IsHead() || s.IsTail())
}

// TODO - edge case, doesn't handle when stacked at start of game
type BoardSpace struct {
	Snake *SnakeSegment
	Hazard bool
	Food bool
}
func (b BoardSpace) IsEmpty() bool {
	return !b.Hazard && !b.Food && !b.ContainsSnake()
}

func (b BoardSpace) ContainsSnakeHead() bool {
	return b.Snake != nil && b.Snake.IsHead()
}

func (b BoardSpace) ContainsSnakeTail() bool {
	return b.Snake != nil && b.Snake.IsTail()
}

func (b BoardSpace) ContainsSnake() bool {
	return b.Snake != nil
}


type GameBoard struct {
	width int
	height int
	cells [][]BoardSpace
}

func NewGameBoard(board *Board, activeBattlesnake *Battlesnake) (*GameBoard, error) {
	contents := [][]BoardSpace{}

	for i := 0; i < board.Height; i++ {
		row := make([]BoardSpace, board.Width)
		contents = append(contents, row)
	}

	for i := 0; i < len(board.Food); i++ {
		pos := board.Food[i]
		contents[pos.Y][pos.X].Food = true
	}

	for i := 0; i < len(board.Hazards); i++ {
		pos := board.Hazards[i]
		contents[pos.Y][pos.X].Hazard = true
	}

	for i := 0; i < len(board.Snakes); i++ {
		snake := board.Snakes[i]

		for x := 0; x < len(snake.Body); x++ {
			segment := SnakeSegment{Index: x, Parent: &snake, ActiveBattlesnake: activeBattlesnake.ID == snake.ID}
			pos := snake.Body[x]
			contents[pos.Y][pos.X].Snake = &segment
		}
	}

	gb := GameBoard{ width: board.Width, height: board.Height, cells: contents}

	return &gb, nil
}

func (b *GameBoard) GetCell(pos Coord)(*BoardSpace, error) {
	if b.IsOutOfBounds(pos) {
		return nil, errors.New("Coordinate outside of the board space")
	}
	return &b.cells[pos.Y][pos.X], nil
}

func (b *GameBoard) Print() {

	var buf bytes.Buffer

	for y := 0; y < len(b.cells); y++ {
		for x:=0; x < len(b.cells[y]); x++ {
			cell := b.cells[y][x] 
			if cell.IsEmpty() {
				buf.WriteString("  ")

			} else if cell.Food {
				buf.WriteString("F ")
				
			} else if cell.Hazard {
				buf.WriteString("H ")

			} else if cell.ContainsSnake() {
				buf.WriteString("S ")
			}
			buf.WriteString("\n")
		}
	}
}

func (b *GameBoard) IsOutOfBounds(pos Coord) bool {
	if pos.X < 0 || pos.X >= b.width {
		return true
	}
	
	if pos.Y <0 || pos.Y >= b.height {
		return true
	}

	return false
}

func (b *GameBoard) ContainsHazard(pos Coord) (bool) {
	cell, err := b.GetCell(pos)

	if err != nil {
		return false
	}

	return cell.Hazard
}

func (b *GameBoard) IsASnake(pos Coord) (bool) {
	cell, err := b.GetCell(pos)

	if err != nil {
		return false
	}

	return cell.ContainsSnake()
}


func (b *GameBoard) ContainsFood(pos Coord) (bool) {
	cell, err := b.GetCell(pos)

	if err != nil {
		return false
	}

	return cell.Food
}
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var EmptyTestSnake = Battlesnake{
	Body: []Coord{
	},
	Length: 0,
}

var OneSegmentTestSnake = Battlesnake{
	Body: []Coord{
		{X: 0, Y: 0},
	},
	Length: 1,
}

var TestSnake = Battlesnake{
	Body: []Coord{
		{X: 0, Y: 0},
		{X: 1, Y: 0},
		{X: 2, Y: 0},
		{X: 2, Y: 1},
		{X: 2, Y: 2},
	},
	Length: 5,
}

func TestSnakeSegment_IsHead(t *testing.T) {
	segment := SnakeSegment{Index: 0, Parent: nil, ActiveBattlesnake: false}
	assert.False(t, segment.IsHead())

	segment = SnakeSegment{Index: 0, Parent: &EmptyTestSnake, ActiveBattlesnake: false}
	assert.False(t, segment.IsHead())

	segment = SnakeSegment{Index: 0, Parent: &OneSegmentTestSnake, ActiveBattlesnake: false}
	assert.True(t, segment.IsHead())

	segment = SnakeSegment{Index: 0, Parent: &TestSnake, ActiveBattlesnake: false}
	assert.True(t, segment.IsHead())

	segment = SnakeSegment{Index: 2, Parent: &TestSnake, ActiveBattlesnake: false}
	assert.False(t, segment.IsHead())
}

func TestSnakeSegment_IsTail(t *testing.T) {
	segment := SnakeSegment{Index: 0, Parent: nil, ActiveBattlesnake: false}
	assert.False(t, segment.IsTail())

	segment = SnakeSegment{Index: 0, Parent: &EmptyTestSnake, ActiveBattlesnake: false}
	assert.False(t, segment.IsTail())

	segment = SnakeSegment{Index: 0, Parent: &OneSegmentTestSnake, ActiveBattlesnake: false}
	assert.True(t, segment.IsTail())

	segment = SnakeSegment{Index: 4, Parent: &TestSnake, ActiveBattlesnake: false}
	assert.True(t, segment.IsTail())

	segment = SnakeSegment{Index: 3, Parent: &TestSnake, ActiveBattlesnake: false}
	assert.False(t, segment.IsTail())
}

func TestSnakeSegment_IsBody(t *testing.T) {
	segment := SnakeSegment{Index: 0, Parent: nil, ActiveBattlesnake: false}
	assert.False(t, segment.IsBody())

	segment = SnakeSegment{Index: 0, Parent: &EmptyTestSnake, ActiveBattlesnake: false}
	assert.False(t, segment.IsBody())

	segment = SnakeSegment{Index: 0, Parent: &OneSegmentTestSnake, ActiveBattlesnake: false}
	assert.False(t, segment.IsBody())

	segment = SnakeSegment{Index: 2, Parent: &TestSnake, ActiveBattlesnake: false}
	assert.True(t, segment.IsBody())

	segment = SnakeSegment{Index: 0, Parent: &TestSnake, ActiveBattlesnake: false}
	assert.False(t, segment.IsBody())

	segment = SnakeSegment{Index: 4, Parent: &TestSnake, ActiveBattlesnake: false}
	assert.False(t, segment.IsBody())

}
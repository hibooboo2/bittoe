package main

import (
	"fmt"
	"log"
)

type Board int32

func (b Board) String() string {
	s := make([]player, 9)
	for i := range s {
		s[i] = b.Space((XTopLeft << uint(i)))
	}
	return fmt.Sprintf(`
		 %s| %s| %s
		________
		 %s| %s| %s
		________
		 %s| %s| %s
		`,
		s[0], s[1], s[2],
		s[3], s[4], s[5],
		s[6], s[7], s[8])

}

type player rune

func (p player) String() string {
	return string(p)
}

func (b Board) Space(pos Board) player {
	pos &= ClearBoard
	switch b & (pos | pos<<9) {
	case pos:
		return 'X'
	case pos << 9:
		return 'O'
	case pos | pos<<9:
		panic("How did you play 2 players on tile!")
	default:
		return ' '
	}
}

func (b *Board) Move(pos Board) bool {
	pos = pos & ClearBoard
	if pos >= OBottomRight+1 || pos < XTopLeft {
		log.Printf("Not valid range: [%b]\n", pos)
		return false
	}

	cmp := pos
	if cmp > XBottomRight {
		cmp = cmp << 9
	}

	if (*b&cmp) > 0 || (*b&(cmp<<9)) > 0 {
		log.Println("Invalid move move taken [%b] [%b]\n", *b, cmp)
		return false
	}

	*b = *b | pos
	return true
}

const (
	BoardStart Board = 1<<31 - 1
	None       Board = 1 << iota
	XTopLeft
	XMiddleLeft
	XBottomLeft
	XTopCenter
	XMiddleCenter
	XBottomCenter
	XTopRight
	XMiddleRight
	XBottomRight
	OTopLeft
	OMiddleLeft
	OBottomLeft
	OTopCenter
	OMiddleCenter
	OBottomCenter
	OTopRight
	OMiddleRight
	OBottomRight
	IsO        Board = OTopLeft | OMiddleLeft | OBottomLeft | OTopCenter | OMiddleCenter | OBottomCenter | OTopRight | OMiddleRight | OBottomRight
	IsX        Board = XTopLeft | XMiddleLeft | XBottomLeft | XTopCenter | XMiddleCenter | XBottomCenter | XTopRight | XMiddleRight | XBottomRight
	ClearBoard Board = XTopLeft | XMiddleLeft | XBottomLeft | XTopCenter | XMiddleCenter | XBottomCenter | XTopRight | XMiddleRight | XBottomRight | OTopLeft | OMiddleLeft | OBottomLeft | OTopCenter | OMiddleCenter | OBottomCenter | OTopRight | OMiddleRight | OBottomRight
	BoardMask  Board = 0xFFFFFFF ^ ClearBoard
)

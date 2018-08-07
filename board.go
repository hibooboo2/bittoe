package main

import (
	"fmt"
	"log"
)

type Board uint32

const PlayerShift = 9

func (b *Board) Play(moves chan uint) {
	for m := range moves {
		move := b.ParseMove(m)
		if !b.Move(move) {
			log.Println("Invalid move please try again.")
			continue
		}
		switch b.GameOver() {
		case 'X':
			fmt.Println("Xes win!", b)
			return
		case 'O':
			fmt.Println("Oes win!", b)
			return
		case ' ':
			continue
		case 'C':
			fmt.Println("Cats Game(tie)", b)
			return
		}
		fmt.Println(b.String())
	}
}

func (b Board) ParseMove(m uint) Board {
	if m > 8 {
		panic("Invalid move")
	}
	move := XTopLeft << m
	if b&Player == 0 {
		move = move << PlayerShift
	}
	return move
}

func (b Board) GameOver() player {
	var alreadywon bool
	var won player = ' '
	for win, player := range winner {
		if win&b == win {
			if !alreadywon {
				won = player
				alreadywon = true
			} else {
				if player != won {
					panic("2 People cannot win!")
				}
			}
		}
	}
	return won
}

func (b Board) String() string {
	s := make([]player, 9)
	for i := range s {
		s[i] = b.Space((XTopLeft << uint(i)))
	}
	return fmt.Sprintf(`
		 %s | %s | %s
		__________
		 %s | %s | %s
		__________
		 %s | %s | %s
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
	switch b & (pos | pos<<PlayerShift) {
	case pos:
		return 'X'
	case pos << PlayerShift:
		return 'O'
	case pos | pos<<PlayerShift:
		panic("How did you play 2 players on tile!")
	default:
		return ' '
	}
}

func (b *Board) Move(pos Board) bool {
	pos = pos & ClearBoard
	if !(pos&XTopLeft == XTopLeft ||
		pos&XTopCenter == XTopCenter ||
		pos&XTopRight == XTopRight ||
		pos&XMiddleLeft == XMiddleLeft ||
		pos&XMiddleCenter == XMiddleCenter ||
		pos&XMiddleRight == XMiddleRight ||
		pos&XBottomLeft == XBottomLeft ||
		pos&XBottomCenter == XBottomCenter ||
		pos&XBottomRight == XBottomRight ||
		pos&OTopLeft == OTopLeft ||
		pos&OTopCenter == OTopCenter ||
		pos&OTopRight == OTopRight ||
		pos&OMiddleLeft == OMiddleLeft ||
		pos&OMiddleCenter == OMiddleCenter ||
		pos&OMiddleRight == OMiddleRight ||
		pos&OBottomLeft == OBottomLeft ||
		pos&OBottomCenter == OBottomCenter ||
		pos&OBottomRight == OBottomRight) {
		log.Println("Invalid position %b", pos)
		return false
	}

	if pos&IsX > 0 && *b&Player == 0 {
		log.Printf("PlayerMove:%b CoturrentPlayer:%b", pos&IsX, *b&Player)
		return false
	}

	cmp := pos
	if cmp > XBottomRight {
		cmp = cmp >> PlayerShift
	}

	if (*b & (cmp | cmp<<PlayerShift)) > 0 {
		log.Printf("Invalid move move taken [%b] [%b]\n", *b, pos)
		return false
	}

	*b = (*b | pos) ^ Player
	return true
}

const (
	BoardStart Board = 1 << (31 - iota)
	Player

	None Board = 1 << (iota)
	XTopLeft
	XTopCenter
	XTopRight
	XMiddleLeft
	XMiddleCenter
	XMiddleRight
	XBottomLeft
	XBottomCenter
	XBottomRight
	OTopLeft
	OTopCenter
	OTopRight
	OMiddleLeft
	OMiddleCenter
	OMiddleRight
	OBottomLeft
	OBottomCenter
	OBottomRight
	IsO        Board = OTopLeft | OMiddleLeft | OBottomLeft | OTopCenter | OMiddleCenter | OBottomCenter | OTopRight | OMiddleRight | OBottomRight
	IsX        Board = XTopLeft | XMiddleLeft | XBottomLeft | XTopCenter | XMiddleCenter | XBottomCenter | XTopRight | XMiddleRight | XBottomRight
	ClearBoard Board = XTopLeft | XMiddleLeft | XBottomLeft | XTopCenter | XMiddleCenter | XBottomCenter | XTopRight | XMiddleRight | XBottomRight | OTopLeft | OMiddleLeft | OBottomLeft | OTopCenter | OMiddleCenter | OBottomCenter | OTopRight | OMiddleRight | OBottomRight
	BoardMask  Board = 0xFFFFFFF ^ ClearBoard
	XWin1      Board = XTopLeft | XTopRight | XTopCenter
	XWin2      Board = XWin1 << 3
	XWin3      Board = XWin1 << 6
	XWin4      Board = XTopLeft | XMiddleLeft | XBottomLeft
	XWin5      Board = XWin4 << 3
	XWin6      Board = XWin4 << 6
	XWin7      Board = XTopLeft | XMiddleCenter | XBottomRight
	XWin8      Board = XTopRight | XMiddleCenter | XBottomLeft
	OWin1      Board = XWin1 << PlayerShift
	OWin2      Board = XWin1 << (PlayerShift + 3)
	OWin3      Board = XWin1 << (PlayerShift + 6)
	OWin4      Board = XWin4 << PlayerShift
	OWin5      Board = XWin4 << (PlayerShift + 3)
	OWin6      Board = XWin4 << (PlayerShift + 6)
	OWin7      Board = OTopLeft | OMiddleCenter | OBottomRight
	OWin8      Board = OTopRight | OMiddleCenter | OBottomLeft
)

var winner = map[Board]player{
	XWin1: 'X',
	XWin2: 'X',
	XWin3: 'X',
	XWin4: 'X',
	XWin5: 'X',
	XWin6: 'X',
	XWin7: 'X',
	XWin8: 'X',
	OWin1: 'O',
	OWin2: 'O',
	OWin3: 'O',
	OWin4: 'O',
	OWin5: 'O',
	OWin6: 'O',
	OWin7: 'O',
	OWin8: 'O',
}

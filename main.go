package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	moves := make(chan uint)
	go func() {
		for {
			fmt.Println("What is your move?")
			data, _, _ := r.ReadLine()
			move, err := strconv.Atoi(string(data))
			if err != nil || move < 0 || move > 8 {
				fmt.Println(err, "Please try again enter a number between 0 and 8")
				continue
			}
			moves <- uint(move)
		}
	}()
	b := None
	b.Play(moves)
}

package main

import "fmt"

func main() {
	// fmt.Printf("%b %s\n", XTopLeft, "XTopLeft")
	// fmt.Printf("%b %s\n", XMiddleLeft, "XMiddleLeft")
	// fmt.Printf("%b %s\n", XBottomLeft, "XBottomLeft")
	// fmt.Printf("%b %s\n", XTopCenter, "XTopCenter")
	// fmt.Printf("%b %s\n", XMiddleCenter, "XMiddleCenter")
	// fmt.Printf("%b %s\n", XBottomCenter, "XBottomCenter")
	// fmt.Printf("%b %s\n", XTopRight, "XTopRight")
	// fmt.Printf("%b %s\n", XMiddleRight, "XMiddleRight")
	// fmt.Printf("%b %s\n", XBottomRight, "XBottomRight")
	// fmt.Printf("%b %s\n", OTopLeft, "OTopLeft")
	// fmt.Printf("%b %s\n", OMiddleLeft, "OMiddleLeft")
	// fmt.Printf("%b %s\n", OBottomLeft, "OBottomLeft")
	// fmt.Printf("%b %s\n", OTopCenter, "OTopCenter")
	// fmt.Printf("%b %s\n", OMiddleCenter, "OMiddleCenter")
	// fmt.Printf("%b %s\n", OBottomCenter, "OBottomCenter")
	// fmt.Printf("%b %s\n", OTopRight, "OTopRight")
	// fmt.Printf("%b %s\n", OMiddleRight, "OMiddleRight")
	// fmt.Printf("%b %s\n", OBottomRight, "OBottomRight")
	// fmt.Printf("%b %s\n", ClearBoard, "ClearBoard")
	// fmt.Printf("%b %s\n", BoardMask, "BoardMask")
	// fmt.Printf("%b\n", XTopLeft|XMiddleLeft|XBottomLeft|XTopCenter|XMiddleCenter|XBottomCenter|XTopRight|XMiddleRight|XBottomRight|OTopLeft|OMiddleLeft|OBottomLeft|OTopCenter|OMiddleCenter|OBottomCenter|OTopRight|OMiddleRight|OBottomRight)

	board := None
	fmt.Printf("%s\n", board)
	fmt.Println(board.Move(OTopLeft | XTopRight))
	fmt.Printf("%s\n", board)

	// fmt.Printf("%b %b\n", (XBottomLeft ^ OBottomLeft), (XBottomLeft | XBottomLeft<<9))
	// fmt.Printf("%b\n", XBottomRight)
	// fmt.Printf("%b\n", OBottomRight)
	// fmt.Printf("%b\n", OBottomRight^XBottomRight^OBottomRight)
	// fmt.Printf("%b\n", OBottomRight^XBottomRight^OBottomRight)
}

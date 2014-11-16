package main

import (
	"github.com/nsf/termbox-go"
)

const chars = "nnnnnnnnnbbbbbbbbbuuuuuuuuuBBBBBBBBB"

func nextChar(current int) int {
	current++
	if current >= len(chars) {
		return 0
	}
	return current
}

func printCombinationsTable(sx, sy int, attrs []termbox.Attribute) {
	var bg termbox.Attribute
	currentChar := 0
	y := sy

	allAttrs := []termbox.Attribute{
		0,
		termbox.AttrBold,
		termbox.AttrUnderline,
		termbox.AttrBold | termbox.AttrUnderline,
	}

	drawLine := func() {
		x := sx
		for _, a := range allAttrs {
			for c := termbox.ColorDefault; c <= termbox.ColorWhite; c++ {
				fg := a | c
				termbox.SetCell(x, y, rune(chars[currentChar]), fg, bg)
				currentChar = nextChar(currentChar)
				x++
			}
		}
	}

	for _, a := range attrs {
		for c := termbox.ColorDefault; c <= termbox.ColorWhite; c++ {
			bg = a | c
			drawLine()
			y++
		}
	}
}

func printWide(x, y int, s string) {
	red := false
	for _, r := range s {
		c := termbox.ColorDefault
		if red {
			c = termbox.ColorRed
		}
		termbox.SetCell(x, y, r, termbox.ColorDefault, c)
		x += 2
		red = !red
	}
}

const helloWorld = "こんにちは世界"

func drawAll() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	printCombinationsTable(1, 1, []termbox.Attribute{
		0,
		termbox.AttrBold,
	})
	printCombinationsTable(2+len(chars), 1, []termbox.Attribute{
		termbox.AttrReverse,
	})
	printWide(2+len(chars), 11, helloWorld)
	termbox.Flush()
}

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	drawAll()
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			}
		case termbox.EventResize:
			drawAll()
		}
	}
}

package prettyprint

import (
	"bytes"
	"sort"
)

type zigzagGrid struct {
	Char rune
	X    int
	Y    int
}

func ZigzagView(s string, row int) string {
	// 计算每个字符的 x y 坐标然后排序

	if row == 1 {
		return s
	}

	x := 0
	y := 0

	moveDownStep := row - 1
	moveUpStep := 0

	array := make([]zigzagGrid, 0, len(s))

	for _, char := range s {
		grid := zigzagGrid{
			Char: char,
			X:    x,
			Y:    y,
		}

		array = append(array, grid)

		if moveDownStep > 0 {
			y -= 1
			moveDownStep--

			if moveDownStep == 0 {
				moveUpStep = row - 1
			}
		} else if moveUpStep > 0 {
			y += 1
			x += 1
			moveUpStep--

			if moveUpStep == 0 {
				moveDownStep = row - 1
			}
		}
	}

	sort.Slice(array, func(i, j int) bool {
		a := array[i]
		b := array[j]

		if a.Y > b.Y {
			return true
		}

		if a.Y < b.Y {
			return false
		}

		return a.X < b.X
	})

	buf := new(bytes.Buffer)
	for _, elm := range array {
		buf.WriteRune(elm.Char)

	}

	return buf.String()
}

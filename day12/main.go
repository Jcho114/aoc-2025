package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func main() {
	boxes := processInput()

	wg := sync.WaitGroup{}
	var res atomic.Int64
	var count atomic.Int64

	for _, box := range boxes {
		b := box
		wg.Add(1)

		go (func() {
			if b.canFitPresents() {
				res.Add(1)
			}
			count.Add(1)
			currCount := count.Load()
			wg.Done()
			fmt.Printf("Finished Processing %d/%d Boxes\n", currCount, len(boxes))
		})()
	}

	wg.Wait()
	fmt.Println(res.Load())
}

type Present struct {
	Bitmap          Bitmap
	Area            int
	Transformations []Bitmap
}

func (present *Present) rotateClockwise() {
	present.transpose()
	present.flipHorizontal()
}

func (present *Present) transpose() {
	for i := 0; i < len(present.Bitmap); i++ {
		for j := i + 1; j < len(present.Bitmap[0]); j++ {
			present.Bitmap[i][j], present.Bitmap[j][i] =
				present.Bitmap[j][i], present.Bitmap[i][j]
		}
	}
}

func (present *Present) flipVertical() {
	rows := len(present.Bitmap)
	for i := 0; i < rows/2; i++ {
		present.Bitmap[i], present.Bitmap[rows-i-1] =
			present.Bitmap[rows-i-1], present.Bitmap[i]
	}
}

func (present *Present) flipHorizontal() {
	rows := len(present.Bitmap)
	cols := len(present.Bitmap[0])
	for i := 0; i < cols/2; i++ {
		for j := 0; j < rows; j++ {
			present.Bitmap[j][i], present.Bitmap[j][cols-i-1] =
				present.Bitmap[j][cols-i-1], present.Bitmap[j][i]
		}
	}
}

func bitmapKey(b Bitmap) string {
	var sb strings.Builder
	for _, row := range b {
		sb.WriteString(string(row))
		sb.WriteByte('\n')
	}
	return sb.String()
}

func allTransformations(present *Present) []Bitmap {
	transformations := []Bitmap{}
	seen := map[string]bool{}

	copyPresent := func(p *Present) *Present {
		bm := make([][]rune, len(p.Bitmap))
		for i := range p.Bitmap {
			bm[i] = make([]rune, len(p.Bitmap[i]))
			copy(bm[i], p.Bitmap[i])
		}
		return &Present{Bitmap: bm}
	}

	ops := []func(*Present){
		func(p *Present) {},
		func(p *Present) { p.rotateClockwise() },
		func(p *Present) { p.rotateClockwise(); p.rotateClockwise() },
		func(p *Present) { p.rotateClockwise(); p.rotateClockwise(); p.rotateClockwise() },
		func(p *Present) { p.flipVertical() },
		func(p *Present) { p.flipHorizontal() },
		func(p *Present) { p.rotateClockwise(); p.flipVertical() },
		func(p *Present) { p.rotateClockwise(); p.flipHorizontal() },
	}

	for _, op := range ops {
		t := copyPresent(present)
		op(t)
		key := bitmapKey(t.Bitmap)
		if !seen[key] {
			seen[key] = true
			transformations = append(transformations, t.Bitmap)
		}
	}

	return transformations
}

func (present *Present) prettyPrint() {
	for _, row := range present.Bitmap {
		fmt.Println(row)
	}
	fmt.Println()
}

type Coord struct {
	Row, Col int
}

type PresentId int
type Bitmap [][]rune

const (
	Filled = '#'
	Empty  = '.'
)

type Box struct {
	Width, Length int
	Quota         []int
	Presents      []Present
	Bitmap        Bitmap
	Area          int
}

func (box *Box) validPresentPlacements(rowStart int, colStart int, id PresentId) []Bitmap {
	valid := []Bitmap{}

	for _, transformation := range box.Presents[id].Transformations {
		if box.canAddPresent(rowStart, colStart, transformation) {
			valid = append(valid, transformation)
		}
	}

	return valid
}

func (box *Box) canAddPresent(rowStart int, colStart int, bitmap Bitmap) bool {
	if rowStart+len(bitmap) > len(box.Bitmap) || colStart+len(bitmap[0]) > len(box.Bitmap[0]) {
		return false
	}

	for row := rowStart; row < rowStart+len(bitmap); row++ {
		for col := colStart; col < colStart+len(bitmap[0]); col++ {
			val := bitmap[row-rowStart][col-colStart]
			if val == Filled && box.Bitmap[row][col] == Filled {
				return false
			}
		}
	}
	return true
}

func (box *Box) addPresent(rowStart int, colStart int, id PresentId, bitmap Bitmap) {
	box.Quota[id] -= 1
	for row := rowStart; row < rowStart+len(bitmap); row++ {
		for col := colStart; col < colStart+len(bitmap[0]); col++ {
			val := bitmap[row-rowStart][col-colStart]
			if val == Filled {
				box.Bitmap[row][col] = Filled
			}
		}
	}
}

func (box *Box) prettyPrintBitmap() {
	for _, row := range box.Bitmap {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func (box *Box) removePresent(rowStart int, colStart int, id PresentId, bitmap Bitmap) {
	box.Quota[id] += 1
	for row := rowStart; row < rowStart+len(bitmap); row++ {
		for col := colStart; col < colStart+len(bitmap[0]); col++ {
			val := bitmap[row-rowStart][col-colStart]
			if val == Filled {
				box.Bitmap[row][col] = Empty
			}
		}
	}
}

func (box *Box) canFitPresents() bool {
	necessaryArea := 0
	numPresents := 0
	for id, amount := range box.Quota {
		numPresents += amount
		necessaryArea += box.Presents[id].Area * amount
	}

	if box.Area < necessaryArea {
		return false
	}

	if numPresents*9 < box.Area {
		return true
	}

	return box.canFitPresentsHelper(0, make(map[string]bool))
}

func (box *Box) canFitPresentsHelper(curr PresentId, memo map[string]bool) bool {
	if curr >= PresentId(len(box.Presents)) {
		return true
	}

	key := bitmapKey(box.Bitmap) + "|" + fmt.Sprint(box.Quota[curr:])
	if memo[key] {
		return false
	}

	if box.Quota[curr] == 0 {
		return box.canFitPresentsHelper(curr+1, memo)
	}

	for row := 0; row < box.Length; row++ {
		for col := 0; col < box.Width; col++ {
			placements := box.validPresentPlacements(row, col, curr)
			if len(placements) == 0 {
				continue
			}

			for _, placement := range placements {
				box.addPresent(row, col, curr, placement)

				if box.Quota[curr] > 0 {
					if box.canFitPresentsHelper(curr, memo) {
						return true
					}
				} else {
					if box.canFitPresentsHelper(curr+1, memo) {
						return true
					}
				}

				box.removePresent(row, col, curr, placement)
			}
		}
	}

	memo[key] = true
	return false
}

func processInput() []Box {
	presents := []Present{}
	boxes := []Box{}

	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	presentMode := true
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "x") {
			presentMode = false
		}

		if presentMode {
			present := Present{}
			scanner.Scan()
			line = scanner.Text()
			for line != "" {
				present.Bitmap = append(present.Bitmap, []rune(line))
				scanner.Scan()
				line = scanner.Text()
			}

			area := 0
			for _, row := range present.Bitmap {
				for _, val := range row {
					if val == Filled {
						area += 1
					}
				}
			}
			present.Area = area
			present.Transformations = allTransformations(&present)

			presents = append(presents, present)
		} else {
			split := strings.Split(line, ": ")

			dimensions := strings.Split(split[0], "x")
			width, err := strconv.Atoi(dimensions[0])
			if err != nil {
				panic(err)
			}
			length, err := strconv.Atoi(dimensions[1])
			if err != nil {
				panic(err)
			}

			quota := []int{}
			split = strings.Split(split[1], " ")
			for _, raw := range split {
				val, err := strconv.Atoi(raw)
				if err != nil {
					panic(err)
				}
				quota = append(quota, val)
			}

			bitmap := make([][]rune, length)
			for i := range bitmap {
				bitmap[i] = make([]rune, width)
				for j := range bitmap[i] {
					bitmap[i][j] = Empty
				}
			}

			box := Box{
				Width:    width,
				Length:   length,
				Quota:    quota,
				Presents: presents,
				Bitmap:   bitmap,
				Area:     width * length,
			}
			boxes = append(boxes, box)
		}
	}

	return boxes
}

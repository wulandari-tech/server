package game

import (
	"encoding/json"
)

type Linex [8]Pos

func (l *Linex) At(x Pos) Pos {
	return l[x-1]
}

func (l *Linex) Set(x Pos, val Pos) {
	l[x-1] = val
}

func (l *Linex) Cover(op Linex) *Linex {
	for i, v := range op {
		if v > 0 {
			l[i] = v
		}
	}
	return l
}

func (l *Linex) Len() int {
	for i := 7; i >= 0; i-- {
		if l[i] > 0 {
			return i + 1
		}
	}
	return 0
}

func (l *Linex) Copy(pos, num Pos) (dst Linex) {
	var x1, x2 Pos
	if num >= 0 {
		x1, x2 = pos-1, pos-1+num
	} else {
		x1, x2 = pos+num, pos
	}
	copy(dst[x1:x2], l[x1:x2])
	return
}

func (l *Linex) CopyL(num Pos) (dst Linex) {
	copy(dst[:num], l[:num])
	return
}

func (l *Linex) CopyR5(num Pos) (dst Linex) {
	copy(dst[5-num:5], l[5-num:5])
	return
}

func (l Linex) MarshalJSON() ([]byte, error) {
	return json.Marshal(l[:l.Len()])
}

// (1 ,1) symbol is on left top corner

// Ultra Hot 3x3 bet lines
var BetLinesHot3 = []Linex{
	{2, 2, 2}, // 1
	{1, 1, 1}, // 2
	{3, 3, 3}, // 3
	{1, 2, 3}, // 4
	{3, 2, 1}, // 5
}

// Sizzling Hot bet lines
var BetLinesHot5 = []Linex{
	{2, 2, 2, 2, 2}, // 1
	{1, 1, 1, 1, 1}, // 2
	{3, 3, 3, 3, 3}, // 3
	{1, 2, 3, 2, 1}, // 4
	{3, 2, 1, 2, 3}, // 5
}

// Megajack 21 bet lines
var BetLinesMgj = []Linex{
	{2, 2, 2, 2, 2}, // 1
	{1, 1, 1, 1, 1}, // 2
	{3, 3, 3, 3, 3}, // 3
	{1, 2, 3, 2, 1}, // 4
	{3, 2, 1, 2, 3}, // 5
	{1, 1, 2, 3, 3}, // 6
	{3, 3, 2, 1, 1}, // 7
	{2, 1, 2, 3, 2}, // 8
	{2, 3, 2, 1, 2}, // 9
	{1, 2, 1, 2, 1}, // 10
	{3, 2, 3, 2, 3}, // 11
	{1, 2, 2, 2, 2}, // 12
	{3, 2, 2, 2, 2}, // 13
	{2, 2, 1, 1, 1}, // 14
	{2, 2, 3, 3, 3}, // 15
	{2, 1, 1, 1, 1}, // 16
	{2, 3, 3, 3, 3}, // 17
	{1, 1, 1, 1, 2}, // 18
	{3, 3, 3, 3, 2}, // 19
	{3, 3, 2, 3, 3}, // 20
	{1, 1, 2, 1, 1}, // 21
}

// Novomatic 9 bet lines (old versions of games)
var BetLinesNvm9 = []Linex{
	{2, 2, 2, 2, 2}, // 1
	{1, 1, 1, 1, 1}, // 2
	{3, 3, 3, 3, 3}, // 3
	{1, 2, 3, 2, 1}, // 4
	{3, 2, 1, 2, 3}, // 5
	{2, 3, 3, 3, 2}, // 6
	{2, 1, 1, 1, 2}, // 7
	{3, 3, 2, 1, 1}, // 8
	{1, 1, 2, 3, 3}, // 9
}

// Novomatic 10 bet lines (deluxe versions of games)
var BetLinesNvm10 = []Linex{
	{2, 2, 2, 2, 2}, // 1
	{1, 1, 1, 1, 1}, // 2
	{3, 3, 3, 3, 3}, // 3
	{1, 2, 3, 2, 1}, // 4
	{3, 2, 1, 2, 3}, // 5
	{2, 3, 3, 3, 2}, // 6
	{2, 1, 1, 1, 2}, // 7
	{3, 3, 2, 1, 1}, // 8
	{1, 1, 2, 3, 3}, // 9
	{3, 2, 2, 2, 1}, // 10
}

// Novomatic 20 bet lines (new games)
var BetLinesNvm20 = []Linex{
	{2, 2, 2, 2, 2}, // 1
	{1, 1, 1, 1, 1}, // 2
	{3, 3, 3, 3, 3}, // 3
	{1, 2, 3, 2, 1}, // 4
	{3, 2, 1, 2, 3}, // 5
	{1, 1, 2, 3, 3}, // 6
	{3, 3, 2, 1, 1}, // 7
	{2, 1, 1, 1, 2}, // 8
	{2, 3, 3, 3, 2}, // 9
	{1, 2, 2, 2, 1}, // 10
	{3, 2, 2, 2, 3}, // 11
	{2, 1, 2, 3, 2}, // 12
	{2, 3, 2, 1, 2}, // 13
	{1, 3, 1, 3, 1}, // 14
	{3, 1, 3, 1, 3}, // 15
	{1, 2, 1, 2, 1}, // 16
	{3, 2, 3, 2, 3}, // 17
	{1, 3, 3, 3, 3}, // 18
	{3, 1, 1, 1, 1}, // 19
	{2, 2, 1, 2, 2}, // 20
}

// Novomatic 40 bet lines (screen 5x4)
var BetLinesNvm40 = []Linex{
	{1, 1, 1, 1, 1}, // 1
	{2, 2, 2, 2, 2}, // 2
	{3, 3, 3, 3, 3}, // 3
	{4, 4, 4, 4, 4}, // 4
	{1, 2, 3, 2, 1}, // 5
	{2, 3, 4, 3, 2}, // 6
	{3, 2, 1, 2, 3}, // 7
	{4, 3, 2, 3, 4}, // 8
	{1, 1, 1, 1, 2}, // 9
	{2, 2, 2, 2, 1}, // 10
	{3, 3, 3, 3, 4}, // 11
	{4, 4, 4, 4, 3}, // 12
	{1, 2, 2, 2, 2}, // 13
	{2, 2, 2, 2, 3}, // 14
	{3, 3, 3, 3, 2}, // 15
	{4, 3, 3, 3, 3}, // 16
	{2, 1, 1, 1, 1}, // 17
	{2, 3, 3, 3, 3}, // 18
	{3, 2, 2, 2, 2}, // 19
	{3, 4, 4, 4, 4}, // 20
	{1, 1, 1, 2, 3}, // 21
	{2, 2, 2, 3, 4}, // 22
	{3, 3, 3, 2, 1}, // 23
	{4, 4, 4, 3, 2}, // 24
	{1, 2, 3, 3, 3}, // 25
	{2, 3, 4, 4, 4}, // 26
	{3, 2, 1, 1, 1}, // 27
	{4, 3, 2, 2, 2}, // 28
	{1, 1, 2, 1, 1}, // 29
	{2, 2, 1, 2, 2}, // 30
	{3, 3, 4, 3, 3}, // 31
	{4, 4, 3, 4, 4}, // 32
	{1, 2, 2, 2, 1}, // 33
	{2, 2, 3, 2, 2}, // 34
	{3, 3, 2, 3, 3}, // 35
	{4, 3, 3, 3, 4}, // 36
	{2, 1, 1, 1, 2}, // 37
	{2, 3, 3, 3, 2}, // 38
	{3, 2, 2, 2, 3}, // 39
	{3, 4, 4, 4, 3}, // 40
}

// BetSoft 25 bet lines
var BetLinesBetSoft25 = []Linex{
	{2, 2, 2, 2, 2}, // 1
	{1, 1, 1, 1, 1}, // 2
	{3, 3, 3, 3, 3}, // 3
	{1, 2, 3, 2, 1}, // 4
	{3, 2, 1, 2, 3}, // 5
	{1, 1, 2, 1, 1}, // 6
	{3, 3, 2, 3, 3}, // 7
	{2, 3, 3, 3, 2}, // 8
	{2, 1, 1, 1, 2}, // 9
	{2, 1, 2, 1, 2}, // 10
	{2, 3, 2, 3, 2}, // 11
	{1, 2, 1, 2, 1}, // 12
	{3, 2, 3, 2, 3}, // 13
	{2, 2, 1, 2, 2}, // 14
	{2, 2, 3, 2, 2}, // 15
	{1, 2, 2, 2, 1}, // 16
	{3, 2, 2, 2, 3}, // 17
	{1, 2, 3, 3, 3}, // 18
	{3, 2, 1, 1, 1}, // 19
	{1, 3, 1, 3, 1}, // 20
	{3, 1, 3, 1, 3}, // 21
	{1, 3, 3, 3, 1}, // 22
	{3, 1, 1, 1, 3}, // 23
	{1, 1, 3, 1, 1}, // 24
	{3, 3, 1, 3, 3}, // 25
}

// BetSoft 30 bet lines
var BetLinesBetSoft30 = []Linex{
	{2, 2, 2, 2, 2}, // 1
	{1, 1, 1, 1, 1}, // 2
	{3, 3, 3, 3, 3}, // 3
	{1, 2, 3, 2, 1}, // 4
	{3, 2, 1, 2, 3}, // 5
	{1, 1, 2, 1, 1}, // 6
	{3, 3, 2, 3, 3}, // 7
	{2, 3, 3, 3, 2}, // 8
	{2, 1, 1, 1, 2}, // 9
	{2, 1, 2, 1, 2}, // 10
	{2, 3, 2, 3, 2}, // 11
	{1, 2, 1, 2, 1}, // 12
	{3, 2, 3, 2, 3}, // 13
	{2, 2, 1, 2, 2}, // 14
	{2, 2, 3, 2, 2}, // 15
	{1, 2, 2, 2, 1}, // 16
	{3, 2, 2, 2, 3}, // 17
	{1, 2, 3, 3, 3}, // 18
	{3, 2, 1, 1, 1}, // 19
	{1, 3, 1, 3, 1}, // 20
	{3, 1, 3, 1, 3}, // 21
	{1, 3, 3, 3, 1}, // 22
	{3, 1, 1, 1, 3}, // 23
	{1, 1, 3, 1, 1}, // 24
	{3, 3, 1, 3, 3}, // 25
	{1, 3, 2, 1, 3}, // 26
	{3, 1, 2, 3, 1}, // 27
	{2, 1, 3, 2, 3}, // 28
	{1, 3, 2, 3, 2}, // 29
	{3, 2, 1, 1, 2}, // 30
}

// NetEnt 10 bet lines
var BetLinesNetEnt10 = []Linex{
	{2, 2, 2, 2, 2}, // 1
	{1, 1, 1, 1, 1}, // 2
	{3, 3, 3, 3, 3}, // 3
	{1, 2, 3, 2, 1}, // 4
	{3, 2, 1, 2, 3}, // 5
	{1, 1, 2, 1, 1}, // 6
	{3, 3, 2, 3, 3}, // 7
	{2, 3, 3, 3, 2}, // 8
	{2, 1, 1, 1, 2}, // 9
	{2, 1, 2, 1, 2}, // 10
}

// NetEnt 15 bet lines
var BetLinesNetEnt15 = []Linex{
	{2, 2, 2, 2, 2}, // 1
	{1, 1, 1, 1, 1}, // 2
	{3, 3, 3, 3, 3}, // 3
	{1, 2, 3, 2, 1}, // 4
	{3, 2, 1, 2, 3}, // 5
	{1, 1, 2, 1, 1}, // 6
	{3, 3, 2, 3, 3}, // 7
	{2, 3, 3, 3, 2}, // 8
	{2, 1, 1, 1, 2}, // 9
	{2, 1, 2, 1, 2}, // 10
	{2, 3, 2, 3, 2}, // 11
	{1, 2, 1, 2, 1}, // 12
	{3, 2, 3, 2, 3}, // 13
	{2, 2, 1, 2, 2}, // 14
	{2, 2, 3, 2, 2}, // 15
}

// NetEnt 20 bet lines
var BetLinesNetEnt20 = []Linex{
	{2, 2, 2, 2, 2}, // 1
	{1, 1, 1, 1, 1}, // 2
	{3, 3, 3, 3, 3}, // 3
	{1, 2, 3, 2, 1}, // 4
	{3, 2, 1, 2, 3}, // 5
	{1, 1, 2, 1, 1}, // 6
	{3, 3, 2, 3, 3}, // 7
	{2, 3, 3, 3, 2}, // 8
	{2, 1, 1, 1, 2}, // 9
	{2, 1, 2, 1, 2}, // 10
	{2, 3, 2, 3, 2}, // 11
	{1, 2, 1, 2, 1}, // 12
	{3, 2, 3, 2, 3}, // 13
	{2, 2, 1, 2, 2}, // 14
	{2, 2, 3, 2, 2}, // 15
	{1, 2, 2, 2, 1}, // 16
	{3, 2, 2, 2, 3}, // 17
	{1, 2, 3, 3, 3}, // 18
	{3, 2, 1, 1, 1}, // 19
	{1, 3, 1, 3, 1}, // 20
}

// NetEnt 40 bet lines for 5x4 screen
var BetLinesNetEnt40 = []Linex{
	{2, 2, 2, 2, 2}, // 1
	{3, 3, 3, 3, 3}, // 2
	{1, 1, 1, 1, 1}, // 3
	{4, 4, 4, 4, 4}, // 4
	{2, 3, 4, 3, 2}, // 5
	{3, 2, 1, 2, 3}, // 6
	{1, 1, 2, 3, 4}, // 7
	{4, 4, 3, 2, 1}, // 8
	{2, 1, 1, 1, 2}, // 9
	{3, 4, 4, 4, 3}, // 10
	{1, 2, 3, 4, 4}, // 11
	{4, 3, 2, 1, 1}, // 12
	{2, 1, 2, 3, 2}, // 13
	{3, 4, 3, 2, 3}, // 14
	{1, 2, 1, 2, 1}, // 15
	{4, 3, 4, 3, 4}, // 16
	{2, 3, 2, 1, 2}, // 17
	{3, 2, 3, 4, 3}, // 18
	{1, 2, 2, 2, 1}, // 19
	{4, 3, 3, 3, 4}, // 20
	{2, 2, 3, 4, 4}, // 21
	{3, 3, 2, 1, 1}, // 22
	{2, 2, 1, 2, 2}, // 23
	{3, 3, 4, 3, 3}, // 24
	{2, 3, 3, 3, 4}, // 25
	{3, 2, 2, 2, 1}, // 26
	{1, 1, 2, 1, 1}, // 27
	{4, 4, 3, 4, 4}, // 28
	{1, 2, 3, 3, 4}, // 29
	{4, 3, 2, 2, 1}, // 30
	{1, 1, 1, 2, 3}, // 31
	{4, 4, 4, 3, 2}, // 32
	{2, 1, 1, 2, 3}, // 33
	{3, 4, 4, 3, 2}, // 34
	{1, 2, 2, 3, 4}, // 35
	{4, 3, 3, 2, 1}, // 36
	{2, 1, 2, 3, 4}, // 37
	{3, 4, 3, 2, 1}, // 38
	{1, 2, 3, 4, 3}, // 39
	{4, 3, 2, 1, 2}, // 40
}

// Playtech 15 bet lines
var BetLinesPlt15 = []Linex{
	{2, 2, 2, 2, 2}, // 1
	{1, 1, 1, 1, 1}, // 2
	{3, 3, 3, 3, 3}, // 3
	{1, 2, 3, 2, 1}, // 4
	{3, 2, 1, 2, 3}, // 5
	{2, 1, 1, 1, 2}, // 6
	{2, 3, 3, 3, 2}, // 7
	{1, 1, 2, 3, 3}, // 8
	{3, 3, 2, 1, 1}, // 9
	{2, 3, 2, 1, 2}, // 10
	{2, 1, 2, 3, 2}, // 11
	{1, 2, 2, 2, 1}, // 12
	{3, 2, 2, 2, 3}, // 13
	{1, 2, 1, 2, 1}, // 14
	{3, 2, 3, 2, 3}, // 15
}

// Playtech 30 bet lines
var BetLinesPlt30 = []Linex{
	{2, 2, 2, 2, 2}, // 1
	{1, 1, 1, 1, 1}, // 2
	{3, 3, 3, 3, 3}, // 3
	{1, 2, 3, 2, 1}, // 4
	{3, 2, 1, 2, 3}, // 5
	{2, 1, 1, 1, 2}, // 6
	{2, 3, 3, 3, 2}, // 7
	{1, 1, 2, 3, 3}, // 8
	{3, 3, 2, 1, 1}, // 9
	{2, 3, 2, 1, 2}, // 10
	{2, 1, 2, 3, 2}, // 11
	{1, 2, 2, 2, 1}, // 12
	{3, 2, 2, 2, 3}, // 13
	{1, 2, 1, 2, 1}, // 14
	{3, 2, 3, 2, 3}, // 15
	{2, 2, 1, 2, 2}, // 16
	{2, 2, 3, 2, 2}, // 17
	{1, 1, 3, 1, 1}, // 18
	{3, 3, 1, 3, 3}, // 19
	{1, 3, 3, 3, 1}, // 20
	{3, 1, 1, 1, 3}, // 21
	{2, 3, 1, 3, 2}, // 22
	{2, 1, 3, 1, 2}, // 23
	{1, 3, 1, 3, 1}, // 24
	{3, 1, 3, 1, 3}, // 25
	{1, 3, 2, 1, 3}, // 26
	{3, 1, 2, 3, 1}, // 27
	{2, 1, 3, 2, 3}, // 28
	{1, 3, 2, 3, 1}, // 29
	{3, 2, 1, 1, 2}, // 30
}

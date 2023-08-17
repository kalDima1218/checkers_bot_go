package main

import (
	"math"
	"sync"
)

func _dist(a [2]int, b [2]int) [2]int {
	return [2]int{b[0] - a[0], b[1] - a[1]}
}

func _abs(a int) int {
	return int(math.Abs(float64(a)))
}

func _add(a [2]int, b [2]int) [2]int {
	return [2]int{a[0] + b[0], a[1] + b[1]}
}

func _sub(a [2]int, b [2]int) [2]int {
	return [2]int{a[0] - b[0], a[1] - b[1]}
}

func _div(a [2]int, b int) [2]int {
	return [2]int{a[0] / b, a[1] / b}
}

func _len(a [2]int, b [2]int) int {
	return int(math.Sqrt((math.Pow(float64(b[0]-a[0]), 2) + math.Pow(float64(b[1]-a[1]), 2)) / 2))
}

func _isBetwen(l [2]int, m [2]int, r [2]int) bool {
	return _len(l, m)+_len(m, r) == _len(l, r)
}

func _max(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func _min(a int, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

type Move struct {
	score float64
	game  Game
}

func newMove(score float64, game Game) Move {
	var tmp Move
	tmp.score = score
	tmp.game = game
	return tmp
}

type MapGame struct {
	mx sync.Mutex
	mp map[Game]float64
}

func newMapGame() *MapGame {
	return &MapGame{
		mp: make(map[Game]float64),
	}
}

func (mp *MapGame) get(key Game) (float64, bool) {
	mp.mx.Lock()
	defer mp.mx.Unlock()
	val, ok := mp.mp[key]
	return val, ok
}

func (mp *MapGame) insert(key Game, value float64) {
	mp.mx.Lock()
	defer mp.mx.Unlock()
	mp.mp[key] = value
}

func (mp *MapGame) clear() {
	for k := range mp.mp {
		delete(mp.mp, k)
	}
}

func _init() {
	for i := 1; i <= 7; i++ {
		POSSIBLE_TURNS[i-1] = [2]int{i, i}
		POSSIBLE_TURNS[7+i-1] = [2]int{i, -i}
		POSSIBLE_TURNS[14+i-1] = [2]int{-i, i}
		POSSIBLE_TURNS[21+i-1] = [2]int{-i, -i}
	}
}

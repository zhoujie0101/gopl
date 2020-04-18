package main

import (
	"fmt"
	"math"
	"sync"
)

type Point struct { X, Y float64 }

func Distance(p, q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p *Point) ScaleBy(factor float64) {
    p.X *= factor
    p.Y *= factor
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	fmt.Println(Distance(p, q))  //function call
	fmt.Println(p.Distance(q))  //method call

	p.ScaleBy(2)
}

var cache = struct {
    sync.Mutex
    mapping map[string]string
}{
    mapping : make(map[string]string),
}

func Lookup(key string) string {
    cache.Lock()
    v := cache.mapping[key]
    cache.Unlock()
    return v
}
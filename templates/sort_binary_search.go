package main

import "sort"

type pair struct {
	x int
	y float64
}

type pairs []pair

func (a pairs) Len() int           { return len(a) }
func (a pairs) Less(i, j int) bool { return a[i].x < a[j].x || a[i].y < a[j].y }
func (a pairs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func sortBynarySearch() {
	defer writer.Flush()

	items := pairs{{4, 74.0}, {9, 34.56}, {9, 31.31}}
	sort.Sort(sort.Reverse(items))

	for _, item := range items {
		printf("(%d, %.2f)\n", item.x, item.y)
	}

	// (9, 34.56)
	// (9, 31.31)
	// (4, 74.00)
}

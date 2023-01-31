package data

import (
	"sort"
)

type Row struct {
	timeStamp int32
	price     int32
}

func NewRow(timeStamp int32, price int32) Row {
	return Row{timeStamp, price}
}

type Assets struct {
	data []Row
}

func NewAssets() *Assets {
	return &Assets{[]Row{}}
}

func (assets *Assets) Query(from int32, to int32) int32 {
	start := sort.Search(len(assets.data), func(i int) bool { return assets.data[i].timeStamp >= from })
	end := sort.Search(len(assets.data), func(i int) bool { return assets.data[i].timeStamp > to })
	if end < start || end == 0 || start == len(assets.data) {
		return 0
	}
	targets := assets.data[start:end]
	var sum int32 = 0
	for _, asset := range targets {
		sum += asset.price
	}
	return sum / int32(len(targets))
}

func (assets *Assets) Insert(row Row) {
	i := sort.Search(len(assets.data), func(i int) bool { return assets.data[i].timeStamp > row.timeStamp })
	if len(assets.data) == 0 || i == len(assets.data) {
		assets.data = append(assets.data, row)
		return
	}
	assets.data = append(assets.data[:i+1], assets.data[i:]...)
	assets.data[i] = row
}

package gildedrose

import "fmt"

// Item describes an item sold by the Gilded Rose Inn
type Item struct {
	name    string
	sellin    int
	quality int
}

// String outputs a string representation of an Item
func (i *Item) String() string {
	return fmt.Sprintf("%s: %d days left, quality is %d", i.name, i.sellin, i.quality)
}

// New creates a new Item
func New(name string, days, quality int) *Item {
	return &Item{
		name:    name,
		sellin:    days,
		quality: quality,
	}
}


package gildedrose

type GildedRose struct {
	items []*Item
}

func NewGildedRose(items []*Item) *GildedRose {
	return &GildedRose{items: items}
}

// UpdateQuality ages the item by a day, and updates the quality of the item
func (gildedRose *GildedRose) UpdateQuality(items ...*Item) {
	for _, item := range items {
		if item.name != "Aged Brie" && item.name != "Backstage passes to a TAFKAL80ETC concert" {
			if item.quality > 0 {
				if item.name != "Sulfuras, Hand of Ragnaros" {
					item.quality = item.quality - 1
				}
			}
		} else {
			if item.quality < 50 {
				item.quality = item.quality + 1
				if item.name == "Backstage passes to a TAFKAL80ETC concert" {
					if item.sellin < 11 {
						if item.quality < 50 {
							item.quality = item.quality + 1
						}
					}
					if item.sellin < 6 {
						if item.quality < 50 {
							item.quality = item.quality + 1
						}
					}
				}
			}
		}

		if item.name != "Sulfuras, Hand of Ragnaros" {
			item.sellin = item.sellin - 1
		}

		if item.sellin < 0 {
			if item.name != "Aged Brie" {
				if item.name != "Backstage passes to a TAFKAL80ETC concert" {
					if item.quality > 0 {
						if item.name != "Sulfuras, Hand of Ragnaros" {
							item.quality = item.quality - 1
						}
					}
				} else {
					item.quality = item.quality - item.quality
				}
			} else {
				if item.quality < 50 {
					item.quality = item.quality + 1
				}
			}
		}
	}
}
package gildedrose

import (
	"fmt"
	"strings"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)



type GildedRoseTest struct {
    items []*Item
}

func NewGildedRoseTest() *GildedRoseTest {
    return &GildedRoseTest{
        items: []*Item{
            {name: "+5 Dexterity Vest", sellin: 10, quality: 20},
            {name: "Aged Brie", sellin: 2, quality: 0},
            {name: "Elixir of the Mongoose", sellin: 5, quality: 7},
            {name: "Sulfuras, Hand of Ragnaros", sellin: 0, quality: 80},
            {name: "Sulfuras, Hand of Ragnaros", sellin: -1, quality: 80},
            {name: "Backstage passes to a TAFKAL80ETC concert", sellin: 15, quality: 20},
            {name: "Backstage passes to a TAFKAL80ETC concert", sellin: 10, quality: 49},
            {name: "Backstage passes to a TAFKAL80ETC concert", sellin: 5, quality: 49},
        },
    }
}

func (g *GildedRoseTest) PrintItems() string {
    var itemLog strings.Builder
    itemLog.WriteString("name, sellIn, quality\n")
    for _, item := range g.items {
        itemLog.WriteString(fmt.Sprintf("%s, %d, %d\n", item.name, item.sellin, item.quality))
    }
    itemLog.WriteString("\n")
    return itemLog.String()
}

func TestGildedRoseFirstUpdate(t *testing.T) {
    t.Skip("Waiting for demo")
    g := NewGildedRoseTest()
    log := "Before Update\n"
    log += g.PrintItems()
    gildedRose := NewGildedRose(g.items)
    gildedRose.UpdateQuality()
    log += "After Update\n"
    log += g.PrintItems()
    approvals.VerifyString(t, log)
}


package main

import "fmt"

func main() {
	var i1 Item
	fmt.Println(i1)
	fmt.Printf("%#v\n", i1)

	i2 := Item{1, 2}
	fmt.Printf("i2: %#v\n", i2)

	i3 := Item{ // Here you don't need to obey the order of type Item, you can omit one of them too.
		Y: 10,
		X: 20,
	}
	fmt.Printf("i3: %#v\n", i3)
	fmt.Println(newItem(10, 20))
	fmt.Println(newItem(10, -20))

	i3.Move(100, 200)
	fmt.Printf("i3 (move): %#v\n", i3)

	p1 := Player{
		Name: "Parzival",
		Item: Item{500, 300},
	}
	fmt.Printf("p1: %#v\n", p1)
	fmt.Printf("p1.X: %#v\n", p1.X)
	fmt.Printf("p1.Item.X: %#v\n", p1.Item.X)
	fmt.Println()

	ms := []mover{
		&i1,
		&p1,
		&i2,
	}
	moveAll(ms, 0, 0)
	for _, m := range ms {
		fmt.Println(m)
	}

	k := Jade
	fmt.Println("k:", k)
	fmt.Println("key:", Key(17))

	p1.FoundKey(Jade)
	fmt.Println(p1.Keys)
	p1.FoundKey(Copper)
	fmt.Println(p1.Keys)

}

/*
	Exercise

- Add a "Keys" field to player which is a slice of Key
- Add a "FoundKey(k Key) error" method to player which will add k to Key if it's not there
- Err if k is not one of the known keys
*/
func (p *Player) FoundKey(k Key) error {
	if k < Jade || k >= invalidKey {
		return fmt.Errorf("invalid key: %#v", k)
	}
	if !containsKey(p.Keys, k) {
		p.Keys = append(p.Keys, k)
	}
	return nil
}

func containsKey(keys []Key, k Key) bool {
	for _, k2 := range keys {
		if k2 == k {
			return true
		}
	}
	return false
}

func (k Key) String() string {
	switch k {
	case Jade:
		return "jade"
	case Copper:
		return "copper"
	case Crystal:
		return "crystal"
	}

	return fmt.Sprintf("<Key %d>", k)
}

// Rule of thumb: Accept interfaces, return types
func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

/*
i is called "the receiver"
If you want to mutate, use pointer receiver
One thing have to be explicity clear, it isn't inheritance it's embedding
*/
func (i *Item) Move(x, y int) { // if remove pointer receiver * the values won't change
	i.X = x
	i.Y = y
}

func newItem(x, y int) (*Item, error) { // The * is called "pointer receiver"
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)
	}

	i := Item{
		X: x,
		Y: y,
	}
	// The Go compiler does escape analysis and will allocation i on the heap
	return &i, nil
}

const (
	maxX = 1_000
	maxY = 600
)

// Go's version of enum
const (
	Jade Key = iota + 1
	Copper
	Crystal
	invalidKey // internal (not exported)
)

type Key byte

type mover interface {
	Move(x, y int)
}

type Player struct {
	Name string
	Item // Embed Item
	Keys []Key
}

// Item is an item in the game
type Item struct {
	X int
	Y int
}

/*
TODO Implement sortByDistance(players []Player, x, y int)
*/

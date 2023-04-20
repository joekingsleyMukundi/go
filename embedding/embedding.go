package main

import "fmt"

const (
	Small = iota
	Medium
	Large
)

const (
	Ground = iota
	Air
)

type BelstSize int
type Shipping int

func (b BelstSize) String() string {
	return []string{"Small", "Medium", "large"}[b]
}

func (s Shipping) String() string {
	return []string{"Ground", "Air"}[s]
}

type Convayor interface {
	Convay() BelstSize
}
type Shipper interface {
	Ship() Shipping
}
type WareHouseAutomator interface {
	Convayor
	Shipper
}

type spamMail struct {
	amount int
}

func (s *spamMail) Ship() Shipping {
	return Air
}
func (s *spamMail) Convay() BelstSize {
	return Small
}
func automate(item WareHouseAutomator) {
	fmt.Printf("Convey %V on %v convayor\n", item, item.Convay())
	fmt.Printf("Ship %v via %v\n", item, item.Ship())
}

type ToxicWaste struct {
	weight int
}

func (t *ToxicWaste) Ship() Shipping {
	return Ground
}
func main() {

	mail := spamMail{40000}

	automate(&mail)

	// waste := ToxicWaste{3000}
	// automate(waste)

	// Using embedded interfaces to ensure only items that
	// implement both Conveyor and Shipper can be automated
	// automate(&ToxicWaste{300}) // Won't work!
}

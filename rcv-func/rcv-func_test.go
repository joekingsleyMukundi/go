package main

import "testing"

func CreateParkingLot() ParkingLot {
	lot := ParkingLot{spaces: make([]Space, 4)}
	return lot
}

func TestCreateParkingLot(t *testing.T) {
	lot := CreateParkingLot()
	if len(lot.spaces) != 4 {
		t.Errorf("incorrect value of spaces got %v want %v", len(lot.spaces), 4)
	}
}

func TestOccupySpace(t *testing.T) {
	lot := CreateParkingLot()
	lot.occupySpace(1)
	if !lot.spaces[0].occupied {
		t.Errorf("incorrect value for occupy space got %v want %v", lot.spaces[0].occupied, true)
	}
}

func TestVacateSpace(t *testing.T) {
	lot := CreateParkingLot()
	lot.occupySpace(1)
	lot.occupySpace(2)
	lot.occupySpace(3)

	lot.vacateSpace(2)
	if lot.spaces[1].occupied {
		t.Errorf("incorrect value for occupy space got %v want %v", lot.spaces[0].occupied, true)
	}
}

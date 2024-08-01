package main

type ParkingSystem struct {
	car []int
}

func Constructor(big, medium, small int) ParkingSystem {
	return ParkingSystem{car: []int{big, medium, small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.car[carType-1] != 0{
		this.car[carType-1] -=1
		return true
	}
	return false
}

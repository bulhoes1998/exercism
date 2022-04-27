package speed

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	newCar := Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
		distance:     0,
	}

	return newCar
}

type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	newTrack := Track{
		distance: distance,
	}

	return newTrack
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	switch {
	case car.battery >= car.batteryDrain:
		car.distance = car.distance + car.speed
		car.battery = car.battery - car.batteryDrain
	default:
		return car
	}

	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	switch {
	case (track.distance/car.speed)*car.batteryDrain <= car.battery:
		return true
	default:
		return false
	}
}

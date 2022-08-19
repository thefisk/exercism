package speed

// TODO: define the 'Car' type struct

type Car struct {
    battery int
    batteryDrain int
    speed int
    distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car {
        speed: speed,
        battery: 100,
        batteryDrain: batteryDrain,
    }
}

// TODO: define the 'Track' type struct

type Track struct {
    distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track {
        distance: distance,
    }
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
    var distance int
    var batteryUsed int
    if car.batteryDrain > car.battery {
        batteryUsed = 0
        distance = 0
    } else {
    	distance = car.speed
        batteryUsed = car.batteryDrain
    }
	return Car {
        distance: car.distance + distance,
        battery: car.battery - batteryUsed,
        speed: car.speed,
        batteryDrain: car.batteryDrain,
    }
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	return (car.battery / car.batteryDrain * car.speed) >= track.distance
}

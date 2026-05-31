// WPU Coding challenge
// 21/366

// Question :

// After a hard quarter in the office you decide to get some rest on a vacation. So you will book a flight for you and your girlfriend and try to leave all the mess behind you.
// You will need a rental car in order for you to get around in your vacation. The manager of the car rental makes you some good offers.
// Every day you rent the car costs $40. If you rent the car for 7 or more days, you get $50 off your total. Alternatively, if you rent the car for 3 or more days, you get $20 off your total.
// Write a code that gives out the total amount for different days(d).

// Answer :
package soal


func RentalCarCost(day int) (res int) {

	if day < 3 {
		res = 40 * day
	} else if day >= 3 && day < 7 {
		res = 40*day - 20
	} else {
		res = 40*day - 50
	}

	return res

}

// Aslinya pakai js gada go :
// let rentalCarCost = (d) => d < 3 ? 40 * d : d >= 3 && d < 7 ? 40 * d - 20 : 40 * d - 50
package main

//Requrement
// A program that computes mobile phone cost depending on phone
// usage. The following constraints:
// • A number of time < = 100, total cost is 40 Baht
// • A number of time is between 101 and 200, cost per minute is 0.50
// Baht
// • A number of time > 200, cost per minute is 0.25 Baht
// • If the phone cost is over 100, there is 10% discount.

//Input timeUsage
//Check timeUsage <= 100 | Bill = 40
//Check timeUsage between 101 - 200 | Bill = Bill + Bill * 0.5
//Check timeUsage > 200 | Bill = Bill + Bill * 0.5
//Check Bill > 100 | Bill = Bill + Bill * 0.5

func CalculateBill(usage float64) float64 {
	bill := 0.0

	if usage > 0 {
		bill = 40.0
	}

	if usage > 100 {

		if usage <= 200 {
			bill = bill + ((usage - 100.0) * 0.5)
		} else {
			bill = bill + 50 + ((usage - 200.0) * 0.25)

			if bill > 100 {
				bill = BillOver100(bill)
			}

		}
	}
	return bill
}

func BillOver100(bill float64) float64 {

	if bill > 100 {
		bill = bill * 0.9
	}
	return bill
}

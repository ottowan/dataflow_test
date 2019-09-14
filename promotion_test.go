package main

import "testing"

func Test_calculateBill_Input_0_Should_Be_0(t *testing.T) {
	inputOne := 0.00
	expectedResult := 0.00
	actualResult := CalculateBill(inputOne)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but i got %v", expectedResult, actualResult)
	}
}

// func Test_calculateBill_Input_208_Should_Be_92(t *testing.T) {
// 	inputOne := 208.00
// 	expectedResult := 92.00
// 	actualResult := CalculateBill(inputOne)

// 	if expectedResult != actualResult {
// 		t.Errorf("Expected %v but i got %v", expectedResult, actualResult)
// 	}
// }

func Test_calculateBill_Input_220_Should_Be_92(t *testing.T) {
	inputOne := 220.00
	expectedResult := 92.00
	actualResult := CalculateBill(inputOne)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but i got %v", expectedResult, actualResult)
	}
}

func Test_calculateBill_Input_250_Should_Be_92(t *testing.T) {
	inputOne := 300.00
	expectedResult := 92.00
	actualResult := CalculateBill(inputOne)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but i got %v", expectedResult, actualResult)
	}
}



func Test_BillOver100_Input_0_Should_be_0(t){
	
}

package unit_testing

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(5, 5)
	if sum != 10 {
		t.Error("There in some error in Add function")
	}
}

func TestGetGrade(t *testing.T) {
	score := getGrade(45)
	if score != "F" {
		t.Error("The score F is expected")
	}

	score = getGrade(91)
	if score != "A" {
		t.Error("The score A is expected")
	}

	score = getGrade(65)
	if score != "D" {
		t.Error("The score D is expected")
	}

	score = getGrade(82)
	if score != "B" {
		t.Error("The score B is expected")
	}

	score = getGrade(75)
	if score != "C" {
		t.Error("The score C is expected")
	}

}

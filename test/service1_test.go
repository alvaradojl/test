// Package service1_test tests all functions of service1
package service1_test

import (
	"testing"

	"github.com/alvaradojl/test/pkg/service1"
)

// Add_1And2_Return_3 tests that 1+2 equals 3
func Add_1And2_Return_3(t *testing.T) {
	//arrange
	num1 := 1
	num2 := 2
	expected := 3
	//act
	actual, err := service1.Add(num1, num2)
	//assert
	if err != nil {
		t.Error(err)
	}

	if expected != actual {
		t.Fail()
	}
}

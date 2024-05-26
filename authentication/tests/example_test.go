package tests

import (
	"testing"
	// "users-service/src/middleware"
)

// func TestMinusOne(t *testing.T) {
// 	number := 100
// 	res := middleware.MinisOne(number)

// 	if res != 99 {
// 		t.Errorf("minusOne should return 99...")
// 	}
// }

func TestHello(t *testing.T) {
	res := "hello"
	if res != "hello" {
		t.Errorf("should be hello?")
	}
}

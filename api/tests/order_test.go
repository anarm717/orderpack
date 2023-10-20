package main

import (
	"GoOrderPackProject/api/controllers"
	"fmt"
	"reflect"
	"testing"
)

func TestFulfillOrder(t *testing.T) {
	packSizes := []int{250, 500, 1000, 2000, 5000}

	testCases := []struct {
		orderQuantity int
		expected      map[int]int
	}{
		{1, map[int]int{250: 1}},
		{250, map[int]int{250: 1}},
		{251, map[int]int{500: 1}},
		{501, map[int]int{500: 1, 250: 1}},
		{12001, map[int]int{5000: 2, 2000: 1, 250: 1}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("OrderQuantity=%d", tc.orderQuantity), func(t *testing.T) {
			fulfillment := controllers.GetPacks(tc.orderQuantity, packSizes)
			if reflect.DeepEqual(fulfillment, tc.expected) {
				t.Logf("Fulfillment details match the expected result.")
			} else {
				t.Errorf("Fulfillment details do not match the expected result. Got: %v, Expected: %v", fulfillment, tc.expected)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

type testPair struct {
	input  string
	output Address
}

var tests = []testPair{
	{
		"100 South Main Street",
		Address{
			House:           100,
			StreetDirection: "s",
			StreetName:      "main",
			StreetType:      "st",
		},
	}, {
		"100 South Main Street Apartment 21",
		Address{
			House:           100,
			StreetDirection: "s",
			StreetName:      "main",
			StreetType:      "st",
		},
	}, {
		"100 South Main Street apt 212",
		Address{
			House:           100,
			StreetDirection: "s",
			StreetName:      "main",
			StreetType:      "st",
		},
	}, {
		"100 South Main Street suite 23",
		Address{
			House:           100,
			StreetDirection: "s",
			StreetName:      "main",
			StreetType:      "st",
		},
	}, {
		"100 s Main Street suite 23",
		Address{
			House:           100,
			StreetDirection: "s",
			StreetName:      "main",
			StreetType:      "st",
		},
	}, {
		"100 s Main ln suite 23",
		Address{
			House:           100,
			StreetDirection: "s",
			StreetName:      "main",
			StreetType:      "st",
		},
	},
}

func TestNormalize(t *testing.T) {
	for _, pair := range tests {
		v, err := Normalize(pair.input)
		if err != nil {
			t.Error(err)
		}
		if reflect.DeepEqual(v, pair.output) != true {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", v,
			)
		}
	}
}

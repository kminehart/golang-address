package address

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
		"123 s asdf st",
		Address{
			House:           123,
			StreetDirection: "s",
			StreetName:      "asdf",
			StreetType:      "st",
		},
	}, {
		"123 s asdf rd",
		Address{
			House:           123,
			StreetDirection: "s",
			StreetName:      "asdf",
			StreetType:      "rd",
		},
	}, {
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
			SuiteType:       "apt",
			SuiteNumber:     "21",
		},
	}, {
		"100 South Main Street apt 212",
		Address{
			House:           100,
			StreetDirection: "s",
			StreetName:      "main",
			StreetType:      "st",
			SuiteType:       "apt",
			SuiteNumber:     "212",
		},
	}, {
		"100 South Main Street suite 23",
		Address{
			House:           100,
			StreetDirection: "s",
			StreetName:      "main",
			StreetType:      "st",
			SuiteType:       "ste",
			SuiteNumber:     "23",
		},
	}, {
		"100 s Main Street suite 23",
		Address{
			House:           100,
			StreetDirection: "s",
			StreetName:      "main",
			StreetType:      "st",
			SuiteType:       "ste",
			SuiteNumber:     "23",
		},
	}, {
		"100 s Main ln apt 23",
		Address{
			House:           100,
			StreetDirection: "s",
			StreetName:      "main",
			StreetType:      "ln",
			SuiteType:       "apt",
			SuiteNumber:     "23",
		},
	}, {
		"123 Main st suite 23",
		Address{
			House:       123,
			StreetName:  "main",
			StreetType:  "st",
			SuiteType:   "ste",
			SuiteNumber: "23",
		},
	}, {
		"123 NorTHeast Main DRIVE apartment 235",
		Address{
			House:           123,
			StreetDirection: "ne",
			StreetName:      "main",
			StreetType:      "dr",
			SuiteType:       "apt",
			SuiteNumber:     "235",
		},
	}, {
		"100 Northeast Avenue",
		Address{
			House:      100,
			StreetName: "northeast",
			StreetType: "ave",
		},
	}, {
		"RR 2 Box 152",
		Address{
			StreetName: "rr 2 box 152",
		},
	}, {
		"po box 123",
		Address{
			StreetName: "po box 123",
		},
	}, {
		"110 CR 123",
		Address{
			House:      110,
			StreetName: "county road 123",
		},
	}, {
		"123 county ROAD 110",
		Address{
			House:      123,
			StreetName: "county road 110",
		},
	}, {
		"123 some invalid address 111",
		Address{
			House:      123,
			StreetName: "some invalid address 111",
		},
	}, {
		"123 south main st apt c",
		Address{
			House:           123,
			StreetDirection: "s",
			StreetName:      "main",
			StreetType:      "st",
			SuiteType:       "apt",
			SuiteNumber:     "c",
		},
	}, {
		"stupid shitty address that doesn't validate.",
		Address{
			StreetName: "stupid shitty address that doesnt validate",
		},
	}, {
		"128.5 s main street bitch",
		Address{
			StreetName: "128.5 s main street bitch",
		},
	}, {
    "100 Southwest Maple Tree Blvd",
    Address{
      House:  100,
      StreetDirection: "sw",
      StreetName: "maple tree",
      StreetType: "blvd",
    },
  }, {
    "100 S Imperial Valley Drive",
    Address{
      House:  100,
      StreetDirection: "s",
      StreetName: "imperial valley",
      StreetType: "dr",
    },
  }, {
    "100 s. weird. st.",
    Address{
      House: 100,
      StreetDirection: "s",
      StreetName: "weird",
      StreetType: "st",
    },
  }, {
    "221b s main st",
    Address{
      StreetName: "221b s main st",
    },
  }, {
    "221 completely unrecognizable street name",
    Address{
      House: 221,
      StreetName: "completely unrecognizable street name",
    },
  }, {
    "100 Circle Drive",
    Address{
      House: 100,
      StreetName: "circle",
      StreetType: "dr",
    },
  },
}

func TestNormalize(t *testing.T) {
	for _, pair := range tests {
		v, err := Normalize(pair.input)
		if err != nil {
			t.Log(err)
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

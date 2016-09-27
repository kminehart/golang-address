// golang-address parses and validates US addresses according to USPS standards.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Normalize(s string) (Address, error) {
	// This iterator is for walking across the address
	i := 0
	components := strings.Split(s, " ")
	if len(components) <= 1 {
		return Address{}, fmt.Errorf("That's not an address.")
	}
	// 0: House Number
	// 1: Street Direction
	// 2: Possibly Street Direction again, example: South West Main Street
	// However, we can not rightfully assume that if they separate the two, that they were intended to be together.
	// Example:  South East Steet.
	//   South:  Caridnal Direction,
	//   East:   Street name,
	//   Street: Strry type.
	// 2: Street Name.
	// 3: Secondary Address Unit Designator
	// 4: Secondary Address unit Designator Number

	var address Address
	var err error
	address.House, err = strconv.Atoi(components[i])

	if err != nil {
		return Address{}, fmt.Errorf("The house number %s is not valid.", components[0])
	}

	i++

	if direction, ok := CardinalDirectionAbbreviations[components[i]]; ok {
		address.StreetDirection = direction
		i++
	}

	address.StreetName = components[i]
	i++

	return address, nil
}

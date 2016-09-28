// golang-address parses and validates US addresses according to USPS standards.
package address

import (
	"fmt"
	"strconv"
	"strings"
)

func Normalize(s string) (Address, error) {
	// This iterator is for walking across the address
	var componentQueue queue

	components := strings.Split(s, " ")
	max := len(components)

	if max <= 1 {
		return Address{}, fmt.Errorf("That's not an address.")
	}

	for i, _ := range components {
		components[i] = strings.ToLower(components[i])
		componentQueue.Push(&components[i]) // First in, first out.
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
	var val *string

	val, err = componentQueue.Pop()

	if err != nil {
		return Address{}, fmt.Errorf("Error when getting house number: %s\n", err.Error())
	}
	// Convert house number to integer
	address.House, err = strconv.Atoi(*val)

	if err != nil {
		return Address{}, fmt.Errorf("The house number %s is not valid.", *val)
	}

	val = componentQueue.Peek()

	if val == nil {
		return address, fmt.Errorf("Error when getting street direction: %s\n", err.Error())
	}

	// It is maybe possible for a cardinal direction to be completely omitted.
	if direction, ok := CardinalDirectionAbbreviations[*val]; ok {
		val, err = componentQueue.Pop()
		address.StreetDirection = direction
	}

	val, err = componentQueue.Pop()

	if err != nil {
		return address, fmt.Errorf("Error when getting street name:  %s\n", err.Error())
	}

	address.StreetName = *val

	val, err = componentQueue.Pop()

	if err != nil {
		return address, fmt.Errorf("Error when getting street type: %s\n", err.Error())
	}

	if streetType, ok := StreetTypeAbbreviations[*val]; ok {
		address.StreetType = streetType
	} else {
		return address, fmt.Errorf("%s is not a valid street type.", *val)
	}

	val, err = componentQueue.Pop()

	if err != nil {
		return address, nil
	}

	if aptType, ok := SuiteTypeAbbreviations[*val]; ok {
		address.SuiteType = aptType

		val, err = componentQueue.Pop()

		if err != nil {
			return address, fmt.Errorf("An apartment/suite type was found, but the number was not.  Error:%s\n", err.Error())
		}

		address.SuiteNumber, err = strconv.Atoi(*val)

		if err != nil {
			return address, err
		}
	}
	return address, nil
}

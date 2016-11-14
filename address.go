// golang-address parses and validates US addresses according to USPS standards.
package address

import (
	"fmt"
  "regexp"
	"strconv"
	"strings"
)

func del(s *[]string, i int) {
	*s = append((*s)[:(i)], (*s)[(i)+1:]...)
}

// Returns the first number found.
//    IF it finds a suite type first
//    THEN there must not be a house/builing number.
func (a *Address) getHouseNumber(s *[]string) (int, error) {
	for i, e := range *s {
		if _, ok := SuiteTypeAbbreviations[e]; ok {
			return 0, fmt.Errorf("Found a suite type before a house number.  There must not be a house number.")
		}
		n, err := strconv.Atoi(e)
		if err != nil {
			continue
		}
		// Found a number!
		del(s, i)
		return n, nil
	}
	return 0, fmt.Errorf("Could not find a suitable house number.")
}

func (a *Address) getStreetDirection(s *[]string) (string, error) {
	for i, e := range *s {
		if n, ok := CardinalDirectionAbbreviations[e]; ok {
			del(s, i)
			return n, nil
		}
	}
	return "", fmt.Errorf("No suitable street direction was found in the address.")
}

// Returns the first street-type found.
//	IF no street-types are in the string that match the streettypeabbreviations map,
//	THEN an empty string will be returned.
func (a *Address) getStreetType(s *[]string) (string, error) {
  var t string
  var m int
	for i, e := range *s {
		if n, ok := StreetTypeAbbreviations[e]; ok {
      // Is there anything proceeding the street type?
      if len(*s) > i + 1 {
        if isAbbreviation((*s)[i+1]) == true {
          // Is it "apt", "suite", "s"...?
          // then do nothing
        } else {
          return "", nil
        }
      }
      m = i
      t = n
			// return n, nil
		}
	}
  if t == "" {
    return "", fmt.Errorf("No suitable street type was found in the address.")
  }
  del(s, m)
  return t, nil
}

func isAbbreviation(s string) bool {
  if _, ok := StreetTypeAbbreviations[s]; ok {
    return true
  }
  if _, ok := CardinalDirectionAbbreviations[s]; ok {
    return true
  }
  if _, ok := RuralBoxes[s]; ok {
    return true
  }
  if _, ok := SuiteTypeAbbreviations[s]; ok {
    return true
  }
  return false
}
// Will return the first non-numeric element in the array of strings
// that does not exist in any of the "abbreviation" maps.
//   IF there no non-numeric strings that do not exist in any of the maps,
//      it will return the first element that is non-numeric.
//   IF the (finally) found string is supposed to have a pair, like po->box or county->road,
//      it will find the matching pair and return the 2 as one string.
//      IF the final string exists in the "NumbersRequired" array, then it will grab the number immediately following the latest index.
// Problem childs:
//    100 Northwest Avenue (Cardinal direction street-names)
//    po box 123
//    100 county road 23
func (a *Address) getStreetName(s *[]string) (r string, err error) {
	// The indices to remove at the end of the function.
	var removeQueue queue

	for i, e := range *s {
		_, err := strconv.Atoi(e)
		if err == nil {
			continue
		}

    if isAbbreviation(e) == true {
      continue
    }

		removeQueue.Push(i)
		r = e
		break
	}

	if r == "" {
		for i, e := range *s {
			_, err := strconv.Atoi(e)
			if err == nil {
				continue
			}
			if _, ok := CardinalDirectionAbbreviations[e]; ok {
				r = e
				removeQueue.Push(i)
				break
			}
			if n, ok := RuralBoxes[e]; ok {
				r = n
				removeQueue.Push(i)
				break
			}
			if _, ok := SuiteTypeAbbreviations[e]; ok {
				r = e
				removeQueue.Push(i)
				break
			}
		}
	}

	if p, ok := Pairs[r]; ok {
		var n string
		// Find our match, make sure it's in here.
		for i, e := range *s {
			if e == p {
				n = p
				removeQueue.Push(i)
			}
		}

		if n == "" {
			return "", fmt.Errorf("Badly formatted address.  Found a value that expected a pair, but the pair does not exist. (Example:  \"PO\" but no \"BOX\", \"County\" but no \"Road\")")
		}

		// You would think you could just join the 2 at this point and call it done, but nope. cr has to become "County road" at some point...
		if b, ok := RuralBoxes[n]; ok {
			r = b
		} else {
			// Join the strings.  rr, 23 becomes rr 23
			st := strings.Join([]string{
				r,
				n,
			}, " ")
			r = st
		}
	}

	if r == "" {
		return "", fmt.Errorf("Could not find a proper street name.")
	}

	// Now check to see if the street name requires a number
	for _, n := range NumberRequired {
		if r != n {
			continue
		}
		// The last index of a removed element (also read:  valid element).
		l := removeQueue.PeekLast()

		// That's an L, not a 1
		_, err = strconv.Atoi((*s)[l+1])

		// It's not a number so we don't want that crap.
		if err != nil {
			return "", fmt.Errorf("This kind of address requires a number in the street name.  Example:  CR 123;  County Road 100")
			continue
		}
		st := strings.Join([]string{
			r,
			(*s)[(l + 1)],
		}, " ")
		r = st
		removeQueue.Push(l + 1)
	}

	for {
		t, err := removeQueue.Pop()
		if err != nil {
			// end of queue
			break
		}
		del(s, t)

		ra := removeQueue.Get()
		for x, y := range ra {
			if y > t {
				ra[x] = y - 1
			}
		}
	}

	return r, nil
}

func (a *Address) getSuite(s *[]string) (suiteType string, suiteNumber string, err error) {
	for i, e := range *s {
		if n, ok := SuiteTypeAbbreviations[e]; ok {
			suiteType = n
			if len(*s) > (i + 1) {
				suiteNumber = (*s)[i+1]
				del(s, i)
				del(s, i)
				return suiteType, suiteNumber, nil
			} else {
				del(s, i)
				return suiteType, "", fmt.Errorf("Suite type was found, but a suite number was not.")
			}
		}
	}
	return "", "", fmt.Errorf("No suitable suite type was found in the address.")
}

func (a *Address) finalize(s *[]string) {
	if len(*s) <= 0 {
		return
	}

	for i := 0; i < len(*s); i++ {
		st := strings.Join([]string{
			a.StreetName,
			(*s)[i],
		}, " ")
		a.StreetName = st
	}
}
func Normalize(s string) (a Address, err error) {
  if err != nil {
    return Address{}, err
  }

	t := strings.Fields(s)

	// Lowercase it all.
	for i := range t {
		t[i] = strings.ToLower(t[i])
	}

  // Strip out miscellaneous characters.
  regIsNumeric, err := regexp.Compile("^(\\d|\\.)+$")
  regMatchInvalid, err := regexp.Compile("[^a-zA-Z\\d\\s:]")

  for i := range t {
    // Match numeric fields and don't replace their decimals.
    if regIsNumeric.MatchString(t[i]) == false {
      // Remove stray periods and whatnot.
      t[i] = regMatchInvalid.ReplaceAllString(t[i], "")
    }
  }
	// Every address has a street name.  Start with that.
	a.StreetName, err = a.getStreetName(&t)
	if err != nil {
		return Address{}, err
	}

	a.House, err = a.getHouseNumber(&t)
	if err != nil {
		a.finalize(&t)
		return a, nil
	}

	a.StreetType, err = a.getStreetType(&t)

	a.SuiteType, a.SuiteNumber, err = a.getSuite(&t)

	a.StreetDirection, err = a.getStreetDirection(&t)

	a.finalize(&t)
	return a, nil
}

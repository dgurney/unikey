package validator

/*
   Copyright (C) 2020 Daniel Gurney
   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func checkdigitCheck(k int64) bool {
	// Check digit cannot be 0 or >= 8.
	if k%10 == 0 || k%10 >= 8 {
		return false
	}
	return true
}

// Mod7OEM is a mod7 OEM key
type Mod7OEM struct {
	First  string
	Second string
	Third  string
	Fourth string
	Is95   bool
}

// Mod7ElevenCD is an 11-digit mod7 CD key
type Mod7ElevenCD struct {
	First  string
	Second string
}

// Mod7CD is an 10-digit mod7 CD key
type Mod7CD struct {
	First  string
	Second string
	Is95   bool
}

// Validate validates an 11-digit mod7 CD key
func (e Mod7ElevenCD) Validate() error {
	// +1 to account for the dash
	if len(e.First)+len(e.Second)+1 != 12 {
		return errors.New("key is not the correct length")
	}

	first, err := strconv.ParseInt(e.First[0:4], 10, 0)
	if err != nil {
		return errors.New("first segment is not a number")
	}
	main, err := strconv.ParseInt(e.Second[0:7], 10, 0)
	if err != nil {
		return errors.New("second segment is not a number")
	}

	// Error is safe to discard since we checked if it's a number before.
	third, _ := strconv.ParseInt(e.First[2:3], 10, 0)
	last := first % 10
	if last != third+1 && last != third+2 {
		switch {
		case third == 8 && last != 9 && last != 0:
			return fmt.Errorf("last digit of the first segment should be 9 or 0, not %d", last)
		case third+1 >= 9 && last == 0 || third+2 >= 9 && last == 1:
			break
		default:
			return fmt.Errorf("last digit of the first segment should be %d or %d, not %d", (third+1)%10, (third+2)%10, last)
		}
	}

	if !checkdigitCheck(main) {
		return fmt.Errorf("check digit of the second segment should be > 0 and < 8, not %d", main%10)
	}
	sum := digitsum(main)
	if sum%7 != 0 {
		return fmt.Errorf("digit sum of the second segment should be divisible by 7, %d is not", sum)
	}
	return nil
}

// Validate validates a 10-digit mod7 CD key
func (c Mod7CD) Validate() error {
	// +1 to account for the dash
	if len(c.First)+len(c.Second)+1 != 11 {
		return errors.New("key is not the correct length")

	}

	site, err := strconv.ParseInt(c.First[0:3], 10, 0)
	if err != nil && !c.Is95 {
		return errors.New("first segment is not a number")
	}
	main, err := strconv.ParseInt(c.Second[0:7], 10, 0)
	if err != nil {
		return errors.New("last segment is not a number")
	}

	invalidSites := map[int64]int{333: 333, 444: 444, 555: 555, 666: 666, 777: 777, 888: 888, 999: 999}
	_, invalid := invalidSites[site]
	if invalid {
		return errors.New("site number should not be 333, 444, 555, 666, 777, 888, or 999")
	}
	if !checkdigitCheck(main) && !c.Is95 {
		return fmt.Errorf("check digit of the second segment should be > 0 and < 8, not %d", main%10)
	}
	sum := digitsum(main)
	if sum%7 != 0 {
		return fmt.Errorf("digit sum of the second segment should be divisible by 7, %d is not", sum)
	}
	return nil
}

// Validate validates a mod7 OEM key
func (o Mod7OEM) Validate() error {
	// +3 to account for dashes
	if len(o.First)+len(o.Second)+len(o.Third)+len(o.Fourth)+3 != 23 {
		return errors.New("key is not the correct length")
	}

	_, err := strconv.ParseInt(o.First[0:5], 10, 0)
	if err != nil {
		return errors.New("first segment is not a number")
	}
	th, err := strconv.ParseInt(o.Third[0:7], 10, 0)
	if err != nil {
		return errors.New("third segment is not a number")
	}
	_, err = strconv.ParseInt(o.Fourth[0:], 10, 0)
	if err != nil {
		return errors.New("fourth segment is not a number")
	}
	julian, err := strconv.ParseInt(o.First[0:3], 10, 0)
	if julian == 0 || julian > 366 {
		return fmt.Errorf("date should be within 001-366, not %03d", julian)
	}

	year := o.First[3:5]
	validYears := map[string]string{"95": "95", "96": "96", "97": "97", "98": "98", "99": "99", "00": "00", "01": "01", "02": "02", "03": "03"}
	if o.Is95 {
		validYears = map[string]string{"95": "95", "96": "96", "97": "97", "98": "98", "99": "99", "00": "00", "01": "01", "02": "02"}
	}
	_, valid := validYears[year]
	if !valid {
		switch {
		default:
			return fmt.Errorf("year should be within 95-03, not %s", year)
		case o.Is95:
			return fmt.Errorf("year should be within 95-02 for Windows 95, not %s", year)
		}
	}

	if strings.ToUpper(o.Second) != "OEM" {
		return fmt.Errorf("second segment should be OEM, not %s", o.Second)
	}

	third := o.Third[0:7]
	if string(third[0]) != "0" {
		return fmt.Errorf("third segment beginning should be 0, not %s", third[0:1])
	}
	if !checkdigitCheck(th) {
		return fmt.Errorf("check digit of the third segment should be > 0 and < 8, not %d", th%10)
	}
	sum := digitsum(th)
	if sum%7 != 0 {
		return fmt.Errorf("digit sum of the third segment should be divisible by 7, %d is not", sum)
	}

	return nil
}

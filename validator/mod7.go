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
	"strings"
)

func validCheckDigit(k int64) bool {
	checkDigit := k % 10
	return checkDigit > 0 && checkDigit < 8
}

// Mod7OEM is a mod7 OEM key.
type Mod7OEM struct {
	First  string
	Second string
	Third  string
	Fourth string
	Is95   bool // Windows 95 accepts years only through 02.
}

// Mod7ElevenCD is an 11-digit mod7 CD key.
type Mod7ElevenCD struct {
	First                string
	Second               string
	EnableCheckDigitRule bool // Some products, including Office 97, do not enforce this rule.
}

// Mod7CD is a 10-digit mod7 CD key.
type Mod7CD struct {
	First  string
	Second string
	Is95   bool // Windows 95 permits nonnumeric site IDs and any check digit.
}

// Validate validates an 11-digit mod7 CD key.
func (e Mod7ElevenCD) Validate() error {
	if len(e.First) != 4 || len(e.Second) != 7 {
		return errors.New("key is not the correct length")
	}

	first, ok := parseDecimal(e.First)
	if !ok {
		return errors.New("first segment is not a number")
	}
	main, ok := parseDecimal(e.Second)
	if !ok {
		return errors.New("second segment is not a number")
	}

	third := int64(e.First[2] - '0')
	last := first % 10
	expectedFirst := (third + 1) % 10
	expectedSecond := (third + 2) % 10
	if last != expectedFirst && last != expectedSecond {
		return fmt.Errorf("last digit of the first segment should be %d or %d, not %d", expectedFirst, expectedSecond, last)
	}

	if e.EnableCheckDigitRule && !validCheckDigit(main) {
		return fmt.Errorf("check digit of the second segment should be > 0 and < 8, not %d", main%10)
	}
	sum := digitSum(main)
	if sum%7 != 0 {
		return fmt.Errorf("digit sum of the second segment should be divisible by 7, %d is not", sum)
	}
	return nil
}

// Validate validates a 10-digit mod7 CD key.
func (c Mod7CD) Validate() error {
	if len(c.First) != 3 || len(c.Second) != 7 {
		return errors.New("key is not the correct length")
	}

	if !c.Is95 {
		if _, ok := parseDecimal(c.First); !ok {
			return errors.New("first segment is not a number")
		}
	}
	main, ok := parseDecimal(c.Second)
	if !ok {
		return errors.New("last segment is not a number")
	}

	switch c.First {
	case "333", "444", "555", "666", "777", "888", "999":
		return errors.New("site number should not be 333, 444, 555, 666, 777, 888, or 999")
	}
	if !validCheckDigit(main) && !c.Is95 {
		return fmt.Errorf("check digit of the second segment should be > 0 and < 8, not %d", main%10)
	}
	sum := digitSum(main)
	if sum%7 != 0 {
		return fmt.Errorf("digit sum of the second segment should be divisible by 7, %d is not", sum)
	}
	return nil
}

// Validate validates a mod7 OEM key.
func (o Mod7OEM) Validate() error {
	if len(o.First) != 5 || len(o.Second) != 3 || len(o.Third) != 7 || len(o.Fourth) != 5 {
		return errors.New("key is not the correct length")
	}

	if _, ok := parseDecimal(o.First); !ok {
		return errors.New("first segment is not a number")
	}
	third, ok := parseDecimal(o.Third)
	if !ok {
		return errors.New("third segment is not a number")
	}
	if _, ok := parseDecimal(o.Fourth); !ok {
		return errors.New("fourth segment is not a number")
	}

	julian := int64(o.First[0]-'0')*100 + int64(o.First[1]-'0')*10 + int64(o.First[2]-'0')
	if julian == 0 || julian > 366 {
		return fmt.Errorf("date should be within 001-366, not %03d", julian)
	}

	year := o.First[3:5]
	switch year {
	case "95", "96", "97", "98", "99", "00", "01", "02":
	case "03":
		if o.Is95 {
			return fmt.Errorf("year should be within 95-02 for Windows 95, not %s", year)
		}
	default:
		return fmt.Errorf("year should be within 95-03, not %s", year)
	}

	if strings.ToUpper(o.Second) != "OEM" {
		return fmt.Errorf("second segment should be OEM, not %s", o.Second)
	}

	if o.Third[0] != '0' {
		return fmt.Errorf("third segment beginning should be 0, not %s", o.Third[0:1])
	}
	if !validCheckDigit(third) {
		return fmt.Errorf("check digit of the third segment should be > 0 and < 8, not %d", third%10)
	}
	sum := digitSum(third)
	if sum%7 != 0 {
		return fmt.Errorf("digit sum of the third segment should be divisible by 7, %d is not", sum)
	}

	return nil
}

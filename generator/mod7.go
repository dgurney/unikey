package generator

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
	"fmt"
	"math/rand"
	"strconv"
)

// Mod7OEM is a mod7 OEM key
type Mod7OEM struct {
	First  string
	Second string
	Third  int
	Fourth int
}

// Mod7ElevenCD is an 11-digit mod7 CD key
type Mod7ElevenCD struct {
	First  string
	Second int
}

// Mod7CD is an 10-digit mod7 CD key
type Mod7CD struct {
	First  int
	Second int
}

func checkdigitCheck(k int) bool {
	// Check digit cannot be 0 or >= 8.
	if k%10 == 0 || k%10 >= 8 {
		return false
	}
	return true
}

func (c Mod7ElevenCD) String() string {
	return fmt.Sprintf("%s-%07d", c.First, c.Second)
}

// Generate generates an 11-digit mod7 CD key.
func (c Mod7ElevenCD) Generate() (KeyGenerator, error) {
	// Generate the first segment of the key.
	// Formula for last digit: third digit + 1 or 2. If the result is more than 9, it's 0 or 1.
	s := rand.Intn(999)
	last := s % 10
	first := ""
	fourth := 0
	switch {
	default:
		fourth = last + 1
	case rand.Intn(2) == 1:
		fourth = last + 2
	}
	switch {
	case fourth == 10:
		fourth = 0
	case fourth > 10:
		fourth = 1
	}
	first = fmt.Sprintf("%03d%d", s, fourth)

	// Generate the second segment of the key. The digit sum of the seven numbers must be divisible by seven.
	// The last digit is the check digit. The check digit cannot be 0 or >=8.
	second := 0
	for {
		s := rand.Intn(9999999)
		// Perform the actual validation
		sum := digitsum(s)
		if sum%7 == 0 {
			second = s
			if checkdigitCheck(s) {
				break
			}
		}
	}
	c.First = first
	c.Second = second
	return c, nil
}

func (c Mod7CD) String() string {
	return fmt.Sprintf("%03d-%07d", c.First, c.Second)
}

// Generate generates a 10-digit mod7 CD key.
func (c Mod7CD) Generate() (KeyGenerator, error) {
	// Generate the so-called site number, which is the first segment of the key.
	first := rand.Intn(998)
	// Technically 999 could be omitted as we don't generate a number that high, but we include it for posterity anyway.
	invalidSites := []int{333, 444, 555, 666, 777, 888, 999}
	for _, v := range invalidSites {
		if v == first {
			// Site number is invalid, so we replace it with a guaranteed valid number
			first = rand.Intn(300)
		}
	}

	// Generate the second segment of the key. The digit sum of the seven numbers must be divisible by seven.
	// The last digit is the check digit. The check digit cannot be 0 or >=8.
	// Note that Windows 95 does not have a check digit check.
	second := 0
	for {
		second = rand.Intn(9999999)
		// Perform the actual validation
		sum := digitsum(second)
		if sum%7 == 0 {
			if checkdigitCheck(second) {
				break
			}
		}
	}
	c.First = first
	c.Second = second
	return c, nil
}

func (o Mod7OEM) String() string {
	return fmt.Sprintf("%s-%s-0%06d-%05d", o.First, o.Second, o.Third, o.Fourth)
}

// Generate generates a mod7 OEM key
func (o Mod7OEM) Generate() (KeyGenerator, error) {
	// Generate the first segment of the key. The first three digits represent the julian date the COA was printed (001 to 366), and the last two are the year.
	// The year cannot be below 95 or above 03 (not Y2K-compliant D:).
	// The maximum year for Windows 95 is 02.
	d := 0
	first := ""
	nonzero := false
	for !nonzero {
		switch {
		case d != 0:
			nonzero = true
		default:
			d = rand.Intn(366)
		}
	}
	date := fmt.Sprintf("%03d", d)
	// 03 is also valid for many later products, but for Windows 95 it is not
	years := []string{"95", "96", "97", "98", "99", "00", "01", "02"}
	year := years[rand.Intn(len(years))]

	// Check that year is actually a leap year and adjust day 366 to 365 accordingly if not
	y2kify := func(y string) int {
		yearInt, _ := strconv.Atoi(y)
		switch {
		case yearInt >= 95:
			return 1900 + yearInt
		default:
			return 2000 + yearInt
		}
	}(year)
	if !isLeap(y2kify) && date == "366" {
		date = "365"
	}

	first = date + year

	// The third segment (OEM is the second) must begin with a zero, but otherwise it follows the same rule as the second segment of 10-digit keys:
	// The digit sum must be divisible by seven, and the check digit cannot be 0 or >=8.
	third := 0
	for {
		s := rand.Intn(999999)
		// Perform the actual validation
		sum := digitsum(s)
		if sum%7 == 0 {
			third = s
			if checkdigitCheck(s) {
				break
			}
		}
	}

	// The fourth segment is truly irrelevant
	fourth := rand.Intn(99999)
	o.First = first
	o.Second = "OEM"
	o.Third = third
	o.Fourth = fourth
	return o, nil
}

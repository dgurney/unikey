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
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Mod7OEM is a mod7 OEM key
type Mod7OEM struct {
}

// Mod7ElevenCD is an 11-digit mod7 CD key
type Mod7ElevenCD struct {
}

// Mod7CD is an 10-digit mod7 CD key
type Mod7CD struct {
}

func checkdigitCheck(k int) bool {
	// Check digit cannot be 0 or >= 8.
	if k%10 == 0 || k%10 >= 8 {
		return false
	}
	return true
}

// Generate generates an 11-digit mod7 CD key.
func (Mod7ElevenCD) Generate(ch chan string) {
	// Generate the first segment of the key.
	// Formula for last digit: third digit + 1 or 2. If the result is more than 9, it's 0 or 1.
	s := rand.Intn(999)
	site := fmt.Sprintf("%03d", s)
	last, _ := strconv.Atoi(site[len(site)-1:])
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
	first = fmt.Sprintf("%s%d", site, fourth)

	// Generate the second segment of the key. The digit sum of the seven numbers must be divisible by seven.
	// The last digit is the check digit. The check digit cannot be 0 or >=8.
	second := ""
	for {
		s := rand.Intn(9999999)
		// Perform the actual validation
		sum := digitsum(s)
		if sum%7 == 0 {
			second = fmt.Sprintf("%07d", s)
			if checkdigitCheck(s) {
				break
			}
		}
	}
	ch <- first + "-" + second
}

// Generate generates a 10-digit mod7 CD key.
func (Mod7CD) Generate(ch chan string) {
	// Generate the so-called site number, which is the first segment of the key.
	first := ""
	s := rand.Intn(998)
	// Technically 999 could be omitted as we don't generate a number that high, but we include it for posterity anyway.
	invalidSites := []int{333, 444, 555, 666, 777, 888, 999}
	for _, v := range invalidSites {
		if v == s {
			// Site number is invalid, so we replace it with a guaranteed valid number
			s = rand.Intn(300)
		}
	}
	first = fmt.Sprintf("%03d", s)

	// Generate the second segment of the key. The digit sum of the seven numbers must be divisible by seven.
	// The last digit is the check digit. The check digit cannot be 0 or >=8.
	second := ""
	for {
		s := rand.Intn(9999999)
		// Perform the actual validation
		sum := digitsum(s)
		if sum%7 == 0 {
			second = fmt.Sprintf("%07d", s)
			if checkdigitCheck(s) {
				break
			}
		}
	}
	ch <- first + "-" + second
}

// Generate generates a mod7 OEM key
func (o Mod7OEM) Generate(ch chan string) {
	// Generate the first segment of the key. The first three digits represent the julian date the COA was printed (001 to 366), and the last two are the year.
	// The year cannot be below 95 or above 03 (not Y2K-compliant D:).
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
	years := []string{"95", "96", "97", "98", "99", "00", "01", "02", "03"}
	year := years[rand.Intn(len(years))]
	first = date + year

	// The third segment (Mod7OEM is the second) must begin with a zero, but otherwise it follows the same rule as the second segment of 10-digit keys:
	// The digit sum must be divisible by seven, and the check digit cannot be 0 or >=8.
	third := ""
	for {
		s := rand.Intn(999999)
		// Perform the actual validation
		sum := digitsum(s)
		if sum%7 == 0 {
			third = fmt.Sprintf("%06d", s)
			if checkdigitCheck(s) {
				break
			}
		}
	}

	// The fourth segment is truly irrelevant
	f := rand.Intn(99999)
	fourth := fmt.Sprintf("%05d", f)
	ch <- first + "-OEM-0" + third + "-" + fourth
}

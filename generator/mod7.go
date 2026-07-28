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
	"math/rand/v2"
)

// Mod7OEM is a mod7 OEM key.
type Mod7OEM struct {
	First  string
	Second string
	Third  int // String formats this segment with its required leading zero.
	Fourth int
}

// Mod7ElevenCD is an 11-digit mod7 CD key.
type Mod7ElevenCD struct {
	First  int
	Second int
}

// Mod7CD is a 10-digit mod7 CD key.
type Mod7CD struct {
	First  int
	Second int
}

func validCheckDigit(k int) bool {
	checkDigit := k % 10
	return checkDigit > 0 && checkDigit < 8
}

func (c Mod7ElevenCD) String() string {
	return fmt.Sprintf("%04d-%07d", c.First, c.Second)
}

// Generate generates an 11-digit mod7 CD key.
func (c *Mod7ElevenCD) Generate() {
	s := rand.IntN(1_000)
	last := s % 10
	fourth := last + rand.IntN(2) + 1
	first := s*10 + fourth%10

	second := 0
	for {
		second = rand.IntN(10_000_000)
		sum := digitSum(second)
		if sum%7 == 0 && validCheckDigit(second) {
			break
		}
	}
	c.First = first
	c.Second = second
}

func (c Mod7CD) String() string {
	return fmt.Sprintf("%03d-%07d", c.First, c.Second)
}

// Generate generates a 10-digit mod7 CD key.
func (c *Mod7CD) Generate() {
	first := rand.IntN(999)
	switch first {
	case 333, 444, 555, 666, 777, 888:
		first = rand.IntN(300)
	}

	second := 0
	for {
		second = rand.IntN(10_000_000)
		sum := digitSum(second)
		if sum%7 == 0 && validCheckDigit(second) {
			break
		}
	}
	c.First = first
	c.Second = second
}

func (o Mod7OEM) String() string {
	return fmt.Sprintf("%s-%s-0%06d-%05d", o.First, o.Second, o.Third, o.Fourth)
}

// Generate generates a mod7 OEM key.
func (o *Mod7OEM) Generate() {
	d := rand.IntN(366) + 1
	date := fmt.Sprintf("%03d", d)
	years := []string{"95", "96", "97", "98", "99", "00", "01", "02"}
	year := years[rand.IntN(len(years))]

	yearNumber := int(year[0]-'0')*10 + int(year[1]-'0')
	fullYear := 2000 + yearNumber
	if yearNumber >= 95 {
		fullYear = 1900 + yearNumber
	}
	if !isLeap(fullYear) && date == "366" {
		date = "365"
	}

	third := 0
	for {
		third = rand.IntN(1_000_000)
		sum := digitSum(third)
		if sum%7 == 0 && validCheckDigit(third) {
			break
		}
	}

	o.First = date + year
	o.Second = "OEM"
	o.Third = third
	o.Fourth = rand.IntN(100_000)
}

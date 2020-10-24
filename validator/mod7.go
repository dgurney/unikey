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

import "strconv"

// Mod7OEM is a mod7 OEM key
type Mod7OEM struct {
	Key string
}

// Mod7ElevenCD is an 11-digit mod7 CD key
type Mod7ElevenCD struct {
	Key string
}

// Mod7CD is an 10-digit mod7 CD key
type Mod7CD struct {
	Key string
}

// Validate validates an 11-digit mod7 CD key
func (e Mod7ElevenCD) Validate(v chan bool) {
	if len(e.Key) != 12 {
		v <- false
		return
	}

	_, err := strconv.ParseInt(e.Key[0:4], 10, 0)
	if err != nil {
		v <- false
		return
	}
	main, err := strconv.ParseInt(e.Key[5:12], 10, 0)
	if err != nil {
		v <- false
		return
	}

	// Error is safe to discard since we checked if it's a number before.
	last, _ := strconv.ParseInt(e.Key[3:4], 10, 0)
	third, _ := strconv.ParseInt(e.Key[2:3], 10, 0)
	if last != third+1 && last != third+2 {
		switch {
		case third == 8 && last != 9 && last != 0:
			v <- false
			return
		case third+1 >= 9 && last == 0 || third+2 >= 9 && last == 1:
			break
		default:
			v <- false
			return
		}
	}

	if !checkdigitCheck(main) {
		v <- false
		return
	}
	sum := digitsum(main)
	if sum%7 != 0 {
		v <- false
		return
	}
	v <- true
}

// Validate validates a 10-digit mod7 CD key
func (c Mod7CD) Validate(v chan bool) {
	if len(c.Key) != 11 {
		v <- false
		return
	}

	site, err := strconv.ParseInt(c.Key[0:3], 10, 0)
	if err != nil {
		v <- false
		return
	}
	main, err := strconv.ParseInt(c.Key[4:11], 10, 0)
	if err != nil {
		v <- false
		return
	}

	invalidSites := map[int64]int{333: 333, 444: 444, 555: 555, 666: 666, 777: 777, 888: 888, 999: 999}
	_, invalid := invalidSites[site]
	if invalid {
		v <- false
		return
	}
	if !checkdigitCheck(main) {
		v <- false
		return
	}
	sum := digitsum(main)
	if sum%7 != 0 {
		v <- false
		return
	}
	v <- true
}

// Validate validates a mod7 OEM key
func (o Mod7OEM) Validate(v chan bool) {
	if len(o.Key) != 23 {
		v <- false
		return
	}

	_, err := strconv.ParseInt(o.Key[0:5], 10, 0)
	if err != nil {
		v <- false
		return
	}
	th, err := strconv.ParseInt(o.Key[10:17], 10, 0)
	if err != nil {
		v <- false
		return
	}
	_, err = strconv.ParseInt(o.Key[18:], 10, 0)
	if err != nil {
		v <- false
		return
	}
	julian, err := strconv.ParseInt(o.Key[0:3], 10, 0)
	if julian == 0 || julian > 366 {
		v <- false
		return
	}

	year := o.Key[3:5]
	validYears := map[string]string{"95": "95", "96": "96", "97": "97", "98": "98", "99": "99", "00": "00", "01": "01", "02": "02", "03": "03"}
	_, valid := validYears[year]
	if !valid {
		v <- false
		return
	}

	// nice
	if o.Key[6:9] != "OEM" {
		v <- false
		return
	}

	third := o.Key[10:17]
	if string(third[0]) != "0" {
		v <- false
		return
	}
	if !checkdigitCheck(th) {
		v <- false
		return
	}
	sum := digitsum(th)
	if sum%7 != 0 {
		v <- false
		return
	}

	v <- true
}

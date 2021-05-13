package generator

import "testing"

var leapYears = []int{1904, 1908, 1912, 1916, 1920, 1924, 1928, 1932, 1936, 1940, 1944, 1948, 1952, 1956, 1960, 1964, 1968, 1972, 1976, 1980, 1984, 1988, 1992, 1996, 2000, 2004, 2008, 2012, 2016, 2020}
var notLeap = []int{1911, 1945, 1995, 2021}

func TestLeap(t *testing.T) {
	for _, y := range leapYears {
		leap := isLeap(y)
		if !leap {
			t.Fatalf("%d is unexpectedly not a leap year", y)
		}
	}
	for _, y := range notLeap {
		leap := isLeap(y)
		if leap {
			t.Fatalf("%d is unexpectedly a leap year", y)
		}
	}
}

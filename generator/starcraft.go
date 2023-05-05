package generator

import (
	"fmt"
	"math/rand"
	"strconv"
)

type StarCraft struct {
	Key string // 15-character key (including separators), of which the last digit is a check digit
}

// String returns the key with separators in place
func (s StarCraft) String() string {
	return fmt.Sprintf("%s-%s-%s", s.Key[0:4], s.Key[4:9], s.Key[9:13])
}

// Generate generates a StarCraft key without separators
func (s *StarCraft) Generate() error {
	key := fmt.Sprintf("%012d", rand.Int63n(999999999999))
	s.Key = key + strconv.Itoa(s.generateCheckDigit(key))
	return nil
}

func (s StarCraft) generateCheckDigit(key string) int {
	temp := 3
	for i := 0; i < 12; i++ {
		c, err := strconv.ParseInt(key[i:i+1], 10, 0)
		if err != nil {
			return 0
		}
		temp += (2 * temp) ^ int(c)
	}
	return temp % 10
}

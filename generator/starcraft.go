package generator

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

// StarCraft is a generated StarCraft key.
type StarCraft struct {
	key string
}

// String returns the generated key with separators.
func (s StarCraft) String() string {
	if s.key == "" {
		return ""
	}
	return fmt.Sprintf("%s-%s-%s", s.key[0:4], s.key[4:9], s.key[9:13])
}

// Generate creates a StarCraft key.
func (s *StarCraft) Generate() {
	key := fmt.Sprintf("%012d", rand.Int64N(1_000_000_000_000))
	s.key = key + strconv.Itoa(generateStarCraftCheckDigit(key))
}

func generateStarCraftCheckDigit(key string) int {
	temp := 3
	for i := range 12 {
		temp += (2 * temp) ^ int(key[i]-'0')
	}
	return temp % 10
}

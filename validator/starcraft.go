package validator

import (
	"errors"
	"fmt"
	"strconv"
)

type StarCraft struct {
	Key string // 15-character key (including dashes for segmentation), of which the last digit is a check digit
}

// Validate validates a StarCraft CD Key (key can be provided with and without separators)
// Separator character is not checked due to installer enforcing the segmentation
func (s StarCraft) Validate() error {
	if len(s.Key) != 15 && len(s.Key) != 13 {
		return errors.New("key is not in the correct format (should be in the XXXX-XXXXX-XXXX format or 13 digits)")
	}

	bareKey := ""
	switch {
	case len(s.Key) == 13:
		bareKey = s.Key
	default:
		bareKey = s.Key[0:4] + s.Key[5:10] + s.Key[11:15]
	}

	_, err := strconv.ParseInt(bareKey, 10, 64)
	if err != nil {
		return errors.New("key contains non-numeric characters")
	}

	originalCheckDigit, _ := strconv.Atoi(s.Key[len(s.Key)-1:])
	computedCheckDigit := s.generateCheckDigit(bareKey)
	if originalCheckDigit != computedCheckDigit {
		return fmt.Errorf("check digit %d does not match expected %d", originalCheckDigit, computedCheckDigit)
	}
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

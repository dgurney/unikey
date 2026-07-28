package validator

import (
	"errors"
	"fmt"
)

// StarCraft is a segmented or unsegmented StarCraft key.
type StarCraft struct {
	Key string
}

// Validate validates a StarCraft key with or without separators.
func (s StarCraft) Validate() error {
	if len(s.Key) != 15 && len(s.Key) != 13 {
		return errors.New("key is not in the correct format (should be in the XXXX-XXXXX-XXXX format or 13 digits)")
	}

	var bareKey string
	if len(s.Key) == 13 {
		bareKey = s.Key
	} else {
		// Original installers enforce the separator positions independently.
		bareKey = s.Key[0:4] + s.Key[5:10] + s.Key[11:15]
	}

	for i := range bareKey {
		if bareKey[i] < '0' || bareKey[i] > '9' {
			return errors.New("key contains non-numeric characters")
		}
	}

	originalCheckDigit := int(bareKey[len(bareKey)-1] - '0')
	computedCheckDigit := generateStarCraftCheckDigit(bareKey)
	if originalCheckDigit != computedCheckDigit {
		return fmt.Errorf("check digit %d does not match expected %d", originalCheckDigit, computedCheckDigit)
	}
	return nil
}

func generateStarCraftCheckDigit(key string) int {
	temp := 3
	for i := range 12 {
		temp += (2 * temp) ^ int(key[i]-'0')
	}
	return temp % 10
}

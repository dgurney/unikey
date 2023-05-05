package validator

import "testing"

var valid = []KeyValidator{
	// beginning of mod7
	Mod7CD{"111", "1111111", false},
	Mod7CD{"000", "0000007", false},
	Mod7CD{"118", "5688143", false},
	Mod7OEM{"10000", "OEM", "0000007", "00000", false},
	Mod7OEM{"32299", "OEM", "0840621", "16752", false},
	Mod7ElevenCD{"1112", "0000007"},
	Mod7ElevenCD{"9991", "1111111"},
	Mod7ElevenCD{"9990", "1111111"},
	Mod7ElevenCD{"8889", "1111111"},
	Mod7ElevenCD{"8880", "1111111"},

	// Windows 95 does not have a check digit check
	Mod7CD{"111", "1111109", true},
	// Windows 95 does actually allow typing non-integer site numbers
	Mod7CD{"AAA", "1111109", true},
	// end of mod7

	// beginning of starcraft
	StarCraft{"8206-64645-5171"},
	StarCraft{"1234-56789-1234"},
	StarCraft{"2374-81224-9423"},
	// separator not enforced
	StarCraft{"2374A81224B9423"},
	// key can be provided without separators
	StarCraft{"1234567891234"},
	// end of starcraft
}

var invalid = []KeyValidator{
	// beginning of mod7
	Mod7CD{"1", "1", false},
	Mod7CD{"11a", "1111111", false},
	Mod7CD{"111", "a111111", false},
	// Invalid site
	Mod7CD{"333", "5688143", false},
	Mod7CD{"444", "5688143", false},
	Mod7CD{"555", "5688143", false},
	Mod7CD{"666", "5688143", false},
	Mod7CD{"777", "5688143", false},
	Mod7CD{"888", "5688143", false},
	Mod7CD{"999", "5688143", false},
	// Invalid check digit
	Mod7CD{"332", "5683148", false},
	// Invalid main segment
	Mod7CD{"332", "5688313", false},
	Mod7ElevenCD{"1", "1"},
	Mod7ElevenCD{"111a", "1111111"},
	Mod7ElevenCD{"1111", "a111111"},
	// Invalid first segment
	Mod7ElevenCD{"1114", "1111111"},
	Mod7ElevenCD{"1117", "1111111"},
	Mod7ElevenCD{"8881", "1111111"},
	Mod7ElevenCD{"8885", "1111111"},
	Mod7ElevenCD{"9992", "1111111"},
	// Invalid digit sum
	Mod7ElevenCD{"0001", "5688144"},
	// Invalid check digit
	Mod7ElevenCD{"1112", "1111118"},
	Mod7OEM{"1", "1", "1", "1", false},
	Mod7OEM{"1000a", "OEM", "0000007", "11111", false},
	Mod7OEM{"10000", "OEM", "000000a", "11111", false},
	Mod7OEM{"10000", "OEM", "0000007", "1111a", false},
	// Invalid date
	Mod7OEM{"00099", "OEM", "0840621", "16752", false},
	Mod7OEM{"36799", "OEM", "0840621", "16752", false},
	// Invalid year
	Mod7OEM{"10094", "OEM", "0840621", "16752", false},
	Mod7OEM{"10019", "OEM", "0840621", "16752", false},
	// Invalid check digit
	Mod7OEM{"10000", "OEM", "0140628", "12345", false},
	// Invalid third segment starting digit
	Mod7OEM{"10000", "OEM", "8040621", "12345", false},
	// Invalid digit sum
	Mod7OEM{"10000", "OEM", "0000006", "12345", false},
	// Invalid second segment
	Mod7OEM{"10000", "SEX", "0000007", "12345", false},

	// Windows 95 does not allow year 03
	Mod7OEM{"10003", "OEM", "0000007", "12345", true},
	// end of mod7

	// beginning of starcraft
	// wrong check digits
	StarCraft{"8206-64645-5172"},
	StarCraft{"1234-56789-1230"},

	// typo
	StarCraft{"8260-64645-5172"},

	// letter in key
	StarCraft{"a206-64645-5172"},

	// short
	StarCraft{"8206"},
	// end of starcraft
}

func TestValidation(t *testing.T) {
	for _, v := range valid {
		err := v.Validate()
		if err != nil {
			t.Errorf("valid key %v did not pass validation! (%s)", v, err)
		}
	}
	for _, i := range invalid {
		err := i.Validate()
		if err == nil {
			t.Errorf("invalid key %v passed validation!", i)
		}
	}
}

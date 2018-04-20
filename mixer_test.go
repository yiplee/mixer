package mixer

import (
	"testing"
)

func TestBase62(t *testing.T) {
	testData := map[uint]string{
		0:       "a",
		1:       "b",
		26:      "A",
		52:      "0",
		62:      "ba",
		1000000: "emjc",
	}

	mixer := Base62Mixer()
	for num, text := range testData {
		if result := mixer.Encoding(num); result != text {
			t.Errorf("encoding %d should be %s but get %s", num, text, result)
		}

		if result, err := mixer.Decoding(text); err != nil || result != num {
			if err != nil {
				t.Error(err)
			} else {
				t.Errorf("decoding %s should be %d but get %d", text, num, result)
			}
		}
	}
}

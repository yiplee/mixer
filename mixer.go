package mixer

import (
	"errors"
)

const (
	base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

type Mixer struct {
	chars   []byte
	base    uint
	charMap map[byte]uint
}

func (self *Mixer) Encoding(num uint) string {
	if self == nil {
		return ""
	}

	bytes := make([]byte, 0, 6)
	for num >= 0 {
		reminder := num % self.base
		bytes = append(bytes, self.chars[reminder])

		if num >= self.base {
			num = num / self.base
		} else {
			break
		}
	}

	for l, r := 0, len(bytes)-1; l < r; l, r = l+1, r-1 {
		bytes[l], bytes[r] = bytes[r], bytes[l]
	}

	return string(bytes)
}

func (self *Mixer) Decoding(text string) (result uint, err error) {
	if self != nil {
		for _, char := range text {
			if num, found := self.charMap[byte(char-0)]; found {
				result = result*self.base + num
			} else {
				err = errors.New("invalid character")
				return
			}
		}
	}

	return
}

func New(chars string) (*Mixer, error) {
	if length := len(chars); length >= 2 {
		bytes := []byte(chars)
		charMap := map[byte]uint{}
		for idx, char := range bytes {
			charMap[char] = uint(idx)
		}

		if len(charMap) != length {
			return nil, errors.New("chars is invalid")
		}

		return &Mixer{
			chars:   bytes,
			charMap: charMap,
			base:    uint(length),
		}, nil
	} else {
		return nil, errors.New("chars is too short,2 at least")
	}
}

func Base62Mixer() Mixer {
	if mixer, err := New(base62Chars); err == nil {
		return *mixer
	} else {
		panic(err)
	}
}

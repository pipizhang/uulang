package uulang

import (
	"github.com/liuzl/gocc"
)

// S2T Simplified Chinese to Traditional Chinese
func S2T(in string) (string, error) {
	s2t, err := gocc.New("s2t")
	if err != nil {
		return "", err
	}
	out, err := s2t.Convert(in)
	if err != nil {
		return "", err
	}
	return out, nil
}

// T2S Traditional Chinese to Simplified Chinese
func T2S(in string) (string, error) {
	s2t, err := gocc.New("t2s")
	if err != nil {
		return "", err
	}
	out, err := s2t.Convert(in)
	if err != nil {
		return "", err
	}
	return out, nil
}

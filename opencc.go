package uulang

import (
	"github.com/liuzl/gocc"
)

var (
	s2t *gocc.OpenCC
	t2s *gocc.OpenCC
)

func init() {
	var err error
	s2t, err = gocc.New("s2t")
	if err != nil {
		panic("Failed to load OpenCC(s2t)")
	}
	t2s, err = gocc.New("t2s")
	if err != nil {
		panic("Failed to load OpenCC(t2s)")
	}
}

// S2T Simplified Chinese to Traditional Chinese
func S2T(in string) (string, error) {
	out, err := s2t.Convert(in)
	if err != nil {
		return "", err
	}
	return out, nil
}

// T2S Traditional Chinese to Simplified Chinese
func T2S(in string) (string, error) {
	out, err := t2s.Convert(in)
	if err != nil {
		return "", err
	}
	return out, nil
}

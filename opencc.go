package uulang

import (
	"fmt"
	"os"

	"github.com/liuzl/gocc"
)

var (
	s2t *gocc.OpenCC
	t2s *gocc.OpenCC
)

func init() {
	var err error

	config_path := os.Getenv("GOCC_CONFIG_PATH")
	if len(config_path) > 0 {
		gocc.Dir = &config_path
	}

	s2t, err = gocc.New("s2t")
	if err != nil {
		panic(fmt.Sprintf("Failed to load OpenCC(s2t), %s", err))
	}
	t2s, err = gocc.New("t2s")
	if err != nil {
		panic(fmt.Sprintf("Failed to load OpenCC(t2s), %s", err))
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

// IsSimplified guess word whether is a simpliefed word
func IsSimplified(in string) bool {
	t, _ := S2T(in)
	if in != t {
		return true
	}

	s, _ := T2S(t)
	if t == s {
		return true
	}

	return false
}

// IsTraditional guess word whether is a traditional word
func IsTraditional(in string) bool {
	s, _ := T2S(in)
	if in != s {
		return true
	}

	t, _ := S2T(s)
	if t == s {
		return true
	}

	return false
}

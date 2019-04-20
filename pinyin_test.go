package uulang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPinyinTone(t *testing.T) {
	py := PinyinTone("中国")
	ep := pySlice{[]string{"zhōng", "zhòng"}, []string{"guó"}}

	assert.Equal(t, ep, py)
}

func TestPinyinTone2(t *testing.T) {
	py := PinyinTone2("中国")
	ep := pySlice{[]string{"zho1ng", "zho4ng"}, []string{"guo2"}}

	assert.Equal(t, ep, py)
}

func TestPinyinTone3(t *testing.T) {
	py := PinyinTone3("中国")
	ep := pySlice{[]string{"zhong1", "zhong4"}, []string{"guo2"}}

	assert.Equal(t, ep, py)
}

func TestLazyPinyinTone(t *testing.T) {
	py := LazyPinyinTone("中国")
	ep := lzSlice{"zhōng", "guó"}

	assert.Equal(t, 2, len(ep))
	assert.Equal(t, ep, py)
}

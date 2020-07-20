package uulang

import (
	_ "fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var loaded bool = false

func setup() {
	if !loaded {
		LoadSegDictionary("./data/dictionary.txt")
		loaded = true
	}
}

func TestCut(t *testing.T) {
	setup()

	swords := Cut("显示分词结果")
	ep := []string{"显示", "分词", "结果"}

	for k, v := range swords {
		assert.Equal(t, v.Word, ep[k])
	}
}

func TestCut4Search(t *testing.T) {
	setup()

	swords := Cut4Search("中华人民共和国")
	ep := strings.Split("中华 人民 共和 国 共和国 人民共和国 中华人民共和国", " ")

	for k, v := range swords {
		assert.Equal(t, v.Word, ep[k])
	}
}

func TestCutWithPinyin(t *testing.T) {
	setup()

	sword := CutWithPinyin("鞋濕了")

	assert.Equal(t, sword[0].PinyinT[0], 2)
	assert.Equal(t, sword[1].Pos, "unknown")
	assert.Equal(t, sword[2].Pos, "ul")
}

func TestCut2Words(t *testing.T) {
	setup()

	words := Cut2Words("中华人民共和国共和国", true)
	assert.Equal(t, "中华 人民 共和 国 共和国 人民共和国 中华人民共和国", strings.Join(words, " "))

	words = Cut2Words("中华人民共和国", false)
	assert.Equal(t, "中华人民共和国", strings.Join(words, " "))

	words = Cut2Words("昨天是阴天今天也是阴天", true)
	assert.Equal(t, "昨天 是 阴天 今天 也", strings.Join(words, " "))
}

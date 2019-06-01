package uulang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestS2T(t *testing.T) {
	text, _ := S2T("如果这都不算爱")
	assert.Equal(t, "如果這都不算愛", text)
}

func TestIsSimplified(t *testing.T) {
	assert.True(t, IsSimplified("新闻"))
	assert.False(t, IsSimplified("新聞"))
	assert.False(t, IsSimplified("可信度都值得懷疑"))
	assert.True(t, IsSimplified("可信度都值得"))
	assert.True(t, IsSimplified("你好，世界"))
	assert.True(t, IsSimplified("你好牛"))
}

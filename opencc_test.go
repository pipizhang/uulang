package uulang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestS2T(t *testing.T) {
	text, _ := S2T("如果这都不算爱")
	assert.Equal(t, "如果這都不算愛", text)
}

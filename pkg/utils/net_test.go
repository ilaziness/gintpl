package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInternalIP(t *testing.T) {
	ip, err := GetInternalIP()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "192.168.8.155", ip)
}

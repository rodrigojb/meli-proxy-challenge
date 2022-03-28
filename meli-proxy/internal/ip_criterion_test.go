package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIpCriterionApply(t *testing.T) {
	assert := assert.New(t)
	var r Request
	r.Host = "localhost"
	var c IpCriterion
	c.IP = "localhost"

	assert.True(c.ShouldApplied(r))
}

func TestIpCriterionDoesntApply(t *testing.T) {
	assert := assert.New(t)
	var r Request
	r.Host = "10.50.95.7"
	var c IpCriterion
	c.IP = "localhost"

	assert.False(c.ShouldApplied(r))
}

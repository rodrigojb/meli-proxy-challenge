package internal_test

import (
	"github.com/rodrigojb/meli-proxy/meli-proxy/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathCriterionApply(t *testing.T) {
	var r internal.Request
	r.Path = "/categories/MLA3530"
	var c internal.PathCriterion
	c.Path = "/categories"

	assert.True(t, c.ShouldApplied(r))
}

func TestPathCriterionDoesntApply(t *testing.T) {
	r := internal.Request{
		Path: "/sites/MLA/listing_exposures",
	}

	c := internal.PathCriterion{
		Path: "/categories",
	}

	assert.False(t, c.ShouldApplied(r))
}

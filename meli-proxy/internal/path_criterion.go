package internal

import (
	"strings"
	"time"
)

type PathCriterion struct {
	Path   string
	Limit  int
	Period time.Duration
}

func (c PathCriterion) ShouldApplied(req Request) bool {
	return strings.Contains(req.Path, c.Path)
}

func (c PathCriterion) GetKey() string {
	return c.Path
}

func (c PathCriterion) GetLimit() int {
	return c.Limit
}

func (c PathCriterion) GetPeriod() time.Duration {
	return c.Period * time.Second
}

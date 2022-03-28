package internal

import (
	"strings"
	"time"
)

type IpCriterion struct {
	IP     string
	Limit  int
	Period time.Duration
}

func (c IpCriterion) ShouldApplied(req Request) bool {
	return strings.Compare(req.Host, c.IP) == 0
}

func (c IpCriterion) GetKey() string {
	return c.IP
}

func (c IpCriterion) GetLimit() int {
	return c.Limit
}

func (c IpCriterion) GetPeriod() time.Duration {
	return c.Period * time.Second
}

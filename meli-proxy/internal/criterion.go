package internal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var TooManyRequestErr = errors.New("too many request")

type Criterion interface {
	ShouldApplied(req Request) bool
	GetKey() string
	GetLimit() int
	GetPeriod() time.Duration
}

type Request struct {
	Host string
	Path string
}

func LimitRequest(ctx context.Context, rdb redis.UniversalClient, criteria []Criterion, req Request) error {

	var appliedCriteria []Criterion
	for _, c := range criteria {
		if !c.ShouldApplied(req) {
			continue
		}

		appliedCriteria = append(appliedCriteria, c)

		val, err := rdb.Get(ctx, c.GetKey()).Result()

		if err == redis.Nil {
			continue
		}

		if err != nil {
			return fmt.Errorf("getting: %v", err)
		}

		cs := toCriterionState(val)

		if time.Now().After(cs.ExpiredAt) {
			continue
		}

		if cs.Count <= 1 {
			return TooManyRequestErr
		}
	}

	for _, c := range appliedCriteria {

		v, err := rdb.Get(ctx, c.GetKey()).Result()
		if err == redis.Nil {
			cs := CriterionState{
				Count:     c.GetLimit(),
				ExpiredAt: time.Now().Add(c.GetPeriod()),
			}

			val := fromCriterionState(cs)

			if _, err = rdb.SetNX(ctx, c.GetKey(), val, 0).Result(); err != nil {
				return fmt.Errorf("setting: %v", err)
			}
			continue
		}

		if err != nil {
			return fmt.Errorf("getting: %v", err)
		}

		cs := toCriterionState(v)
		cs.Count -= 1

		if time.Now().After(cs.ExpiredAt) {
			cs.Count = c.GetLimit()
			cs.ExpiredAt = time.Now().Add(c.GetPeriod())
		}

		val := fromCriterionState(cs)

		if err := rdb.Set(ctx, c.GetKey(), val, 0).Err(); err != nil {
			return fmt.Errorf("decrementing: %v", err)
		}
	}

	return nil
}

type CriterionState struct {
	Count     int
	ExpiredAt time.Time
}

func fromCriterionState(c CriterionState) string {
	m, _ := json.Marshal(c)
	return string(m)
}

func toCriterionState(s string) CriterionState {
	var c CriterionState
	_ = json.Unmarshal([]byte(s), &c)
	return c
}

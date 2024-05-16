package suncalc_test

import (
	"Lesson/suncalc"
	"testing"
)

func Test(t *testing.T) {
	rise, set := suncalc.Calc(nil, "Moscow")//вот тут вместо nil и нужно вставлять mock
}
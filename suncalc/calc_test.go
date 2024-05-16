package suncalc_test

import (
	"Lesson/suncalc"
	"Lesson/suncalc/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	mock := mocks.NewGeocoder(t)
	mock.On("GetCoordsByName", "Moscow").Return(55.7558, 37.6176, nil)
	rise, set := suncalc.Calc(mock, "Moscow")//вот тут вместо nil и нужно вставлять mock
	assert.Equal(t, "2000-01-01 05:59:15 +0000 UTC", rise.String())
	assert.Equal(t, "2000-01-01 13:06:05 +0000 UTC", set.String())
}
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type testOr struct {
	channels []<-chan interface{}
}

var testOrs = []testOr{
	{[]<-chan interface{}{sig(2 * time.Hour), sig(5 * time.Minute), sig(1 * time.Second), sig(1 * time.Hour), sig(1 * time.Minute)}},
}

func TestOr(t *testing.T) {
	for _, test := range testOrs {
		start := time.Now()
		<-or(test.channels...)
		assert.Less(t, time.Since(start)-time.Second, time.Millisecond*50)
	}
}
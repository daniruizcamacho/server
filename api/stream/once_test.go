package stream

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Execute(t *testing.T) {
	executeOnce := once{}
	executionOne := make(chan struct{})
	executionTwo := make(chan struct{})
	go executeOnce.Do(func() {
		executionOne <- struct{}{}
	})
	go executeOnce.Do(func() {
		executionTwo <- struct{}{}
	})

	select {
	case <-executionOne:
		// expected
	case <-executionTwo:
		// expected
	case <-time.After(100 * time.Millisecond):
		t.Fatal("one do should be executed")
	}

	select {
	case <-executionOne:
		t.Fatal("should only execute one")
	case <-executionTwo:
		t.Fatal("should only execute one")
	case <-time.After(100 * time.Millisecond):
		// expected
	}

	assert.False(t, executeOnce.mayExecute())
}

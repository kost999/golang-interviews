package channelsmerge

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func write(ch chan int, c int) {
	defer close(ch)
	for i := 0; i < c; i++ {
		ch <- i
	}
}

func TestCase1(t *testing.T) {
	c0 := make(chan int)
	c1 := make(chan int)
	c2 := make(chan int)

	go write(c0, 2)
	go write(c1, 5)
	go write(c2, 10)

	r := Merge(c0, c1, c2)

	count := 0
	for range r {
		count++
	}
	assert.Equal(t, 17, count)
}

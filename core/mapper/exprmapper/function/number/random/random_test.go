/*
Sniperkit-Bot
- Status: analyzed
*/

package random

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sniperkit/snk.fork.tibcosoftware-flogo-lib/core/mapper/exprmapper/expression"
)

var s = &Random{}

func TestSample(t *testing.T) {
	final1 := s.Eval(100)
	assert.NotNil(t, final1)
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`number.random(100000)`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	fmt.Println(v)
}

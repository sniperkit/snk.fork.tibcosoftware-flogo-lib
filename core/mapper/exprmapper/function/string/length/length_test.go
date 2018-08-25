/*
Sniperkit-Bot
- Status: analyzed
*/

package length

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sniperkit/snk.fork.tibcosoftware-flogo-lib/core/mapper/exprmapper/expression"
)

var s = &Length{}

func TestStaticFunc_String_length(t *testing.T) {
	final11 := s.Eval("TIBCO FLOGO")
	fmt.Println(final11)
	assert.Equal(t, 11, final11)

	final2 := s.Eval("你好, FLOGO")
	fmt.Println(final2)
	assert.Equal(t, 13, final2)
}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`string.length("seafood,name")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	assert.Equal(t, 12, v)
}

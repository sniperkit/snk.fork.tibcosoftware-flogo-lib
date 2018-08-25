/*
Sniperkit-Bot
- Status: analyzed
*/

package equalsignorecase

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sniperkit/snk.fork.tibcosoftware-flogo-lib/core/mapper/exprmapper/expression"
)

var s = &EqualsIgnoreCase{}

func TestStaticFunc_Starts_with(t *testing.T) {
	final1 := s.Eval("TIBCO FLOGO", "TIBCO")
	fmt.Println(final1)
	assert.Equal(t, false, final1)

	final2 := s.Eval("TIBCO", "tibco")
	fmt.Println(final2)
	assert.Equal(t, true, final2)

}

func TestExpression(t *testing.T) {
	fun, err := expression.ParseExpression(`string.equalsIgnoreCase("TIBCO FLOGO", "TIBCO FLOGO")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	assert.Equal(t, true, v)
}

func TestExpressionIgnoreCase(t *testing.T) {
	fun, err := expression.ParseExpression(`string.equalsIgnoreCase("TIBCO flogo", "TIBCO FLOGO")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.NotNil(t, v)
	assert.Equal(t, true, v)
}

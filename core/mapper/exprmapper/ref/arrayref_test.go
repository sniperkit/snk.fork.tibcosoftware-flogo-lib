/*
Sniperkit-Bot
- Status: analyzed
*/

package ref

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFieldNameFromArrayRef(t *testing.T) {
	s := "$.field.name"
	field := GetFieldNameFromArrayRef(s)
	assert.Equal(t, "field.name", field)
}

func TestGetFieldNameFromArrayRef2(t *testing.T) {
	s := "$.field"
	field := GetFieldNameFromArrayRef(s)
	assert.Equal(t, "field", field)
}

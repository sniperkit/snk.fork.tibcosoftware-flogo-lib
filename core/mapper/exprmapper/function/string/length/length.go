/*
Sniperkit-Bot
- Status: analyzed
*/

package length

import (
	"github.com/sniperkit/snk.fork.tibcosoftware-flogo-lib/core/mapper/exprmapper/expression/function"
	"github.com/sniperkit/snk.fork.tibcosoftware-flogo-lib/logger"
)

var log = logger.GetLogger("length-function")

type Length struct {
}

func init() {
	function.Registry(&Length{})
}

func (s *Length) GetName() string {
	return "length"
}

func (s *Length) GetCategory() string {
	return "string"
}

func (s *Length) Eval(str string) int {
	log.Debugf("Return the length of a string \"%s\"", str)
	var l int
	l = len(str)
	log.Debugf("Done calculating the length %d", l)
	return l
}

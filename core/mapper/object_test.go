/*
Sniperkit-Bot
- Status: analyzed
*/

package mapper

import (
	"testing"

	"github.com/sniperkit/snk.fork.tibcosoftware-flogo-lib/core/data"
)

func Test_evalExpr(t *testing.T) {
	evalExpr("{{1}}", nil, &data.BasicResolver{})
}

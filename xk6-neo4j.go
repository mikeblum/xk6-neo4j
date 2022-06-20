package neo4j

import (
	"go.k6.io/k6/js/modules"
	"github.com/mikeblum/xk6-neo4j/neo4j"
)

func init() {
	modules.Register("k6/x/neo4j", neo4j.New())
}

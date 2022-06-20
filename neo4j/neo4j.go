package neo4j

import (
	"go.k6.io/k6/js/modules"
)

// forked from the xk6-sql impl:
// https://github.com/grafana/xk6-sql/blob/9c6ae906d6eaec10344ea8bb73ec4b168505a914/sql.go

// RootModule is the global module object type. It is instantiated once per test
// run and will be used to create `k6/x/neo4j` module instances for each VU.
type RootModule struct{}

// Neo4j represents an instance of the Neo4j module for every VU.
type Neo4j struct {
	vu modules.VU
}

var (
	_ modules.Module   = &RootModule{}
	_ modules.Instance = &Neo4j{}
)

// NewModuleInstance implements the modules.Module interface to return
// a new instance for each VU.
func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &Neo4j{vu: vu}
}

// Exports implements the modules.Instance interface and returns the exports
// of the JS module.
func (neo4j *Neo4j) Exports() modules.Exports {
	return modules.Exports{Default: neo4j}
}

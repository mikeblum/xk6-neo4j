package neo4j

import (
	"fmt"
	"go.k6.io/k6/js/modules"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// forked from the xk6-sql impl:
// https://github.com/grafana/xk6-sql/blob/9c6ae906d6eaec10344ea8bb73ec4b168505a914/sql.go

// RootModule is the global module object type. It is instantiated once per test
// run and will be used to create `k6/x/neo4j` module instances for each VU.
type RootModule struct{}

// Neo4j represents an instance of the Neo4j module for every VU.
type Neo4j struct {
	conf 	*Conf
	driver  *neo4j.Driver
	vu 		modules.VU
}

var (
	_ modules.Module   = &RootModule{}
	_ modules.Instance = &Neo4j{}
)

// New returns a pointer to a new RootModule instance.
func New() *RootModule {
    return &RootModule{}
}

// NewModuleInstance implements the modules.Module interface to return
// a new instance for each VU.
func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &Neo4j{
		conf: NewConf(),
		vu: vu,
	}
}

// Exports implements the modules.Instance interface and returns the exports
// of the JS module.
func (neo4j *Neo4j) Exports() modules.Exports {
	return modules.Exports{Default: neo4j}
}

func (neo4j *Neo4j) OpenWithConf() (*Neo4j, error) {
	if neo4j.conf == nil {
		return nil, fmt.Errorf("no conf found")
	}
	return neo4j.Open(neo4j.getEnvOr(NEO4J_ADDR, NEO4J_DEFAULT_ADDR), 
		neo4j.getEnvOr(NEO4J_USERNAME, NEO4J_DEFAULT_USERNAME), 
		neo4j.getEnvOr(NEO4J_PASSWORD, NEO4J_DEFAULT_PASSWORD))
}

// Open a neo4j db connection:
// "neo4j://localhost:7687", "neo4j", "?????"
func (neo4j *Neo4j) Open(connectionString string, username string, password string) (*Neo4j, error) {
	// short-circuit if a driver is configured
	if neo4j.driver != nil {
		return neo4j, nil
	}
	driver, err := createDriver(connectionString, username, password)
	if err != nil {
		return nil, err
	}
	neo4j.driver = &driver
	return neo4j, err
}

// Verify a neo4j db connection
func (neo4j *Neo4j) Verify() (bool, error) {
	if neo4j.driver == nil {
		return false, fmt.Errorf("no driver configured")
	}
	driver := *neo4j.driver
	err := driver.VerifyConnectivity()
	return err == nil, err
}

func createDriver(connectionString string, username string, password string) (neo4j.Driver, error) {
	return neo4j.NewDriver(connectionString, neo4j.BasicAuth(username, password, ""))
}

// call on application exit
func (neo4j *Neo4j) Close() error {
	if neo4j.driver == nil {
		return fmt.Errorf("driver improperly closed")
	}
	driver := *neo4j.driver
	return driver.Close()
}

func (neo4j *Neo4j) getEnvOr(envVar string, defaultTo string) string {
	val, ok := neo4j.getEnv(envVar); if ok {
		return val
	} else {
		return defaultTo
	}
}

func (neo4j *Neo4j) getEnv(envVar string) (string, bool) {
	return neo4j.conf.GetString(envVar), neo4j.conf.IsSet(envVar)
}

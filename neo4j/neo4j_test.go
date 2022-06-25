package neo4j

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
)

type Neo4jTestSuite struct {
    suite.Suite
	neo4j *Neo4j 
}

func (s *Neo4jTestSuite) SetupTest() {
    s.neo4j = &Neo4j{
		conf: NewConf(),
	}
}

func (s *Neo4jTestSuite) TestGetEnvPasswordIsLocalhost() {
	pswd, ok := s.neo4j.getEnv(NEO4J_PASSWORD)
    assert.Equal(s.T(), "localhost", pswd)
    assert.True(s.T(), ok)
}

func TestNeo4jTestSuite(t *testing.T) {
    suite.Run(t, new(Neo4jTestSuite))
}

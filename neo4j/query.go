package neo4j

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// RETURN 1;
func (n *Neo4j) Return() (*int64, error) {
	driver := *n.driver
    session := driver.NewSession(neo4j.SessionConfig{})
    defer session.Close()
    result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var result, err = tx.Run("RETURN 1;", nil)
		// driver error
		if err != nil {
			return nil, err
		}

    	if record, err := result.Single(); err != nil {
			return nil, err
		} else {
			return record.Values[0].(int64), err
		}
	})

    if err != nil {
        return nil, err
    }
	count := result.(int64)
    return &count, nil
}

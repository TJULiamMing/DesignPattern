package bridge

import (
	"fmt"
	"math/rand"
)

type MySQLFetcher struct {
	Config string
}

func NewMySQLFetcher(config string) *MySQLFetcher {
	return &MySQLFetcher{
		Config: config,
	}
}

func (mf MySQLFetcher) Fetch(sql string) []interface{} {
	fmt.Println("MySQL Fetcher implemented SQL: " + sql)
	rows := make([]interface{}, 0)
	rows = append(rows, rand.Perm(10))
	return rows
}

type PostgreSQLFetcher struct {
	Config string
}

func NewPostgreSQLFetcher(config string) *PostgreSQLFetcher {
	return &PostgreSQLFetcher{
		Config: config,
	}
}

func (pf PostgreSQLFetcher) Fetch(sql string) []interface{} {
	fmt.Println("PostgreSQL Fetcher implemented SQL: " + sql)
	rows := make([]interface{}, 0)
	rows = append(rows, rand.Perm(10), rand.Perm(10))
	return rows
}

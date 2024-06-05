package bridge

type IDataFetcher interface {
	Fetch(sql string) []interface{}
}

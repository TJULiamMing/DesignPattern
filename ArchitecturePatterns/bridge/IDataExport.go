package bridge

type IDataExport interface {
	Fetcher(fetcher IDataFetcher)
	Export(sql string) interface{}
}

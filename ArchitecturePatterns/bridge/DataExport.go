package bridge

import (
	"fmt"
	"reflect"
)

type JsonExport struct {
	jsonFetcher IDataFetcher
}

func (je *JsonExport) Fetcher(fetcher IDataFetcher) {
	je.jsonFetcher = fetcher
}

func NewJsonExport(jsonFetcher IDataFetcher) IDataExport {
	return &JsonExport{jsonFetcher: jsonFetcher}
}

func (je *JsonExport) Export(sql string) interface{} {
	rows := je.jsonFetcher.Fetch(sql)
	fmt.Printf("jsonFetcher get %v rows\n", reflect.ValueOf(rows).Len())
	return rows
}

type CsvExport struct {
	csvFetcher IDataFetcher
}

func NewCsvExport(csvFetcher IDataFetcher) *CsvExport {
	return &CsvExport{csvFetcher: csvFetcher}
}

func (cse *CsvExport) Fetcher(fetcher IDataFetcher) {
	cse.csvFetcher = fetcher
}

func (cse *CsvExport) Export(sql string) interface{} {
	rows := cse.csvFetcher.Fetch(sql)
	fmt.Printf("csvFetcher get %v rows\n", reflect.ValueOf(rows).Len())
	return rows
}

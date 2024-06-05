package main

import "designPattern/ArchitecturePatterns/bridge"

func main() {
	mf := bridge.NewMySQLFetcher("mysql://root:@tcp(127.0.0.1:3306)/architecture")
	csvExport := bridge.NewCsvExport(mf)
	csvExport.Export("select")

	pf := bridge.NewPostgreSQLFetcher("postgres://root:@127.0.0.1:5432/architecture")
	csvExport.Fetcher(pf)
	csvExport.Export("drop table")

	jsonExport := bridge.NewJsonExport(mf)
	jsonExport.Export("select for update")
}

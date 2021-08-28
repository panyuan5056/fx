package xdb

func DbFind(config map[string]string) ([]string, map[string][]string, map[string][]map[string]interface{}, map[string]string) {
	manage := Run(config)
	tables := manage.Tables()
	schemas := map[string][]string{}
	data := map[string][]map[string]interface{}{}
	sqls := map[string]string{}
	for _, table := range tables {
		schemas[table] = manage.Xschemas(table)
		data[table] = manage.Data(table)
		sqls[table] = manage.CreateTableSql(table)
	}
	return tables, schemas, data, sqls
}

func DbWrite(config map[string]string, tables []string, datas map[string][]map[string]interface{}, sqls map[string]string) {
	manage := Run(config)
	for _, table := range tables {
		sql := sqls[table]
		data := datas[table]
		manage.Db.Raw(sql)
		manage.Db.Table(table).CreateInBatches(data, 1000)
	}
}

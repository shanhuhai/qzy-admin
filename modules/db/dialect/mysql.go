package dialect

type mysql struct {
	commonDialect
}

func (mysql) GetName() string {
	return "mysql"
}

func (mysql) ShowColumnsWithComment(schema, table string) string {
	return `SELECT 
			COLUMN_NAME Field, DATA_TYPE Type, IS_NULLABLE 'Null', COLUMN_KEY 'Key', COLUMN_DEFAULT 'Default', EXTRA Extra, COLUMN_COMMENT Comment 
		FROM information_schema.COLUMNS 
		WHERE 
			table_name = '` + table + `'
		AND
			table_schema = '` + schema + `'`
}

func (mysql) ShowColumns(table string) string {
	return "show columns in " + table
}

func (mysql) ShowTables() string {
	return "show tables"
}

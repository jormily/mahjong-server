package database

func DBStartup(dsn string,maxIdle int,maxOpen int,showSql bool) func() {
	return MustStartup(
		dsn,
		MaxIdleConns(maxIdle),
		MaxIdleConns(maxOpen),
		ShowSQL(showSql),
	)
}

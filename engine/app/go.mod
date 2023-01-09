module app

go 1.19

replace sevian.com/whendy => ./../whendy

require sevian.com/whendy v0.0.0-00010101000000-000000000000

replace sevian.com/tool => ./../tool

replace sevian.com/session => ./../session

require sevian.com/session v0.0.0-00010101000000-000000000000

replace sevian.com/memory => ./../session/memory

require sevian.com/memory v0.0.0-00010101000000-000000000000

require (
	
	sevian.com/tool v0.0.0-00010101000000-000000000000
)



require github.com/go-sql-driver/mysql v1.7.0 // indirect

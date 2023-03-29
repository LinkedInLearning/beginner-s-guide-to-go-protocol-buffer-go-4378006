module example.com/customerservice

go 1.18

require example.com/customer v0.0.0-00010101000000-000000000000

require github.com/mattn/go-sqlite3 v1.14.16 // indirect

replace example.com/customer => ./customer

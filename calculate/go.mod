module example.com/calculate

go 1.19

replace example.com/timeanddate => ../timeanddate

require (
	example.com/timeanddate v0.0.0-00010101000000-000000000000 // indirect
	example.com/utils v0.0.0-00010101000000-000000000000 // indirect
)

replace example.com/utils => ../utils

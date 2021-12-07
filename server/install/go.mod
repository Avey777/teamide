module install

go 1.15

replace (
	base => ../base
	config => ../config
	db => ../db
	server => ../server
	worker => ../worker
)

require (
	base v0.0.0-00010101000000-000000000000
	config v0.0.0-00010101000000-000000000000
	db v0.0.0-00010101000000-000000000000
)

module github.com/rogeruiz/csv-normalizer

go 1.16

require (
	github.com/gocarina/gocsv v0.0.0-20210408192840-02d7211d929d
	norm/padding v0.0.0-00010101000000-000000000000
	norm/timeshift v0.0.0-00010101000000-000000000000
)

replace norm/padding => ./padding

replace norm/timeshift => ./timeshift

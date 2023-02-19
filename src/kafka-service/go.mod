module xmservice.com/kafka

go 1.18

replace xmservice.com/entity => ../entity

replace xmservice.com/repository => ../repository

replace xmservice.com/migration => ../migration

require (
	github.com/segmentio/kafka-go v0.4.38
	xmservice.com/entity v0.0.0-00010101000000-000000000000
	xmservice.com/repository v0.0.0-00010101000000-000000000000
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	gorm.io/driver/postgres v1.4.8 // indirect
	gorm.io/gorm v1.24.5 // indirect
	xmservice.com/migration v0.0.0-00010101000000-000000000000 // indirect
)

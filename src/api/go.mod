module xmservice.com/api

go 1.18

replace xmservice.com/services => ../services

replace xmservice.com/entity => ../entity

replace xmservice.com/migration => ../migration

replace xmservice.com/repository => ../repository

require (
	github.com/gin-gonic/gin v1.8.2
	xmservice.com/auth v0.0.0-00010101000000-000000000000
	xmservice.com/services v0.0.0-00010101000000-000000000000
)

require (
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.11.1 // indirect
	github.com/goccy/go-json v0.9.11 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/segmentio/kafka-go v0.4.38 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	github.com/vladimiroff/jwt-go/v3 v3.2.1 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/net v0.6.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/postgres v1.4.8 // indirect
	gorm.io/gorm v1.24.5 // indirect
	xmservice.com/entity v0.0.0-00010101000000-000000000000 // indirect
	xmservice.com/kafka v0.0.0-00010101000000-000000000000 // indirect
	xmservice.com/migration v0.0.0-00010101000000-000000000000 // indirect
	xmservice.com/repository v0.0.0-00010101000000-000000000000 // indirect
)

replace xmservice.com/auth => ../auth

replace xmservice.com/kafka => ../kafka-service

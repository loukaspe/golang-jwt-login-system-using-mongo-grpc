module github.com/loukaspe/auth/proxy

go 1.19

replace proto => ./proto

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/golang-jwt/jwt/v4 v4.4.3
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.4.0
	github.com/sirupsen/logrus v1.9.0
	github.com/stretchr/testify v1.8.1
	google.golang.org/grpc v1.52.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

generate-proto:
	make build-utilities
	@docker run --volume "$(PWD)":/mongo-handler --workdir /mongo-handler \
	utilities-build /bin/bash -c "protoc --proto_path=proto --go_out=./proto --go_opt=paths=source_relative --go-grpc_out=./proto --go-grpc_opt=paths=source_relative proto/*/*.proto proto/*/*.proto"

generate-mock:
ifndef FILE
	@echo "FILE parameter is missing"
	@exit 1
endif
	make build-utilities
	@docker run --volume "$(PWD)":/mongo-handler --workdir /mongo-handler \
	utilities-build /bin/bash -c "mockgen -source=${FILE} -destination=mocks/mock_${FILE}"

tests-unit:
	make build-utilities
	@docker run \
			--rm \
			--volume "$(PWD)":/mongo-handler \
			--workdir /mongo-handler \
			utilities-build godotenv -f .env go test -short -cover -count=1 ./...

tests-benchmark:
	make build-utilities
	@docker run \
			--rm \
			--volume "$(PWD)":/mongo-handler \
			--workdir /mongo-handler \
			utilities-build godotenv -f .env go test ./... -bench=.

tests-file:
	make build-utilities
	@docker run \
			--rm \
			--volume "$(PWD)":/mongo-handler \
			--workdir /mongo-handler \
			utilities-build godotenv -f .env go test -count=1 -v ${FILE}

tests-package:
	make build-utilities
	@docker run \
			--rm \
			--volume "$(PWD)":/mongo-handler \
			--workdir /mongo-handler \
			utilities-build godotenv -f .env go test -count=1  ${PACKAGE}


tests-all:
	make build-utilities
	@docker run \
			--rm \
			--volume "$(PWD)":/mongo-handler \
			--workdir /mongo-handler \
			utilities-build godotenv -f .env go test ./... -cover -count=1

tests-all-with-coverage:
	make build-utilities
	@docker run \
			--rm \
			--volume "$(PWD)":/mongo-handler \
			--workdir /mongo-handler \
			utilities-build godotenv -f .env go test -count=1 -v -coverpkg=./... -coverprofile=profile.cov ./... ; go tool cover -func profile.cov


build-utilities:
	@docker build \
			--tag utilities-build \
			-f Dockerfile.utilities .

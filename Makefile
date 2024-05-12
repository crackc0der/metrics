up:
	docker-compose up -d 

down:
	@docker ps -aq | xargs docker rm -f || true
	# if [ -d "data" ]; then \
	# 	sudo rm -r data; \
	# fi

status:
	docker container ls

stoped:
	docker container ls -a

run-server: 
	cd server && make run &

lint-server:
	cd server && make lint

lint-client:
	cd client && make lint

gen-grpc-server:
	protoc --proto_path server/api/note_v1 --go_out=server/api/note_v1 --go_opt=paths=source_relative --go-grpc_out=server/api/note_v1 --go-grpc_opt=paths=source_relative server/api/note_v1/api.proto

gen-grpc-client:
	protoc --proto_path client/api/note_v1 --go_out=client/api/note_v1 --go_opt=paths=source_relative --go-grpc_out=client/api/note_v1 --go-grpc_opt=paths=source_relative client/api/note_v1/api.proto

golint-server:
	cd server/ && golangci-lint run --enable-all

golint-client:
	cd client/ && golangci-lint run --enable-all
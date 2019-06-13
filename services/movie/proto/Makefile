regenerate-docker:
	docker run --rm -it \
		-v `pwd`:/messages \
		-w /messages \
		ritterho/vss-micro-ci \
		make regenerate

regenerate:
	protoc -I=. --proto_path=${GOPATH}/src --micro_out=. --go_out=. movie.proto
		
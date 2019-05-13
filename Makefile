all: test

clean:
	rm -rf bin
	rm -rf tests/*_easyjson.go

build:
	go build -o bin/easyjson ./easyjson

generate: build
	bin/easyjson -stubs \
		./tests/snake.go \
		./tests/data.go \
		./tests/omitempty.go \
		./tests/nothing.go \
		./tests/named_type.go \
		./tests/custom_map_key_type.go \
		./tests/embedded_type.go \
		./tests/reference_to_pointer.go \

	bin/easyjson -all ./tests/data.go
	bin/easyjson -all ./tests/nothing.go
	bin/easyjson -all ./tests/errors.go
	bin/easyjson -snake_case ./tests/snake.go
	bin/easyjson -omit_empty ./tests/omitempty.go
	bin/easyjson -build_tags=use_easyjson ./benchmark/data.go
	bin/easyjson ./tests/nested_easy.go
	bin/easyjson ./tests/named_type.go
	bin/easyjson ./tests/custom_map_key_type.go
	bin/easyjson ./tests/embedded_type.go
	bin/easyjson ./tests/reference_to_pointer.go
	bin/easyjson ./tests/key_marshaler_map.go
	bin/easyjson -disallow_unknown_fields ./tests/disallow_unknown.go

golint:
	go build -o bin/golint golang.org/x/lint/golint

test: generate root golint
	go test \
		./tests \
		./jlexer \
		./gen \
		./buffer
	go test -benchmem -tags use_easyjson -bench . ./benchmark
	bin/golint -set_exit_status ./tests/*_easyjson.go

bench-other: generate root
	@go test -benchmem -bench . ./benchmark
	@go test -benchmem -tags use_ffjson -bench . ./benchmark
	@go test -benchmem -tags use_jsoniter -bench . ./benchmark
	@go test -benchmem -tags use_codec -bench . ./benchmark

bench-python:
	benchmark/ujson.sh


.PHONY: root clean generate test build

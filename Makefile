all: test

clean:
	rm -rf .root
	rm -rf tests/*_easyjson.go
	rm -rf benchmark/*_easyjson.go

build:
	go build -i -o .root/bin/easyjson ./easyjson

generate: build
	.root/bin/easyjson -stubs \
		./tests/snake.go \
		./tests/data.go \
		./tests/omitempty.go \
		./tests/nothing.go \
		./tests/named_type.go \
		./tests/custom_map_key_type.go \
		./tests/embedded_type.go \
		./tests/reference_to_pointer.go \

	.root/bin/easyjson -all ./tests/data.go
	.root/bin/easyjson -all ./tests/nothing.go
	.root/bin/easyjson -all ./tests/errors.go
	.root/bin/easyjson -snake_case ./tests/snake.go
	.root/bin/easyjson -omit_empty ./tests/omitempty.go
	.root/bin/easyjson -build_tags=use_easyjson ./benchmark/data.go
	.root/bin/easyjson ./tests/nested_easy.go
	.root/bin/easyjson ./tests/named_type.go
	.root/bin/easyjson ./tests/custom_map_key_type.go
	.root/bin/easyjson ./tests/embedded_type.go
	.root/bin/easyjson ./tests/reference_to_pointer.go
	.root/bin/easyjson ./tests/key_marshaler_map.go
	.root/bin/easyjson -disallow_unknown_fields ./tests/disallow_unknown.go

test: generate
	go test \
		./tests \
		./jlexer \
		./gen \
		./buffer
	cd benchmark && go test -benchmem -tags use_easyjson -bench .
	golint -set_exit_status ./tests/*_easyjson.go

bench-other: generate
	cd benchmark && make

bench-python:
	benchmark/ujson.sh


.PHONY: clean generate test build

internal/prototest/prototest.pb.go:
	protoc \
	--go_out=. \
		--go_opt=module=github.com/iterate/maskmerge \
	./internal/prototest/prototest.proto
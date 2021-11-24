.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/permission/healthy.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/permission/shared.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/permission/scope.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/permission/rule.proto

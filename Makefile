run:
	@go run main.go


install:
	@go install github.com/ServiceWeaver/weaver/cmd/weaver@latest


status:
	@weaver single status


dashboard:
	@weaver single dashboard

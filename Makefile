run:
	@SERVICEWEAVER_CONFIG=weaver.toml go run .


reverse:
	@curl "localhost:12345/hello?name=john"


install:
	@go install github.com/ServiceWeaver/weaver/cmd/weaver@latest


status:
	@weaver single status


dashboard:
	@weaver single dashboard


deploy:
	@go build
	@weaver multi deploy weaver.toml

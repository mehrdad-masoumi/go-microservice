run_app:
	air

run_otp_consumer:
	go run cmd/consumer/otp_consumer.go

run_swagger:
	swag init

run_test:
	go test ./...

run_mockery:
	mockery  --output ./mocks --dir ./ --all
template-generate:
	qtc -dir=./internal/templates/v1/ -skipLineComments
	git add .

generate-html: template-generate
	go run ./cmd/generate-companies/main.go
	git add .

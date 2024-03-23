generate-template:
	qtc -dir=./internal/templates/v1/ -skipLineComments
	git add .

generate-html: generate-template
	go run ./cmd/generate-companies/main.go
	git add .

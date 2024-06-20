generate-template:
	qtc -dir=./internal/templates/v1/ -skipLineComments
	git add .

generate-html: generate-template
	go run ./cmd/generate-main-pages/main.go
	go run ./cmd/generate-ukraine-universities-rating-by-faang/main.go
	git add .

generate-template:
	qtc -dir=./internal/templates/v1/ -skipLineComments
	git add .

generate-html: generate-template
	go run ./cmd/generate-main-pages/main.go
	go run ./cmd/generate-ukrainian-universities-rating-by-faang/main.go
	go run ./cmd/generate-ukrainian-courses-employment/main.go
	git add .

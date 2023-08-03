generate:
	@openapi-generator generate -i https://api.pdfgen.dev/v1/swagger.yaml -g go -c config.yaml -p pdfgen -o . --git-repo-id pdfgen-sdk-go --git-user-id pdfgen-dev

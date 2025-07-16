generate-client-openapi-generator:
	rm -rf ./client/*
	openapi-generator generate \
	  -i  /Users/xenia/go/moira/docs/swagger.yaml \
	  -g go \
	  -o pkg/client \
	  -c config.json


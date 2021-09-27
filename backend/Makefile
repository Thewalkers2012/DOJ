.PHONY: test server generateDocs serveDocs

test:
	go test -v cover ./..

server:
	go run main.go

generateDocs:
	swagger generate spec -o ./swagger/swagger.yaml
	
serveDocs:
	swagger serve -F=swagger --port 36666 ./swagger/swagger.yaml

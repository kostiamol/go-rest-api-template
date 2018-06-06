swaggen:
		@ cd ./cmd/go-rest-api-template && \
		swagger generate spec -o ../../swagger.json

swagserv: 
		@ swagger serve -p=8090 -F=swagger swagger.json


api:
	docker-compose exec -w /go/src/github.com/hgkcho/realWorldHttp api bash

serve:
	docker-compose exec -w /go/src/github.com/hgkcho/realWorldHttp/cmd/server api bash -c "go run ."

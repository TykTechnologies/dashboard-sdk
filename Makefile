dashboard:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v7.10.0 generate \
         --generator-name go \
        --input-spec /local/swagger.yml \
        --output /local/pkg/dashboard/ \
        --skip-overwrite \
        --git-host github.com \
        --git-user-id TykTechnologies \
         --git-repo-id dashboard-sdk \
         --package-name dashboard \
         --api-name-suffix API \
         --minimal-update \
        --global-property skipFormModel=true \
        --global-property skipFormModel=true \
        --global-property apis,apiTests=false,apiDocs=true \
        --global-property models,modelTests=true,modelDocs=true \
        --global-property supportingFiles \
        -c /local/config.json\
        --name-mappings _id=MID


	sudo rm -rf pkg/dashboard/model_server_variable.go
	sudo python3 file_replace.py
	sudo mv pkg/dashboard/docs/XTykApiGateway.md pkg/dashboard/docs/XTykapi.md
	sudo gofmt -s -w .
	go mod tidy



validate-swagger:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v7.10.0 validate \
	  -i /local/swagger.yml
tests:
	go test ./... -count=1  -cover





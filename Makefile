dashboard:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
         --generator-name go \
        --input-spec /local/swagger.yml \
        --output /local/pkg/dashboard/ \
        --skip-overwrite \
        --git-host github.com \
        --git-user-id TykTechnologies \
         --git-repo-id dashboard-sdk/pkg/dashboard \
         --package-name dashboard \
         --api-name-suffix API \
        --global-property skipFormModel=true \
        --global-property skipFormModel=true \
        --global-property apis,apiTests=false,apiDocs=true \
        --global-property models,modelTests=false,modelDocs=false \
        --global-property supportingFiles \
        --additional-properties generateInterfaces=true \
        --reserved-words-mappings _id=MID,interface=customInterface

	sudo rm -rf pkg/dashboard/go.mod pkg/dashboard/go.sum pkg/dashboard/model_server_variable.go
	sudo python3 file_replace.py
	sudo gofmt -s -w .



validate-swagger:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v6.2.0 validate \
	  -i /local/swagger.yml





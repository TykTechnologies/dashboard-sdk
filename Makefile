dashboard:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
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
        --global-property apis,apiTests=true,apiDocs=true \
        --global-property models,modelTests=true,modelDocs=true \
        --global-property supportingFiles \
        -c /local/config.json\
        --reserved-words-mappings _id=MID,interface=customInterface

	sudo rm -rf pkg/dashboard/model_server_variable.go
	sudo python3 file_replace.py
	sudo mv pkg/dashboard/docs/XTykapiGateway.md pkg/dashboard/docs/XTykapi.md
	sudo gofmt -s -w .



validate-swagger:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v6.2.0 validate \
	  -i /local/swagger.yml
tests:
	go test ./... -count=1  -cover





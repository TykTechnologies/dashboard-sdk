dashboard:
	 docker run --rm -v "${PWD}:/local" swaggerapi/swagger-codegen-cli-v3 generate \
        -i /local/swagger.yml \
        --additional-properties=packageName=dashboard \
        -l go \
        -o /local/dashboard







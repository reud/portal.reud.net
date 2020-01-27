echo "swagger codegen starting"
docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli generate \
    -i local/swagger.yaml \
    -l typescript-node \
    -o /local/frontend/src/gen
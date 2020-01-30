codegen() {
  # frontend codegen
  docker run --rm -v "${PWD}":/local swaggerapi/swagger-codegen-cli generate \
    -i local/swagger.yaml \
    -l typescript-node \
    -o /local/frontend/src/gen
  # backend codegen
  # TODO: mainの構成がわかって、自力でmain.goが作れたら --exclude-main を入れる
  swagger generate server -f swagger.yaml -t backend/gen --principal map[string]*json.RawMessage
}

allocator() {
  if [ $1 = "codegen" ]; then
    codegen
  else
    echo "usage script { codegen }"
  fi
}

allocator "$1"

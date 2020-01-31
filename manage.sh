#!/bin/bash
codegen() {
  # frontend codegen
  openapi-generator generate \
    -i swagger.yaml \
    -g typescript-axios \
    -o frontend/src/gen
  # backend codegen
  # TODO: mainの構成がわかって、自力でmain.goが作れたら --exclude-main を入れる
  swagger generate server -f swagger.yaml -t backend/gen --principal map[string]*json.RawMessage
}

allocator() {
  if [ $1 = "codegen" ]; then
    codegen
  elif [ $1 = "createdb" ]; then
    docker run -d --name local-reud-net-db -p 5432:5432 postgres:11.6
  elif [ $1 = "removedb" ]; then
    docker stop local-reud-net-db
    docker rm local-reud-net-db
  else
    echo "usage script { codegen }"
  fi
}

allocator "$1"

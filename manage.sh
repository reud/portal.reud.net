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
  elif [ $1 = "run" ]; then
    go build backend/cmd/portal-reud-net-server
    sudo ./portal-reud-net-server \
      --tls-certificate /etc/letsencrypt/live/portal-reud-net-backend.japaneast.cloudapp.azure.com/fullchain.pem \
      --tls-key /etc/letsencrypt/live/portal-reud-net-backend.japaneast.cloudapp.azure.com/privkey.pem \
      --tls-port 443 \
      --host 0.0.0.0
  else
    echo "usage script { codegen | createdb | removedb | run }"
  fi
}

allocator "$1"

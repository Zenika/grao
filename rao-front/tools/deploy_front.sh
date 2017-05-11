#!/bin/bash

ENV=$1

main(){
  getenv
  build
  deploy $HOST
}

getenv(){
  if  [ "$ENV" == "prod" ]
  then
    HOST=""
    # will throw an error
  else
    HOST="35.156.212.179"
    USER="ubuntu"
    DEST="/var/www/html"
    BUILD_PARAMS=""
  fi
}

build(){
  npm run build $BUILD_PARAMS
}

deploy(){
  rsync -r --chmod=u+rwx,g+rwx,o+r dist/* $USER@$HOST:$DEST
}

main

#!/bin/bash

COMMAND="$GOPATH/bin/godocdown"
DEST="documentation"
declare -a PACKAGES=(
  "document"
  "tree"
  "conv"
  "search"
)

setup(){
  if ! type "$COMMAND" >/dev/null 2>&1; then
      echo "INFO: installing godocdown please wait ..."
      go get github.com/robertkrimen/godocdown/godocdown
  else
      echo "INFO: godocdown is allready installed"
  fi
  mkdir -p "$DEST"
}

clean(){
  echo "INFO: cleaning documentation folder"
  if [ -f $DEST/README.md ]; then
    rm -f $DEST/README.md
  fi
}

doc(){
  for i in "${PACKAGES[@]}"
  do
      echo "INFO: generating documentation for $i package"
      "$COMMAND" "$i" >> "$DEST/README.md"
  done
}

main(){
  setup
  clean
  doc
}

main && echo "Documentation was appended to $DEST/README.md"

#!/bin/bash

# This util scripts generates package documentation
# as github mardown that can be committed to the
# repository

# To document a package use godoc comment in source file
# and add it to the $PACKAGE array below
# Packages will be appended to a README.md file in
# same order as they were declared in the $PACKAGE array

# /!\ Working installation of golang
# with an existing $GOPATH is assumed here
COMMAND="$GOPATH/bin/godocdown"
DIR="$(dirname $(realpath $0))"
DEST=$(realpath "$DIR/../_documentation")
SRC="$DIR/.."

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
      "$COMMAND" "$SRC/$i" >> "$DEST/README.md"
  done
}

main(){
  setup
  clean
  doc
}

main && echo "Documentation was appended to $DEST/README.md"

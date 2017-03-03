# Réponse aux Appels d'Offres

```
- P1: Lire le chemin racine des documents dans le fichier de config (ou variable d'env)
- P1: Au lancement : parcourir tous les documents et les indexer dans Algolia
- P1: Tous les soirs, parcourir les nouveaux documents et les indexer dans Algolia
- P2: services pour rechercher dans Algolia
- P2: offrir une IHM pour rechercher les documents via mots clés
 ```

## Set up

Install [go](https://golang.org/cmd/go/)<br>
Set up  [GOPATH](https://golang.org/doc/code.html#GOPATH) environment variable

```shell
ln -s $PWD $GOPATH/src/github.com/Zenika/rao
```
## Deps

```shell
go get golang.org/x/oauth2
go get github.com/stacktic/dropbox
go get -u github.com/sajari/docconv
go get github.com/JalfResi/justext
go get github.com/algolia/algoliasearch-client-go/algoliasearch
```

## Build

```shell
go build -o bin/rao
```

## Env

```shell
export RAO_DBX_KEY="dropbox_key"
export RAO_DBX_SECRET="dropbox_secret"
export RAO_DBX_TOKEN="dropbox_token"
export RAO_DBX_ROOT="dropbox_root_path"
export RAO_DOCD_PORT="docd_listening_port"
export RAO_ALGOLIA_ID="algolia_api_client_id"
export RAO_ALG_KEY="algolia_api_key"
```

## Run

```shell
## install docd server
pushd $GOPATH/src/github.com/sajari/docconv/docd && go install && popd
## launch docd server
nohup $GOPATH/bin/docd &
## run app
$GOPATH/src/github.com/Zenika/rao/bin/rao
```

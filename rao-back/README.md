# Réponse aux Appels d'Offres

```
- P1: Lire le chemin racine des documents dans le fichier de config (ou variable d'env)
- P1: Au lancement : parcourir tous les documents et les indexer dans Algolia
- P1: Tous les soirs, parcourir les nouveaux documents et les indexer dans Algolia
- P2: services pour rechercher dans Algolia
- P2: offrir une IHM pour rechercher les documents via mots clés
 ```
## Deps

```shell
go get golang.org/x/oauth2
go get github.com/stacktic/dropbox
go get -u github.com/sajari/docconv
go get github.com/JalfResi/justext
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
```

## Run

```shell
## install docd server
pushd $GOPATH/src/github.com/sajari/docconv/docd && go install && popd
## launch docd server
nohup $GOPATH/bin/docd &$
## run app
./bin/rao
```

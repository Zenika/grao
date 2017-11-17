# GRAO

## Set up

Install [go](https://golang.org/cmd/go/)<br>
Set up  [GOPATH](https://golang.org/doc/code.html#GOPATH) environment variable

```shell
ln -s $PWD $GOPATH/src/github.com/Zenika/RAO
```
## Deps

```shell
apt-get install tidy
apt-get install wv
apt-get install popplerutils
apt-get install unrtf
go get golang.org/x/oauth2
go get github.com/stacktic/dropbox
go get -u github.com/sajari/docconv
go get github.com/JalfResi/justext
go get github.com/algolia/algoliasearch-client-go/algoliasearch
go get github.com/robfig/cron
go get gopkg.in/square/go-jose.v2
go get github.com/auth0-community/go-auth0
go get github.com/rs/cors
go get -u github.com/gorilla/mux
```

## Build

```shell
go build -o bin/rao
```

## Env

```shell
export GRAO_APP_PORT="8090"
export GRAO_DBX_KEY="dropbox_key"
export GRAO_DBX_SECRET="dropbox_secret"
export GRAO_DBX_TOKEN="dropbox_token"
export GRAO_DBX_ROOT="dropbox_root_path"
export GRAO_DBX_CURSOR="cursor_file"
export GRAO_ALGOLIA_ID="algolia_api_client_id"
export GRAO_ALGOLIA_KEY="algolia_api_key"
export GRAO_LOG_FILE="/tmp/rao.log"
export GRAO_LOG_LEVEL="(DEBUG|WARNING|ERROR|FATAL)"
export GRAO_POLL_EVERY="@daily"
export RAO_POLL_FROM="rao_filter_regexp_string"
export BDC_POLL_FROM="bdc_filter_regexp_string"
export AUTH0_AUDIENCE="https://grao.zenika.com/api/v1"
export AUTH0_DOMAIN="zenika.eu.auth0.com"
export AUTH0_JWKS_URI="https://zenika.eu.auth0.com/.well-known/jwks.json"
export AUTH0_ISSUER="https://zenika.eu.auth0.com/"

```

## Indexes

Configuration of indexes can be performed by posting on the following endpoint

```
http://{host}/api/v1/{index_name}/settings
```

Payloads are available [in this repository](config)

## Run

```shell
## install docd server
pushd $GOPATH/src/github.com/sajari/docconv/docd && go install && popd
## launch docd server
nohup $GOPATH/bin/docd &
## run app
$GOPATH/src/github.com/Zenika/RAO/bin/rao
```

## Source Code Documentation

Source code doc is available [in this repository](_documentation)

The source code documentation README.md file is generated using the
docdown.sh script available [in this repository](_tools/docdown.sh)

## Other Source & Credits

[docconv](https://github.com/sajari/docconv) by Sajari

[dropbox/dropbox.go](https://github.com/stacktic/dropbox/blob/master/dropbox.go) by Arnaud Ysmal

[algoliasearch-client-go](https://github.com/algolia/algoliasearch-client-go) by Algolia

## Configure Auth0 login flow

 - [Create your auth0 api](https://manage.auth0.com/#/apis)
 
   - Name it 
   - Use identifier https://grao.zenika.com/api/v1
   - Keep RS256 algorithm for signing
   
 - Configure environment variables for your backend server
 
   - AUTH0_AUDIENCE will be https://grao.zenika.com/api/v1
   - AUTHO_DOMAIN will be our auth0 domain
   - AUTH0_JWKS_URI will be {AUTH0_DOMAIN}.well-known/jwks.json

*As we use RS256 AUTH0_JWKS_URI is where we get the public key used to sign the token*

 - Use a client on the same domain to access the API
 
 - Filter client from your domain
 
   - [Add a rule to Auth0](https://manage.auth0.com/#/rules)
    - Chose Email domain whitelist and restrict the list to our domain
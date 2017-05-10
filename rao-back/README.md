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
```

## Build

```shell
go build -o bin/rao
```

## Env

```shell
export GRAO_DBX_KEY="dropbox_key"
export GRAO_DBX_SECRET="dropbox_secret"
export GRAO_DBX_TOKEN="dropbox_token"
export GRAO_DBX_ROOT="dropbox_root_path"
export GRAO_DBX_CURSOR="cursor_file"
export GRAO_ALGOLIA_ID="algolia_api_client_id"
export GRAO_ALG_KEY="algolia_api_key"
export GRAO_LOG_FILE="/tmp/rao.log"
export GRAO_LOG_LEVEL="(DEBUG|WARNING|ERROR|FATAL)"
export GRAO_POLL_EVERY="@daily"
export RAO_POLL_FROM="rao_filter_regexp_string"
export BDC_POLL_FROM="bdc_filter_regexp_string"
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

Source code doc is available [here in the repository](documentation)

## Other Source & Credits

[docconv](https://github.com/sajari/docconv) by Sajari

[dropbox/dropbox.go](https://github.com/stacktic/dropbox/blob/master/dropbox.go) by Arnaud Ysmal

[algoliasearch-client-go](https://github.com/algolia/algoliasearch-client-go) by Algolia

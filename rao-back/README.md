# GRAO

## DESCRIPTION


>Grao is an application that allows indexing documents stored in dropbox. 
It provides a search engine user interface for various types of documents.

Basically the idea is to:
  - Extract contents of documents from dropbbox every 24 hours
  - Convert contents of those documents to full text chunks
  - Send it to the algolia platform for indexation
  - Provide search endpoints authenticated with OAuth/Auth0 for searching
  
## PREREQUISITES

Install [docker](https://docs.docker.com/install/)<br>
Install [docker-compose](https://docs.docker.com/compose/install/)<br>

Additionally you may want to use [make](https://www.gnu.org/software/make/)

## DEVELOPMENT

> Application relies on a couple of environment variables.
They define external services dependencies (Auth0, dropbox, algolia) and configuration parameters.
This configuration can be set in a `.env` file used by docker-compose

```shell
# Application server port
GRAO_APP_PORT=8090
# Credentials used to access the dropbox account
GRAO_DBX_KEY=dropbox_key
GRAO_DBX_SECRET=dropbox_secret
GRAO_DBX_TOKEN=secret_token
# Recursive polling of dropbox tree starts here
GRAO_DBX_ROOT=/zenika # that's an example and should be updated
# Cursor file is used to maintain state between polls and permit diff 
GRAO_DBX_CURSOR=/cursor
# Credentials used to access algolia
GRAO_ALGOLIA_ID=algolia_api_client_id
GRAO_ALGOLIA_KEY=algolia_api_key
# This where log ends
GRAO_LOG_FILE=rao.log
# This could be (DEBUG|INFOR|WARN|ERROR)
GRAO_LOG_LEVEL=DEBUG
# We don't poll continuously at the moment (we could) 
GRAO_POLL_EVERY=@daily
# Unfortunatly we need regular expressions to tell the application where documents are stored according to their types
## En gros dans chaque agence les commerciaux sont sensés ranger leurs document de tel ou tel type au même endroit,
## donc à partir de l'arborescence dropbox d'un commercial donné on peut déduire les expressions à définir ci après
RAO_POLL_FROM=BDC_FILTER_PATTERN =`(?i)^.+/_{1,2}clients(_|\s){1}(?P<Agence>[\w&\s]+)/(?P<Client>[^/]+)/(?P<Projet>[^/]+)/BON DE COMMANDE/(?P<Consultant>[^/]+)` # Calls for Bids / Appels d'offre
BDC_POLL_FROM=/ # Purchase Orders / Bons de commande
# This configuration is needed to use AUTH0 as an authentication proxy
AUTH0_AUDIENCE=https://grao.zenika.com/api/v1
AUTH0_DOMAIN=zenika.eu.auth0.com
AUTH0_JWKS_URI=https://zenika.eu.auth0.com/.well-known/jwks.json
AUTH0_ISSUER=https://zenika.eu.auth0.com/
# Application relies on docd for document to text conversions
DOCD_PORT=8888
DOCD_HOST=docd ## docd is provided as a docker image described in the project

```

### The docd container

Converting documents to fulltext is delegated to an external service called docd.
This service must be up and running in development as well and is provided as a docker container in the current repo.

### Go Tools

 - Development image uses [fresh](https://github.com/pilu/fresh) for hot build/reload of the application on file edition
 - Dependencies are managed with [dep](https://golang.github.io/dep/)
 
> The tools are provided by the docker image used for development.
Assuming that docker and docker-compose are installed there is no need to install fresh or dep.

### Managing project steps using docker-compose

  - build
    
    `MODE=BUILD docker-compose up dev`
    
  - start (dev mode)
  
    `MODE=DEV docker-compose up`
    
  - test
  
    `MODE=test docker-compose up dev`
    
     
### Managing project steps using make

  - build
    
    `make build`
    
  - start (dev mode)
    
    `make start`
    
  - test
    
    `make test`
    

## Deployment

  - Application build ends up in the _dist folder
  - You will need to set up application environment according to production use
  - Make sure docd is running for document to text conversions
  
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

package auth

import (
	"github.com/stacktic/dropbox"
	"os"
)

var db *dropbox.Dropbox = nil

func RequireDropboxClient() *dropbox.Dropbox {
	if db != nil {
		return db
	}
	key := os.Getenv("RAO_DBX_KEY")
	secret := os.Getenv("RAO_DBX_SECRET")
	token := os.Getenv("RAO_DBX_TOKEN")
	db := dropbox.NewDropbox()
	db.SetAppInfo(key, secret)
	db.SetAccessToken(token)
	return db
}

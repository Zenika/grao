package auth

import (
	"os"

	"github.com/stacktic/dropbox"
)

var db *dropbox.Dropbox = nil

func RequireDropboxClient() *dropbox.Dropbox {
	if db != nil {
		return db
	}
	key := os.Getenv("GRAO_DBX_KEY")
	secret := os.Getenv("GRAO_DBX_SECRET")
	token := os.Getenv("GRAO_DBX_TOKEN")
	db := dropbox.NewDropbox()
	db.SetAppInfo(key, secret)
	db.SetAccessToken(token)
	return db
}

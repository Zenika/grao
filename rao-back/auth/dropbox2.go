package auth

import (
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"os"
)

var dbx files.Client = nil

func RequireDropbox2Client() files.Client{
	if dbx != nil {
		return dbx
	}
	//key := os.Getenv("GRAO_DBX_KEY")
	//secret := os.Getenv("GRAO_DBX_SECRET")
	//token := os.Getenv("GRAO_DBX_TOKEN")
	//db := dropbox.NewDropbox()
	//db.SetAppInfo(key, secret)
	//db.SetAccessToken(token)
	//return db
	config := dropbox.Config{
		Token:    os.Getenv("GRAO_DBX_TOKEN"),
		LogLevel: dropbox.LogInfo, // if needed, set the desired logging level. Default is off
	}
	dbx := files.New(config)
	return dbx
}
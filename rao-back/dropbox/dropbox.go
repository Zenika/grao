package dropbox

import (
    "encoding/json"
    "github.com/Zenika/RAO/auth"
    "net/http"
    "log"
    "os"
    "io"
)

var db = auth.RequireDropboxClient()

func GetRootFolder(w http.ResponseWriter, r *http.Request){
  log.Println("getting root folder")
  rootFolder := os.Getenv("RAO_DBX_ROOT")
  if len(rootFolder) == 0 {
    rootFolder = "Zenika - Clients"
  }
  // db := auth.RequireDropboxClient()
  md, err:= db.Metadata(rootFolder, true, false, "", "", 1000)
  if err != nil {
    log.Fatal(err)
  }
  walk(rootFolder)
  j, err := json.Marshal(md)
  if err == nil {
    w.Write([]byte(j))
  } else {
    log.Fatal(err)
  }
}


func walk(root string){
  if len(root) == 0 {
    root = "Zenika - Clients"
  }
  // db := auth.RequireDropboxClient()
  entry, err:= db.Metadata(root, true, false, "", "", 0)
  if err != nil {
    log.Fatal(err)
  }
  contents := entry.Contents
  for _, e := range contents  {
    if !e.IsDir {
				log.Println(e.Path)
				log.Println(e.MimeType)
        return;
		}
		walk(e.Path)
	}
}


func download(src) {
	var input io.ReadCloser
	var err error
	if input, _, err = db.Download(src, "", 0); err != nil {
		log.Fatal(err)
	}
	defer input.Close()
  // ...
}

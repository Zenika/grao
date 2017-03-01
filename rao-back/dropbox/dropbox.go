package dropbox

import (
    "github.com/Zenika/RAO/auth"
    "github.com/Zenika/RAO/doc"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "bytes"
    "log"
    "os"
    "io"
)

var db = auth.RequireDropboxClient()

func GetRootFolder(w http.ResponseWriter, r *http.Request){
  rootFolder := os.Getenv("RAO_DBX_ROOT")
  md, err:= db.Metadata(rootFolder, true, false, "", "", 1000)
  check(err)
  walk(rootFolder)
  j, err := json.Marshal(md)
  check(err)
  w.Write([]byte(j))
}


func walk(root string){
  if len(root) == 0 {
    root = "Zenika - Clients"
  }
  entry, err:= db.Metadata(root, true, false, "", "", 0)
  check(err)
  contents := entry.Contents
  for _, e := range contents  {
    if !e.IsDir {
				log.Println(e.Path)
				log.Println(e.MimeType)
        rc, _ := download(e.Path)
        body := convert(rc, e.MimeType)
        log.Println(body)
        return
		}
		walk(e.Path)
	}
}

func convert(rc io.ReadCloser, mime string)(string){
  if "application/pdf" != mime {
    return "not a pdf"
  }
  buffer, err := ioutil.ReadAll(rc)
  defer rc.Close()
  check(err)
  body, _, err := doc.Convert(buffer, mime)
  check(err)
  return string(body[:])
}


func download(src string)(io.ReadCloser, int64) {
  reader, size, err:= db.Download(src, "", 0)
  check(err)
  return reader, size
}

func downloadToFile(src string, dst string) {
  var input io.ReadCloser
	var fd *os.File
  var err error
  fd, err = os.Create(dst)
  check(err)
  defer fd.Close()
  if input, _, err = db.Download(src, "", 0); err != nil {
		os.Remove(dst)
		log.Fatal(err)
	}
  defer input.Close()
  if _, err := io.Copy(fd, input); err != nil {
		os.Remove(dst)
	}
}

func streamToByte(stream io.Reader) []byte {
  buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

func streamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}


func check(err error) {
  if err != nil {
      log.Fatal(err)
  }
}

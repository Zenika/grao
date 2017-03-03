package docd

import (
    "github.com/Zenika/RAO/log"
    "mime/multipart"
    "net/textproto"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "bytes"
    "fmt"
    "os"
)

type ConversionResponse struct {
    Body string             `json:"body"`
    Meta map[string]string  `json:"meta"`
    MSecs uint32            `json:"msecs"`
}

var MIMES = []string { "application/pdf" }

type Docd struct {}

func (docd Docd) Convert(input []byte, mimeType string) ([]byte, error) {
    if "application/pdf" != mimeType {
      return make([]byte, 0), nil
    }
    port := os.Getenv("RAO_DOCD_PORT")
    if(len(port) == 0){
      port = "8888"
    }
    convertUrl := fmt.Sprintf("http://localhost:%v/convert", port)
    convertParam := "input"

    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)

    h := make(textproto.MIMEHeader)
    h.Set("Content-Disposition", `form-data; name="`+convertParam+`"; filename="noname"`)
    h.Set("Content-Type", mimeType)
    part, err := writer.CreatePart(h)
    if err != nil {
        log.Error(err, log.ERROR)
        return nil, err
    }
    _, err = part.Write(input)
    if err != nil {
        log.Error(err, log.ERROR)
        return nil, err
    }
    err = writer.Close()
    if err != nil {
      log.Error(err, log.ERROR)
      return nil, err
    }
    client := &http.Client{}

    request, err := http.NewRequest("POST", convertUrl, body)
    if err != nil {
        log.Error(err, log.ERROR)
        return nil, err
    }
    request.Header["Content-Type"] = []string{"multipart/form-data; boundary="+writer.Boundary()}
    resp, err := client.Do(request)
    if err != nil {
        log.Error(err, log.ERROR)
        return nil, err
    }
    defer resp.Body.Close()
    jsonBlob, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Error(err, log.ERROR)
        return nil, err
    }
    converted := new(ConversionResponse)
    err = json.Unmarshal(jsonBlob, &converted)
    if err != nil {
        log.Error(err, log.ERROR)
        return nil, err
    }
    return []byte(converted.Body), nil
}

func New() *Docd {
    return &Docd {}
}

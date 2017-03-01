package doc

import (
    "encoding/json"
    "io/ioutil"
    "bytes"
    "net/http"
    "mime/multipart"
    "net/textproto"
)

type ConversionResponse struct {
    Body string             `json:"body"`
    Meta map[string]string  `json:"meta"`
    MSecs uint32            `json:"msecs"`
}

func Convert(input []byte, mimeType string) ([]byte, map[string]string, error) {

    convertUrl := "http://localhost:8888/convert"
    convertParam := "input"

    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)

    h := make(textproto.MIMEHeader)
    h.Set("Content-Disposition", `form-data; name="`+convertParam+`"; filename="noname"`)
    h.Set("Content-Type", mimeType)
    part, err := writer.CreatePart(h)
    if err != nil {
        return nil, nil, err
    }
    _, err = part.Write(input)
    if err != nil {
        return nil, nil, err
    }
    err = writer.Close()
    if err != nil {
      return nil, nil, err
    }
    client := &http.Client{}

    request, err := http.NewRequest("POST", convertUrl, body)
    if err != nil {
        return nil, nil, err
    }
    request.Header["Content-Type"] = []string{"multipart/form-data; boundary="+writer.Boundary()}
    resp, err := client.Do(request)
    if err != nil {
        return nil, nil, err
    }
    defer resp.Body.Close()
    jsonBlob, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, nil, err
    }
    converted := new(ConversionResponse)
    err = json.Unmarshal(jsonBlob, &converted)
    if err != nil {
        return nil, nil, err
    }
    return []byte(converted.Body), converted.Meta, nil
}

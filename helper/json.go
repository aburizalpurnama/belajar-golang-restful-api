package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	decodeErr := decoder.Decode(result)
	PanicIfError(decodeErr)
}

func WriteToResponseBody(writer http.ResponseWriter, webResponse interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encodeErr := encoder.Encode(webResponse)
	PanicIfError(encodeErr)
}

package httprpc

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"net/http"
)

type Response struct {
	Data proto.Message
}

// TODO change err

type ResponseError struct {
	Code int
	Msg  string
}

type AbortError struct {
	Code int
	Msg  string
}

func (r Response) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)

	m := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}
	jsonBytes, err := m.Marshal(r.Data)

	if err != nil {
		return err
	}

	_, err = w.Write(jsonBytes)
	return err
}

func (r Response) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"application/json; charset=utf-8"}
	}
}

func ResponseErr(code int, err error) *ResponseError {
	return &ResponseError{
		Code: code,
		Msg:  err.Error(),
	}
}

func AbortErr(code int, err error) *AbortError {
	return &AbortError{
		Code: code,
		Msg:  err.Error(),
	}
}

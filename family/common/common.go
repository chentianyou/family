package common

import (
	"encoding/json"
	"github.com/go-errors/errors"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
)

func EncodeJson(obj interface{}) string {
	jsonStr, _ := json.Marshal(obj)
	return string(jsonStr)
}

func DecodeJson(str []byte, obj interface{}) error {
	err := json.Unmarshal(str, obj)
	if err != nil {
		return errors.New(err)
	}
	return err
}

func GetStringFromBody(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, math.MaxInt64))
	if err != nil {
		return nil, errors.New("从body中读取string失败")
	}
	if err := r.Body.Close(); err != nil {
		return nil, errors.New("无法关闭body的读操作符")
	}
	return body, nil
}

func HttpResponseMessage(w http.ResponseWriter, code int, msg string) {
	message := ResponseMessage{}
	message.Code = code
	message.Message = msg
	HttpResponse(w, message)
}

func HttpResponse(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	content, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
	}
	w.Write(content)
}
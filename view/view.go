package view

import (
	"CATechDojo/controller/request"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ReadUserRequest(r *http.Request) (request.UserRequest, error) {
	body := r.Body
	defer body.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, body); err != nil {
		log.Println(err)
		return request.UserRequest{}, err
	}

	var reqBody request.UserRequest
	if err := json.Unmarshal(buf.Bytes(), &reqBody); err != nil {
		log.Println(err)
		return request.UserRequest{}, err
	}

	return reqBody, nil
}

func ReadGachaDrawRequest(r *http.Request) (request.Times, error) {
	body := r.Body
	defer body.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, body); err != nil {
		log.Println(err)
		return request.Times{}, err
	}

	var reqBody request.Times
	if err := json.Unmarshal(buf.Bytes(), &reqBody); err != nil {
		log.Println(err)
		return request.Times{}, err
	}

	return reqBody, nil
}

func WriteResponse(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	res, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if _, err := w.Write(res); err != nil {
		return err
	}

	return nil
}

package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sderohan/go-auth-server/pkg/models"
	"github.com/sderohan/go-auth-server/pkg/validator"
)

func setStatusCode(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}

func setContentType(w http.ResponseWriter, key, contentType string) {
	w.Header().Add(key, contentType)
}

func getRespnse(data interface{}, statusCode uint, success bool) *models.AuthResponse {
	return &models.AuthResponse{
		Success:    success,
		StatusCode: statusCode,
		Data:       data,
	}
}

func GetRequestBody(req *http.Request) (*models.AuthRequest, error) {
	var authRequest *models.AuthRequest

	reqBody, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		return authRequest, nil
	}

	err = json.Unmarshal(reqBody, &authRequest)
	if err != nil {
		return authRequest, err
	}

	return authRequest, nil
}

func ValidateRequestContent(reqModel *models.AuthRequest) []*validator.IError {
	reqErrors := validator.Validate(*reqModel)
	if reqErrors != nil {
		return reqErrors
	}
	return nil
}

func WriteResponse(w http.ResponseWriter, data interface{}, contentTypeKey, contentTypeValue string, statusCode uint, success bool) {
	setContentType(w, contentTypeKey, contentTypeValue)
	setStatusCode(w, int(statusCode))
	response := getRespnse(data, statusCode, success)
	rawResponse, _ := MarshalResponse(response)
	w.Write(rawResponse)
}

func MarshalResponse(model interface{}) ([]byte, error) {
	return json.Marshal(model)
}

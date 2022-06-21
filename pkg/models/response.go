package models

type Response struct {
	Success      string      `json:"success"`
	ResponseCode uint        `json:"responseCode"`
	Data         interface{} `json:"data"`
}

package user

import ()

type SuccessAPIResponse struct {
	Message string `json:"message"`
	User    User   `json:"user"`
}

type ErrorAPIResponse struct {
	Message string `json:"message"`
}

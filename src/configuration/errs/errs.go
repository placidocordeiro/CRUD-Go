package errs;

import "net/http"

type Errs struct {
	Message string `json:"message"`;
	Err string `json:"error"`;
	Code int `json:"code"`;
	Causes []Causes `json:"causes"`;
};

type Causes struct {
	Field string `json:"field"`;
	Message string `json:"message"`;
};

func (r *Errs) Error() string {
	return r.Message;
}

func NewError(message string, err string, code int, causes []Causes) *Errs {
	return &Errs{
		Message: message,
		Err: err,
		Code: code,
		Causes: causes,
	};
}

func NewBadRequestError(message string) *Errs {
	return &Errs{
		Message: message,
		Err: "Bad Request",
		Code: http.StatusBadRequest,
		Causes: nil,
	};
}

func NewBadRequestValidationError(message string, causes []Causes) *Errs {
	return &Errs{
		Message: message,
		Err: "Bad Request",
		Code: http.StatusBadRequest,
		Causes: causes,
	};
}

func NewInternalServerError(message string) *Errs {
	return &Errs{
		Message: message,
		Err: "Internal Server Error",
		Code: http.StatusInternalServerError,
		Causes: nil,
	};
}

func NewUnauthorizedError(message string) *Errs {
	return &Errs{
		Message: message,
		Err: "Unauthorized",
		Code: http.StatusUnauthorized,
		Causes: nil,
	};
}

func NewForbiddenError(message string) *Errs {
	return &Errs{
		Message: message,
		Err: "Forbidden",
		Code: http.StatusForbidden,
		Causes: nil,
	};
}
package evaluationhttpadapter

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"useorigin.com/insurance-api/errors"
)

type evaluationHttpHandler struct{}

func NewEvaluationHandler() EvaluationHttpHandler {
	return &evaluationHttpHandler{}
}

func (c *evaluationHttpHandler) Evaluation(resp http.ResponseWriter, r *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	evaluation := new(Evaluation)

	err := json.NewDecoder(r.Body).Decode(evaluation)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(errors.NewError("error unmarshalling the request"))
		return
	}

	v := validator.New()
	err = v.Struct(evaluation)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		var error errors.Error
		for _, err := range err.(validator.ValidationErrors) {
			message := "validation error on " + err.Namespace()
			error.Message = append(error.Message, message)
		}
		json.NewEncoder(resp).Encode(error)
		return
	}
}

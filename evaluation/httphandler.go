package evaluation

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
)

type EvaluationHttpHandler interface {
	Evaluation(resp http.ResponseWriter, r *http.Request)
}

type evaluationHttpHandler struct{}

func NewEvaluationHandler() EvaluationHttpHandler {
	return &evaluationHttpHandler{}
}

func (c *evaluationHttpHandler) Evaluation(resp http.ResponseWriter, r *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	evaluation := new(Evaluation)

	err := json.NewDecoder(r.Body).Decode(evaluation)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		// ajustar
		resp.Write([]byte(`{"message": "error unmarshalling the request"}`))
		return
	}

	v := validator.New()
	err = v.Struct(evaluation)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		// ajustar
		//resp.Write([]byte(`{"message": "error unmarshalling the request"}`))
		json.NewEncoder(resp).Encode(Error{Message: err.Error()})
		return
	}
}

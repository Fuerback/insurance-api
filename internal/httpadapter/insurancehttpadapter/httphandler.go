package insurancehttpadapter

import (
	"encoding/json"
	"log"
	"net/http"
	"useorigin.com/insurance-api/internal/service/insuranceservice"

	"github.com/go-playground/validator"
	"useorigin.com/insurance-api/errors"
)

type evaluationHttpHandler struct {
	insurance insuranceservice.Insurance
}

func NewEvaluationHandler(insurance insuranceservice.Insurance) InsuranceHttpHandler {
	return &evaluationHttpHandler{insurance: insurance}
}

func (c *evaluationHttpHandler) EvaluateUserProfile(resp http.ResponseWriter, r *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	riskProfileRequest := new(RiskProfileRequest)

	log.Println("starting EvaluateUserProfile")

	err := json.NewDecoder(r.Body).Decode(riskProfileRequest)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(errors.NewError("error unmarshalling the request"))
		log.Println("EvaluateUserProfile - error unmarshalling the request")
		return
	}

	v := validator.New()
	err = v.Struct(riskProfileRequest)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		var error errors.Error
		for _, err := range err.(validator.ValidationErrors) {
			message := "validation error on " + err.Namespace()
			error.Message = append(error.Message, message)
		}
		json.NewEncoder(resp).Encode(error)
		log.Println("EvaluateUserProfile - error validating input")
		return
	}

	result := c.insurance.EvaluateUserProfile(riskProfileRequest.toDomain())
	resp.WriteHeader(http.StatusOK)
	log.Println("EvaluateUserProfile finished with success")
	json.NewEncoder(resp).Encode(result)
}

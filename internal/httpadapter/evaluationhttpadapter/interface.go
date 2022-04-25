package evaluationhttpadapter

import "net/http"

type EvaluationHttpHandler interface {
	Evaluation(resp http.ResponseWriter, r *http.Request)
}

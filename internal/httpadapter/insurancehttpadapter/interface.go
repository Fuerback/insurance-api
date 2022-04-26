package insurancehttpadapter

import "net/http"

type InsuranceHttpHandler interface {
	EvaluateUserProfile(resp http.ResponseWriter, r *http.Request)
}

package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"useorigin.com/insurance-api/config/env"
	"useorigin.com/insurance-api/errors"
	"useorigin.com/insurance-api/internal/httpadapter/evaluationhttpadapter"
	"useorigin.com/insurance-api/internal/service/evaluationservice"
	"useorigin.com/insurance-api/internal/service/rulesengine"
	"useorigin.com/insurance-api/server"
)

var (
	serverURL     string
	evaluationURL string
)

func TestMain(m *testing.M) {
	port := env.GetEnvWithDefaultAsString("PORT", ":8000")
	serverURL = "http://localhost" + port
	evaluationURL = serverURL + "/evaluation"

	service := evaluationservice.NewService()
	handler := evaluationhttpadapter.NewEvaluationHandler(service)
	go server.NewServer(handler).Run()

	os.Exit(m.Run())
}

func TestEvaluationInput(t *testing.T) {
	var tests = []struct {
		name                string
		eval                evaluationhttpadapter.UserInformation
		want, errorMessages int
	}{
		{
			"single with no house and vehicle",
			evaluationhttpadapter.NewEvaluation(1, 1, 1, rulesengine.SINGLE, []int{1, 0, 1}, nil, nil),
			http.StatusOK,
			0,
		},
		{
			"married with house and vehicle",
			evaluationhttpadapter.NewEvaluation(1, 1, 1, rulesengine.MARRIED, []int{1, 0, 1}, &evaluationhttpadapter.House{OwnershipStatus: rulesengine.OWNED}, &evaluationhttpadapter.Vehicle{Year: 2015}),
			http.StatusOK,
			0,
		},
		{
			"no house ownership status",
			evaluationhttpadapter.NewEvaluation(1, 1, 1, rulesengine.MARRIED, []int{1, 0, 1}, &evaluationhttpadapter.House{}, nil),
			http.StatusBadRequest,
			1,
		},
		{
			"no vehicle year",
			evaluationhttpadapter.NewEvaluation(1, 1, 1, rulesengine.MARRIED, []int{1, 0, 1}, &evaluationhttpadapter.House{OwnershipStatus: rulesengine.OWNED}, &evaluationhttpadapter.Vehicle{}),
			http.StatusBadRequest,
			1,
		},
		{
			"invalid age, dependents and income",
			evaluationhttpadapter.NewEvaluation(-1, -1, -1, rulesengine.MARRIED, []int{1, 0, 1}, &evaluationhttpadapter.House{OwnershipStatus: rulesengine.OWNED}, &evaluationhttpadapter.Vehicle{Year: 2015}),
			http.StatusBadRequest,
			3,
		},
		{
			"invalid martial status",
			evaluationhttpadapter.NewEvaluation(1, 1, 1, "unknown", []int{1, 0, 1}, &evaluationhttpadapter.House{OwnershipStatus: rulesengine.OWNED}, &evaluationhttpadapter.Vehicle{Year: 2015}),
			http.StatusBadRequest,
			1,
		},
		{
			"invalid ownership status",
			evaluationhttpadapter.NewEvaluation(1, 1, 1, rulesengine.MARRIED, []int{1, 0, 1}, &evaluationhttpadapter.House{OwnershipStatus: "unknown"}, &evaluationhttpadapter.Vehicle{Year: 2015}),
			http.StatusBadRequest,
			1,
		},
		{
			"incomplete risk questions",
			evaluationhttpadapter.NewEvaluation(1, 1, 1, rulesengine.MARRIED, []int{1, 0}, &evaluationhttpadapter.House{OwnershipStatus: rulesengine.OWNED}, &evaluationhttpadapter.Vehicle{Year: 2015}),
			http.StatusBadRequest,
			1,
		},
		{
			"no required fields",
			evaluationhttpadapter.UserInformation{},
			http.StatusBadRequest,
			5,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.name)
		t.Run(testname, func(t *testing.T) {
			payload, _ := json.Marshal(tt.eval)
			resp, err := http.Post(evaluationURL, "", bytes.NewBuffer(payload))

			assert.Nilf(t, err, "error when evaluation %s: %s", tt.name, err)

			assert.Equalf(t, resp.StatusCode, tt.want, "got %d, want %d", resp.StatusCode, tt.want)

			if resp.StatusCode == http.StatusBadRequest {
				var error errors.Error
				b, _ := ioutil.ReadAll(resp.Body)
				_ = json.Unmarshal(b, &error)
				assert.Equal(t, tt.errorMessages, len(error.Message), "length error messages are different than expected")
			}
		})
	}
}

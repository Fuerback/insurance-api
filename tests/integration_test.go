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
	"useorigin.com/insurance-api/internal/httpadapter/insurancehttpadapter"
	"useorigin.com/insurance-api/internal/service/insuranceservice"
	"useorigin.com/insurance-api/internal/service/rulesengine"
	"useorigin.com/insurance-api/server"
)

var (
	serverURL     string
	evaluationURL string
)

func TestMain(m *testing.M) {
	port := env.GetEnvWithDefaultAsString("PORT", ":8080")
	serverURL = "http://localhost" + port
	evaluationURL = serverURL + "/evaluation"

	engine := rulesengine.NewEngine()
	service := insuranceservice.NewService(engine)
	handler := insurancehttpadapter.NewEvaluationHandler(service)
	go server.NewServer(handler).Run()

	os.Exit(m.Run())
}

func TestEvaluationInput(t *testing.T) {
	var tests = []struct {
		name                string
		eval                insurancehttpadapter.RiskProfileRequest
		want, errorMessages int
	}{
		{
			"single with no house and vehicle",
			insurancehttpadapter.NewEvaluation(1, 1, 1, rulesengine.SINGLE, []int{1, 0, 1}, nil, nil),
			http.StatusOK,
			0,
		},
		{
			"married with house and vehicle",
			insurancehttpadapter.NewEvaluation(1, 1, 1, rulesengine.MARRIED, []int{1, 0, 1}, &insurancehttpadapter.House{OwnershipStatus: rulesengine.OWNED}, &insurancehttpadapter.Vehicle{Year: 2015}),
			http.StatusOK,
			0,
		},
		{
			"no house ownership status",
			insurancehttpadapter.NewEvaluation(1, 1, 1, rulesengine.MARRIED, []int{1, 0, 1}, &insurancehttpadapter.House{}, nil),
			http.StatusBadRequest,
			1,
		},
		{
			"no vehicle year",
			insurancehttpadapter.NewEvaluation(1, 1, 1, rulesengine.MARRIED, []int{1, 0, 1}, &insurancehttpadapter.House{OwnershipStatus: rulesengine.OWNED}, &insurancehttpadapter.Vehicle{}),
			http.StatusBadRequest,
			1,
		},
		{
			"invalid age, dependents and income",
			insurancehttpadapter.NewEvaluation(-1, -1, -1, rulesengine.MARRIED, []int{1, 0, 1}, &insurancehttpadapter.House{OwnershipStatus: rulesengine.OWNED}, &insurancehttpadapter.Vehicle{Year: 2015}),
			http.StatusBadRequest,
			3,
		},
		{
			"invalid martial status",
			insurancehttpadapter.NewEvaluation(1, 1, 1, "unknown", []int{1, 0, 1}, &insurancehttpadapter.House{OwnershipStatus: rulesengine.OWNED}, &insurancehttpadapter.Vehicle{Year: 2015}),
			http.StatusBadRequest,
			1,
		},
		{
			"invalid ownership status",
			insurancehttpadapter.NewEvaluation(1, 1, 1, rulesengine.MARRIED, []int{1, 0, 1}, &insurancehttpadapter.House{OwnershipStatus: "unknown"}, &insurancehttpadapter.Vehicle{Year: 2015}),
			http.StatusBadRequest,
			1,
		},
		{
			"incomplete risk questions",
			insurancehttpadapter.NewEvaluation(1, 1, 1, rulesengine.MARRIED, []int{1, 0}, &insurancehttpadapter.House{OwnershipStatus: rulesengine.OWNED}, &insurancehttpadapter.Vehicle{Year: 2015}),
			http.StatusBadRequest,
			1,
		},
		{
			"no required fields",
			insurancehttpadapter.RiskProfileRequest{},
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

func TestEvaluationResults(t *testing.T) {
	var tests = []struct {
		name                         string
		riskProfile                  insurancehttpadapter.RiskProfileRequest
		auto, home, life, disability string
	}{
		{
			"young single no car and vehicle one dependent",
			insurancehttpadapter.NewEvaluation(20, 1, 10000, rulesengine.SINGLE, []int{1, 0, 1}, nil, nil),
			rulesengine.INELIGIBLE,
			rulesengine.INELIGIBLE,
			rulesengine.REGULAR,
			rulesengine.REGULAR,
		},
		{
			"adult single with car and vehicle",
			insurancehttpadapter.NewEvaluation(31, 0, 1, rulesengine.SINGLE, []int{1, 0, 1}, &insurancehttpadapter.House{OwnershipStatus: rulesengine.OWNED}, &insurancehttpadapter.Vehicle{Year: 2015}),
			rulesengine.REGULAR,
			rulesengine.REGULAR,
			rulesengine.REGULAR,
			rulesengine.REGULAR,
		},
		{
			"no income married mortgaged new car",
			insurancehttpadapter.NewEvaluation(31, 0, 0, rulesengine.MARRIED, []int{1, 1, 1}, &insurancehttpadapter.House{OwnershipStatus: rulesengine.MORTGAGED}, &insurancehttpadapter.Vehicle{Year: 2020}),
			rulesengine.RESPONSIBLE,
			rulesengine.RESPONSIBLE,
			rulesengine.RESPONSIBLE,
			rulesengine.INELIGIBLE,
		},
		{
			"high income single owned old car",
			insurancehttpadapter.NewEvaluation(41, 1, 205000, rulesengine.SINGLE, []int{0, 0, 0}, &insurancehttpadapter.House{OwnershipStatus: rulesengine.OWNED}, &insurancehttpadapter.Vehicle{Year: 2010}),
			rulesengine.ECONOMIC,
			rulesengine.ECONOMIC,
			rulesengine.ECONOMIC,
			rulesengine.ECONOMIC,
		},
		{
			"old married mortgaged old car",
			insurancehttpadapter.NewEvaluation(61, 3, 150000, rulesengine.MARRIED, []int{1, 0, 0}, &insurancehttpadapter.House{OwnershipStatus: rulesengine.MORTGAGED}, &insurancehttpadapter.Vehicle{Year: 2010}),
			rulesengine.REGULAR,
			rulesengine.REGULAR,
			rulesengine.INELIGIBLE,
			rulesengine.INELIGIBLE,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.name)
		t.Run(testname, func(t *testing.T) {
			payload, _ := json.Marshal(tt.riskProfile)
			resp, err := http.Post(evaluationURL, "", bytes.NewBuffer(payload))

			assert.Nilf(t, err, "error when evaluation %s: %s", tt.name, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var response insuranceservice.InsuranceSuggest
			b, _ := ioutil.ReadAll(resp.Body)
			_ = json.Unmarshal(b, &response)

			assert.Equal(t, tt.home, response.Home)
			assert.Equal(t, tt.life, response.Life)
			assert.Equal(t, tt.auto, response.Auto)
			assert.Equal(t, tt.disability, response.Disability)
		})
	}
}

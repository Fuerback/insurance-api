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
	"useorigin.com/insurance-api/errors"
	"useorigin.com/insurance-api/evaluation"
	"useorigin.com/insurance-api/server"
)

var (
	serverURL     string
	evaluationURL string
)

func TestMain(m *testing.M) {
	serverURL = "http://localhost:8000" //os.Getenv("SERVER_URL")
	evaluationURL = serverURL + "/evaluation"

	go server.NewServer(evaluation.NewEvaluationHandler()).Run()

	os.Exit(m.Run())
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		name                string
		eval                evaluation.Evaluation
		want, errorMessages int
	}{
		{
			"single with no house and vehicle",
			evaluation.NewEvaluation(1, 1, 1, "single", []bool{true, false, true}, nil, nil),
			http.StatusOK,
			0,
		},
		{
			"married with house and vehicle",
			evaluation.NewEvaluation(1, 1, 1, "married", []bool{true, false, true}, &evaluation.House{OwnershipStatus: "owned"}, &evaluation.Vehicle{Year: 2015}),
			http.StatusOK,
			0,
		},
		{
			"no house ownership status",
			evaluation.NewEvaluation(1, 1, 1, "married", []bool{true, false, true}, &evaluation.House{}, nil),
			http.StatusBadRequest,
			1,
		},
		{
			"no vehicle year",
			evaluation.NewEvaluation(1, 1, 1, "married", []bool{true, false, true}, &evaluation.House{OwnershipStatus: "owned"}, &evaluation.Vehicle{}),
			http.StatusBadRequest,
			1,
		},
		{
			"invalid age, dependents and income",
			evaluation.NewEvaluation(-1, -1, -1, "married", []bool{true, false, true}, &evaluation.House{OwnershipStatus: "owned"}, &evaluation.Vehicle{Year: 2015}),
			http.StatusBadRequest,
			3,
		},
		{
			"invalid martial status",
			evaluation.NewEvaluation(1, 1, 1, "unknown", []bool{true, false, true}, &evaluation.House{OwnershipStatus: "owned"}, &evaluation.Vehicle{Year: 2015}),
			http.StatusBadRequest,
			1,
		},
		{
			"invalid ownership status",
			evaluation.NewEvaluation(1, 1, 1, "married", []bool{true, false, true}, &evaluation.House{OwnershipStatus: "unknown"}, &evaluation.Vehicle{Year: 2015}),
			http.StatusBadRequest,
			1,
		},
		{
			"incomplete risk questions",
			evaluation.NewEvaluation(1, 1, 1, "married", []bool{true, false}, &evaluation.House{OwnershipStatus: "owned"}, &evaluation.Vehicle{Year: 2015}),
			http.StatusBadRequest,
			1,
		},
		{
			"no required fields",
			evaluation.Evaluation{},
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

package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
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

func TestEvaluationInput(t *testing.T) {
	eval := evaluation.Evaluation{
		Age: func() *int {
			age := 10
			return &age
		}(),
		Dependents: func() *int {
			age := 10
			return &age
		}(),
		Income: func() *int {
			age := 10
			return &age
		}(),
		MartialStatus: "single",
		RiskQuestions: []int{0, 1, 0},
	}
	payload, _ := json.Marshal(eval)

	resp, err := http.Post(evaluationURL, "", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

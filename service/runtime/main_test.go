package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("USERS", "Users")         // ✅ テーブル名を設定
	os.Setenv("AWS_LOCALSTACK", "true") // ✅ LocalStack を使用する設定
}

func TestHandler(t *testing.T) {
	tests := []struct {
		name     string
		request  events.APIGatewayProxyRequest
		expected int
	}{
		{
			name: "GET request to /users",
			request: events.APIGatewayProxyRequest{
				HTTPMethod: "GET",
				Path:       "/users",
			},
			expected: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := handler(context.TODO(), tt.request)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, resp.StatusCode)

			var responseBody []map[string]string
			err = json.Unmarshal([]byte(resp.Body), &responseBody)
			assert.NoError(t, err)

			assert.NotEmpty(t, responseBody, "Response body should not be empty")
			log.Println("Test Response:", responseBody)
		})
	}
}

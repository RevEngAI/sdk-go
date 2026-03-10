package sdk

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	revengai "github.com/RevEngAI/sdk-go/v3"
)

func TestGetMeResponse_CreationFieldParsesRfc3339WithTimezone(t *testing.T) {
	jsonStr := `{
		"username": "testuser",
		"user_id": 1,
		"first_name": "Test",
		"last_name": "User",
		"email": "test@example.com",
		"creation": "2024-06-15T10:30:00+05:30",
		"tutorial_seen": true,
		"role": "USER"
	}`

	var response revengai.GetMeResponse
	err := json.Unmarshal([]byte(jsonStr), &response)
	require.NoError(t, err)

	expected := time.Date(2024, 6, 15, 10, 30, 0, 0, time.FixedZone("", 5*3600+30*60))
	assert.True(t, expected.Equal(response.GetCreation()), "expected %v, got %v", expected, response.GetCreation())
}

func TestGetMeResponse_CreationFieldParsesRfc3339WithUtcZ(t *testing.T) {
	jsonStr := `{
		"username": "testuser",
		"user_id": 1,
		"first_name": "Test",
		"last_name": "User",
		"email": "test@example.com",
		"creation": "2024-06-15T10:30:00Z",
		"tutorial_seen": true,
		"role": "USER"
	}`

	var response revengai.GetMeResponse
	err := json.Unmarshal([]byte(jsonStr), &response)
	require.NoError(t, err)

	expected := time.Date(2024, 6, 15, 10, 30, 0, 0, time.UTC)
	assert.True(t, expected.Equal(response.GetCreation()), "expected %v, got %v", expected, response.GetCreation())
}

func TestGetMeResponse_CreationFieldParsesRfc3339WithNegativeOffset(t *testing.T) {
	jsonStr := `{
		"username": "testuser",
		"user_id": 1,
		"first_name": "Test",
		"last_name": "User",
		"email": "test@example.com",
		"creation": "2024-06-15T10:30:00-04:00",
		"tutorial_seen": true,
		"role": "USER"
	}`

	var response revengai.GetMeResponse
	err := json.Unmarshal([]byte(jsonStr), &response)
	require.NoError(t, err)

	expected := time.Date(2024, 6, 15, 10, 30, 0, 0, time.FixedZone("", -4*3600))
	assert.True(t, expected.Equal(response.GetCreation()), "expected %v, got %v", expected, response.GetCreation())
}

func TestGetMeResponse_CreationFieldParsesRfc3339WithFractionalSecondsAndTimezone(t *testing.T) {
	jsonStr := `{
		"username": "testuser",
		"user_id": 1,
		"first_name": "Test",
		"last_name": "User",
		"email": "test@example.com",
		"creation": "2024-06-15T10:30:00.123456+02:00",
		"tutorial_seen": true,
		"role": "USER"
	}`

	var response revengai.GetMeResponse
	err := json.Unmarshal([]byte(jsonStr), &response)
	require.NoError(t, err)

	expected := time.Date(2024, 6, 15, 10, 30, 0, 123456000, time.FixedZone("", 2*3600))
	assert.True(t, expected.Equal(response.GetCreation()), "expected %v, got %v", expected, response.GetCreation())
}

func TestGetMeResponse_CreationFieldParsesRfc3339WithNanosecondPrecision(t *testing.T) {
	jsonStr := `{
		"username": "testuser",
		"user_id": 1,
		"first_name": "Test",
		"last_name": "User",
		"email": "test@example.com",
		"creation": "2024-06-15T10:30:00.123456789+03:00",
		"tutorial_seen": true,
		"role": "USER"
	}`

	var response revengai.GetMeResponse
	err := json.Unmarshal([]byte(jsonStr), &response)
	require.NoError(t, err)

	expected := time.Date(2024, 6, 15, 10, 30, 0, 123456789, time.FixedZone("", 3*3600))
	assert.True(t, expected.Equal(response.GetCreation()), "expected %v, got %v", expected, response.GetCreation())
}

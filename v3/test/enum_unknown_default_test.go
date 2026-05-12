package sdk

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	revengai "github.com/RevEngAI/sdk-go/v3"
)

func TestBinaryTaskStatus_UnknownValueDefaultsToOpenAPISentinel(t *testing.T) {
	var status revengai.BinaryTaskStatus
	err := json.Unmarshal([]byte(`"NOT_A_REAL_STATUS"`), &status)

	require.NoError(t, err)
	assert.Equal(t, revengai.BINARYTASKSTATUS_UNKNOWN_DEFAULT_OPEN_API, status)
}

func TestBinaryTaskStatus_KnownValueIsPreserved(t *testing.T) {
	var status revengai.BinaryTaskStatus
	err := json.Unmarshal([]byte(`"COMPLETED"`), &status)

	require.NoError(t, err)
	assert.Equal(t, revengai.BINARYTASKSTATUS_COMPLETED, status)
}

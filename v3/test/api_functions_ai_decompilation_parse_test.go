package sdk

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	revengai "github.com/RevEngAI/sdk-go/v3"
)

func loadTestData(t *testing.T, filename string) []byte {
	t.Helper()
	data, err := os.ReadFile("testdata/" + filename)
	require.NoError(t, err, "failed to read testdata/%s", filename)
	return data
}

func Test_DecompilationData_Parse(t *testing.T) {
	t.Run("Parse full successful response", func(t *testing.T) {
		rawJSON := loadTestData(t, "ai_decompilation_task_result_success.json")

		var result revengai.DecompilationData
		err := json.Unmarshal(rawJSON, &result)
		require.NoError(t, err)

		assert.Equal(t, "success", result.GetStatus())

		require.True(t, result.HasDecompilation())
		decompilation := result.GetDecompilation()
		assert.Contains(t, decompilation, "pre_cpp_init")
		assert.Contains(t, decompilation, "global_var_0 = global_var_1")
		assert.Contains(t, decompilation, "unknown_func_1(global_var_4")
	})

	t.Run("Parse response with null decompilation", func(t *testing.T) {
		rawJSON := loadTestData(t, "ai_decompilation_task_result_null_data.json")

		var result revengai.DecompilationData
		err := json.Unmarshal(rawJSON, &result)
		require.NoError(t, err)

		assert.Equal(t, "pending", result.GetStatus())
		assert.False(t, result.HasDecompilation(), "HasDecompilation should be false when value is null")
	})

	t.Run("Roundtrip marshal and unmarshal", func(t *testing.T) {
		rawJSON := loadTestData(t, "ai_decompilation_task_result_roundtrip.json")

		var original revengai.DecompilationData
		err := json.Unmarshal(rawJSON, &original)
		require.NoError(t, err)

		marshaled, err := json.Marshal(&original)
		require.NoError(t, err)

		var roundtripped revengai.DecompilationData
		err = json.Unmarshal(marshaled, &roundtripped)
		require.NoError(t, err)

		assert.Equal(t, original.GetStatus(), roundtripped.GetStatus())
		assert.Equal(t, original.GetDecompilation(), roundtripped.GetDecompilation())
	})

	t.Run("Parse response with missing optional fields", func(t *testing.T) {
		rawJSON := loadTestData(t, "ai_decompilation_task_result_missing_optional.json")

		var result revengai.DecompilationData
		err := json.Unmarshal(rawJSON, &result)
		require.NoError(t, err)

		assert.Equal(t, "pending", result.GetStatus())
		assert.False(t, result.HasDecompilation())
	})
}

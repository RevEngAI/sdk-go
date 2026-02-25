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

func Test_GetAiDecompilationTaskResult_Parse(t *testing.T) {
	t.Run("Parse full successful response", func(t *testing.T) {
		rawJSON := loadTestData(t, "ai_decompilation_task_result_success.json")

		var result revengai.BaseResponseGetAiDecompilationTask
		err := json.Unmarshal(rawJSON, &result)
		require.NoError(t, err)

		// Top-level fields
		assert.True(t, result.GetStatus())
		// message is null in JSON; the nullable wrapper marks it as set (key was present)
		// but the underlying value is nil
		assert.True(t, result.HasMessage(), "HasMessage() returns true because the key was present in JSON")
		assert.Equal(t, "", result.GetMessage(), "GetMessage() returns zero value when message is null")

		// Data
		require.True(t, result.HasData())
		data := result.GetData()

		assert.Equal(t, "success", data.Status)

		// Decompilation
		decompilation := data.Decompilation.Get()
		require.NotNil(t, decompilation)
		assert.Contains(t, *decompilation, "pre_cpp_init")
		assert.Contains(t, *decompilation, "global_var_0 = global_var_1")
		assert.Contains(t, *decompilation, "unknown_func_1(global_var_4")

		// Raw decompilation
		rawDecomp := data.RawDecompilation.Get()
		require.NotNil(t, rawDecomp)
		assert.Contains(t, *rawDecomp, "DISASM_FUNCTION_0")
		assert.Contains(t, *rawDecomp, "GLOBAL_VAR_0 = GLOBAL_VAR_1")
		assert.Contains(t, *rawDecomp, "UNMATCHED_FUNCTION_0(GLOBAL_VAR_4")

		// Function mapping
		require.Contains(t, data.FunctionMapping, "DISASM_FUNCTION_0")
		fm := data.FunctionMapping["DISASM_FUNCTION_0"]
		assert.Equal(t, "pre_cpp_init", fm.Name)
		require.NotNil(t, fm.IsExternal)
		assert.False(t, *fm.IsExternal)
		addrVal := fm.Addr.Get()
		require.NotNil(t, addrVal)

		// Function mapping full
		require.True(t, data.FunctionMappingFull.IsSet())
		fmFull := data.FunctionMappingFull.Get()
		require.NotNil(t, fmFull)

		// inverse_string_map is empty
		assert.Empty(t, fmFull.InverseStringMap)

		// inverse_function_map
		require.Contains(t, fmFull.InverseFunctionMap, "DISASM_FUNCTION_0")
		ifm := fmFull.InverseFunctionMap["DISASM_FUNCTION_0"]
		assert.Equal(t, "pre_cpp_init", ifm.Name)
		require.NotNil(t, ifm.IsExternal)
		assert.False(t, *ifm.IsExternal)

		// Empty maps in function_mapping_full
		assert.Empty(t, fmFull.UnmatchedFunctions)
		assert.Empty(t, fmFull.UnmatchedCustomTypes)
		assert.Empty(t, fmFull.UnmatchedStrings)
		assert.Empty(t, fmFull.UnmatchedVars)
		assert.Empty(t, fmFull.UnmatchedGoToLabels)
		assert.Empty(t, fmFull.UnmatchedCustomFunctionPointers)
		assert.Empty(t, fmFull.UnmatchedVariadicLists)
		assert.Empty(t, fmFull.UnmatchedEnums)
		assert.Empty(t, fmFull.Fields)
		assert.Empty(t, fmFull.UnmatchedExternalVars)

		// unmatched_global_vars
		require.Len(t, fmFull.UnmatchedGlobalVars, 8)

		expectedGlobalVars := map[string]string{
			"GLOBAL_VAR_0": "global_var_0",
			"GLOBAL_VAR_1": "global_var_1",
			"GLOBAL_VAR_2": "global_var_2",
			"GLOBAL_VAR_3": "global_var_3",
			"GLOBAL_VAR_4": "global_var_4",
			"GLOBAL_VAR_5": "global_var_5_ptr",
			"GLOBAL_VAR_6": "global_var_6_ptr",
			"GLOBAL_VAR_7": "global_var_7_ptr",
		}
		for key, expectedVal := range expectedGlobalVars {
			require.Contains(t, fmFull.UnmatchedGlobalVars, key)
			assert.Equal(t, expectedVal, fmFull.UnmatchedGlobalVars[key].Value, "mismatch for %s", key)
		}

		// Summary
		require.True(t, data.Summary.IsSet())
		assert.Contains(t, *data.Summary.Get(), "pre_cpp_init")

		// AI Summary
		require.True(t, data.AiSummary.IsSet())
		assert.Contains(t, *data.AiSummary.Get(), "pre_cpp_init")

		// Raw AI Summary
		require.True(t, data.RawAiSummary.IsSet())
		assert.Contains(t, *data.RawAiSummary.Get(), "<code>pre_cpp_init</code>")

		// Predicted function name
		require.True(t, data.PredictedFunctionName.IsSet())
		assert.Equal(t, "pre_initialization", *data.PredictedFunctionName.Get())

		// Meta
		meta := result.GetMeta()
		assert.NotNil(t, meta)
	})

	t.Run("Large addr values can be parsed successfully", func(t *testing.T) {
		rawJSON := loadTestData(t, "ai_decompilation_task_result_addr_overflow.json")

		var result revengai.BaseResponseGetAiDecompilationTask
		err := json.Unmarshal(rawJSON, &result)

		require.NoError(t, err, "parsing should not fail even if addr value exceeds int32 range")
	})

	t.Run("Parse response with null data", func(t *testing.T) {
		rawJSON := loadTestData(t, "ai_decompilation_task_result_null_data.json")

		var result revengai.BaseResponseGetAiDecompilationTask
		err := json.Unmarshal(rawJSON, &result)
		require.NoError(t, err)

		assert.True(t, result.GetStatus())
		// The nullable wrapper sets isSet=true when the key is present in JSON
		// (even if the value is null), so HasData() returns true.
		// However, the underlying value pointer is nil.
		assert.True(t, result.HasData(), "HasData() returns true because the key was present in JSON")
		dataPtr, ok := result.GetDataOk()
		assert.True(t, ok)
		assert.Nil(t, dataPtr, "data value should be nil when JSON value is null")
	})

	t.Run("Roundtrip marshal and unmarshal", func(t *testing.T) {
		rawJSON := loadTestData(t, "ai_decompilation_task_result_roundtrip.json")

		var original revengai.BaseResponseGetAiDecompilationTask
		err := json.Unmarshal(rawJSON, &original)
		require.NoError(t, err)

		marshaled, err := json.Marshal(&original)
		require.NoError(t, err)

		var roundtripped revengai.BaseResponseGetAiDecompilationTask
		err = json.Unmarshal(marshaled, &roundtripped)
		require.NoError(t, err)

		assert.Equal(t, original.GetStatus(), roundtripped.GetStatus())
		assert.Equal(t, original.GetData().Status, roundtripped.GetData().Status)
		assert.Equal(t, *original.GetData().Decompilation.Get(), *roundtripped.GetData().Decompilation.Get())
		assert.Equal(t, *original.GetData().RawDecompilation.Get(), *roundtripped.GetData().RawDecompilation.Get())
		assert.Equal(t, *original.GetData().PredictedFunctionName.Get(), *roundtripped.GetData().PredictedFunctionName.Get())
		assert.Equal(t, *original.GetData().Summary.Get(), *roundtripped.GetData().Summary.Get())
		assert.Equal(t, *original.GetData().AiSummary.Get(), *roundtripped.GetData().AiSummary.Get())
		assert.Equal(t, *original.GetData().RawAiSummary.Get(), *roundtripped.GetData().RawAiSummary.Get())

		originalFM := original.GetData().FunctionMapping["DISASM_FUNCTION_0"]
		roundtrippedFM := roundtripped.GetData().FunctionMapping["DISASM_FUNCTION_0"]
		assert.Equal(t, originalFM.Name, roundtrippedFM.Name)
		assert.Equal(t, *originalFM.IsExternal, *roundtrippedFM.IsExternal)
	})

	t.Run("Parse response with missing optional fields", func(t *testing.T) {
		rawJSON := loadTestData(t, "ai_decompilation_task_result_missing_optional.json")

		var result revengai.BaseResponseGetAiDecompilationTask
		err := json.Unmarshal(rawJSON, &result)
		require.NoError(t, err)

		data := result.GetData()
		assert.Equal(t, "pending", data.Status)

		// Decompilation and raw_decompilation are explicit null
		assert.True(t, data.Decompilation.IsSet())
		assert.Nil(t, data.Decompilation.Get())

		assert.True(t, data.RawDecompilation.IsSet())
		assert.Nil(t, data.RawDecompilation.Get())

		// function_mapping is empty
		assert.Empty(t, data.FunctionMapping)

		// function_mapping_full is explicit null
		assert.True(t, data.FunctionMappingFull.IsSet())
		assert.Nil(t, data.FunctionMappingFull.Get())

		// Optional fields not present at all
		assert.False(t, data.Summary.IsSet())
		assert.False(t, data.AiSummary.IsSet())
		assert.False(t, data.RawAiSummary.IsSet())
		assert.False(t, data.PredictedFunctionName.IsSet())
	})
}

package qontak

// Utility function to convert a slice of KeyValue to a map.
func convertKeyValueToMap(keyValues []KeyValue) []map[string]interface{} {
	result := make([]map[string]interface{}, len(keyValues))
	for i, kv := range keyValues {
		result[i] = map[string]interface{}{
			"key":   kv.Key,
			"value": kv.Value,
		}
	}
	return result
}

// Utility function to convert a slice of KeyValueText to a map.
func convertKeyValueTextToMap(keyValueTexts []KeyValueText) []map[string]interface{} {
	result := make([]map[string]interface{}, len(keyValueTexts))
	for i, kvt := range keyValueTexts {
		result[i] = map[string]interface{}{
			"key":        kvt.Key,
			"value_text": kvt.ValueText,
			"value":      kvt.Value,
		}
	}
	return result
}

// Utility function to convert a slice of ButtonMessage to a map.
func convertButtonsToMap(buttons []ButtonMessage) []map[string]interface{} {
	result := make([]map[string]interface{}, len(buttons))
	for i, button := range buttons {
		result[i] = map[string]interface{}{
			"index": button.Index,
			"type":  button.Type,
			"value": button.Value,
		}
	}
	return result
}

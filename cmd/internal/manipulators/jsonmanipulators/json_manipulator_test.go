package jsonmanipulators

import (
	"encoding/json"
	"fmt"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/manipulators"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/readers"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/writers"
	"testing"
)

func TestInvalidFile(t *testing.T) {
	jsonExample := "{whatever:\"value\"}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if manipulator.CanManipulate("/etc/config.doesnotexist") {
		t.Fatal("Must be able to manipulate JSON files")
	}
}

func TestInvalidJson(t *testing.T) {
	jsonExample := "{whatever:\"value\"}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}
}

func TestSetInvalidFile(t *testing.T) {
	jsonExample := "{\"whatever\":\"value\"}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	err := manipulator.SetValue("/etc/config.doesnotexist", "whatever", "newvalue")

	if err == nil {
		t.Fatal("This should have failed")
	}
}

func TestSetInvalidJson(t *testing.T) {
	jsonExample := "{whatever:\"value\"}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	err := manipulator.SetValue("/etc/config.json", "whatever", "newvalue")

	if err == nil {
		t.Fatal("This should have failed")
	}
}

func TestSetJsonStringField(t *testing.T) {
	jsonExample := "{\"whatever\":\"value\"}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever", "newvalue")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file")
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)

	value, ok := result["whatever"].(string)

	if !ok {
		t.Fatal("Value must be a string")
	}

	if value != "newvalue" {
		t.Fatal("Value must be set to \"newvalue\" (was: \"" + value + "\"")
	}
}

func TestSetJsonNumberField(t *testing.T) {
	jsonExample := "{\"whatever\":5}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever", "6")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file")
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)

	value, ok := result["whatever"].(float64)

	if !ok {
		t.Fatal("Value must be a float")
	}

	if value != 6 {
		t.Fatal("Value must be set to \"6\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestSetJsonNumberFieldWithString(t *testing.T) {
	jsonExample := "{\"whatever\":5}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever", "newvalue")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file")
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)

	value, ok := result["whatever"].(string)

	if !ok {
		t.Fatal("Value must be a string")
	}

	if value != "newvalue" {
		t.Fatal("Value must be set to \"newvalue\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestSetJsonBoolField(t *testing.T) {
	jsonExample := "{\"whatever\":true}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever", "false")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file")
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)

	value, ok := result["whatever"].(bool)

	if !ok {
		t.Fatal("Value must be a float")
	}

	if value {
		t.Fatal("Value must be set to \"false\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestSetJsonBoolFieldWithString(t *testing.T) {
	jsonExample := "{\"whatever\":true}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever", "newvalue")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file")
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)

	value, ok := result["whatever"].(string)

	if !ok {
		t.Fatal("Value must be a string")
	}

	if value != "newvalue" {
		t.Fatal("Value must be set to \"newvalue\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestSetJsonObjectField(t *testing.T) {
	jsonExample := "{\"whatever\":{\"whatever2\":true}}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever", "{\"whatever3\":6}")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file")
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)

	value, ok := result["whatever"].(map[string]any)

	if !ok {
		t.Fatal("Value must be a map")
	}

	value2, ok := value["whatever3"].(float64)

	if !ok {
		t.Fatal("Nested Value must be a map")
	}

	if value2 != 6 {
		t.Fatal("Nested value must be set to \"6\" (was: \"" + fmt.Sprint(value2) + "\"")
	}
}

func TestSetJsonObjectFieldWithString(t *testing.T) {
	jsonExample := "{\"whatever\":{\"whatever2\":true}}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever", "newvalue")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file")
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)

	value, ok := result["whatever"].(string)

	if !ok {
		t.Fatal("Value must be a string")
	}

	if value != "newvalue" {
		t.Fatal("Nested value must be set to \"newvalue\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestSetJsonArrayField(t *testing.T) {
	jsonExample := "{\"whatever\":[\"hi\"]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever", "[\"there\"]")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file")
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)

	value, ok := result["whatever"].([]any)

	if !ok {
		t.Fatal("Value must be an array")
	}

	value2, ok := value[0].(string)

	if !ok {
		t.Fatal("Nested Value must be a string")
	}

	if value2 != "there" {
		t.Fatal("Nested value must be set to \"there\" (was: \"" + fmt.Sprint(value2) + "\"")
	}
}

func TestSetJsonArrayFieldWithString(t *testing.T) {
	jsonExample := "{\"whatever\":[\"hi\"]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever", "there")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file")
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)

	value, ok := result["whatever"].(string)

	if !ok {
		t.Fatal("Value must be an string")
	}

	if value != "there" {
		t.Fatal("Nested value must be set to \"there\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestSetJsonArrayFieldWithArray(t *testing.T) {
	jsonExample := "{\"whatever\":[[\"hi\"]]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever:0", "[\"there\"]")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file")
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)

	value, ok := result["whatever"].([]any)

	if !ok {
		t.Fatal("Value must be an array")
	}

	value2, ok := value[0].([]any)

	if !ok {
		t.Fatal("Nested value must be an array")
	}

	if value2[0] != "there" {
		t.Fatal("Nested value must be set to \"there\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestSetJsonNewField(t *testing.T) {
	jsonExample := "{\"whatever\":\"value\"}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever2", "newvalue")

	if err != nil {
		t.Fatal("Failed to perform replacement")
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)
	value, ok := result["whatever2"].(string)

	if !ok {
		t.Fatal("Value must be a string")
	}

	if value != "newvalue" {
		t.Fatal("New value must be set to \"newvalue\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestSetJsonNewNumberField(t *testing.T) {
	jsonExample := "{\"whatever\":\"value\"}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever2", "5")

	if err != nil {
		t.Fatal("Failed to perform replacement")
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)
	value, ok := result["whatever2"].(string)

	if !ok {
		t.Fatal("Value must be a string")
	}

	if value != "5" {
		t.Fatal("New value must be set to \"5\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestSetJsonNewBooleanField(t *testing.T) {
	jsonExample := "{\"whatever\":\"value\"}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever2", "true")

	if err != nil {
		t.Fatal("Failed to perform replacement")
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)
	value, ok := result["whatever2"].(string)

	if !ok {
		t.Fatal("Value must be a string")
	}

	if value != "true" {
		t.Fatal("New value must be set to \"true\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestSetJsonArrayFieldIndex(t *testing.T) {
	jsonExample := "{\"whatever\":[\"hi\"]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever:0", "there")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)

	value, ok := result["whatever"].([]any)

	if !ok {
		t.Fatal("Value must be an array")
	}

	value2, ok := value[0].(string)

	if !ok {
		t.Fatal("Nested Value must be a string")
	}

	if value2 != "there" {
		t.Fatal("Nested value must be set to \"there\" (was: \"" + fmt.Sprint(value2) + "\"")
	}
}

func TestSetJsonNumberArrayFieldIndexWithNumber(t *testing.T) {
	jsonExample := "{\"whatever\":[10]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever:0", "20")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)

	value, ok := result["whatever"].([]any)

	if !ok {
		t.Fatal("Value must be an array")
	}

	value2, ok := value[0].(float64)

	if !ok {
		t.Fatal("Nested Value must be a string")
	}

	if value2 != 20 {
		t.Fatal("Nested value must be set to 20 (was: \"" + fmt.Sprint(value2) + "\"")
	}
}

func TestSetJsonArrayFieldIndexOutOfBounds(t *testing.T) {
	jsonExample := "{\"whatever\":[\"hi\"]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever:10", "there")

	if err == nil {
		t.Fatal("This should have failed")
	}
}

func TestSetJsonArrayFieldAgainstObject(t *testing.T) {
	jsonExample := "{\"whatever\":{\"hi\":\"there\"}}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever:10", "there")

	if err == nil {
		t.Fatal("This should have failed")
	}
}

func TestSetJsonArrayFieldDoubleIndex(t *testing.T) {
	jsonExample := "{\"whatever\":[\"hi\"]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever:0:0", "there")

	if err == nil {
		t.Fatal("This should have failed")
	}
}

func TestSetJsonArrayFieldIndexWithNumber(t *testing.T) {
	jsonExample := "{\"whatever\":[\"hi\"]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever:0", "10")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)

	value, ok := result["whatever"].([]any)

	if !ok {
		t.Fatal("Value must be an array")
	}

	value2, ok := value[0].(string)

	if !ok {
		t.Fatal("Nested Value must be a string")
	}

	if value2 != "10" {
		t.Fatal("Nested value must be set to \"10\" (was: \"" + fmt.Sprint(value2) + "\"")
	}
}

func TestSetJsonArrayFieldIndexWithBool(t *testing.T) {
	jsonExample := "{\"whatever\":[\"hi\"]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever:0", "true")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)

	value, ok := result["whatever"].([]any)

	if !ok {
		t.Fatal("Value must be an array")
	}

	value2, ok := value[0].(string)

	if !ok {
		t.Fatal("Nested Value must be a string")
	}

	if value2 != "true" {
		t.Fatal("Nested value must be set to \"true\" (was: \"" + fmt.Sprint(value2) + "\"")
	}
}

func TestSetJsonIntArrayFieldIndexWithInt(t *testing.T) {
	jsonExample := "{\"whatever\":[10]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever:0", "20")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)

	value, ok := result["whatever"].([]any)

	if !ok {
		t.Fatal("Value must be an array")
	}

	value2, ok := value[0].(float64)

	if !ok {
		t.Fatal("Nested Value must be a number")
	}

	if value2 != 20 {
		t.Fatal("Nested value must be set to 20 (was: \"" + fmt.Sprint(value2) + "\"")
	}
}

func TestSetJsonIntArrayFieldIndexWithString(t *testing.T) {
	jsonExample := "{\"whatever\":[10]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever:0", "blah")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)

	value, ok := result["whatever"].([]any)

	if !ok {
		t.Fatal("Value must be an array")
	}

	value2, ok := value[0].(string)

	if !ok {
		t.Fatal("Nested Value must be a string")
	}

	if value2 != "blah" {
		t.Fatal("Nested value must be set to blah (was: \"" + fmt.Sprint(value2) + "\"")
	}
}

func TestSetJsonBoolArrayFieldIndexWithBool(t *testing.T) {
	jsonExample := "{\"whatever\":[true]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever:0", "false")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)

	value, ok := result["whatever"].([]any)

	if !ok {
		t.Fatal("Value must be an array")
	}

	value2, ok := value[0].(bool)

	if !ok {
		t.Fatal("Nested Value must be a bool")
	}

	if value2 {
		t.Fatal("Nested value must be set to false (was: \"" + fmt.Sprint(value2) + "\"")
	}
}

func TestSetJsonBoolArrayFieldIndexWithString(t *testing.T) {
	jsonExample := "{\"whatever\":[true]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever:0", "blah")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)

	value, ok := result["whatever"].([]any)

	if !ok {
		t.Fatal("Value must be an array")
	}

	value2, ok := value[0].(string)

	if !ok {
		t.Fatal("Nested Value must be a string")
	}

	if value2 != "blah" {
		t.Fatal("Nested value must be set to blah (was: \"" + fmt.Sprint(value2) + "\"")
	}
}

func TestSetJsonObjectArrayFieldIndexWithObject(t *testing.T) {
	jsonExample := "{\"whatever\":[{\"whatever2\":\"hi\"}]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever:0", "{\"whatever3\":\"there\"}")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)

	value, ok := result["whatever"].([]any)

	if !ok {
		t.Fatal("Value must be an array")
	}

	value2, ok := value[0].(map[string]any)

	if !ok {
		t.Fatal("Nested Value must be a map")
	}

	if value2["whatever3"] != "there" {
		t.Fatal("Nested value must be set to \"there\" (was: \"" + fmt.Sprint(value2) + "\"")
	}
}

func TestSetJsonObjectArrayFieldIndexWithString(t *testing.T) {
	jsonExample := "{\"whatever\":[{\"whatever2\":\"hi\"}]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever:0", "there")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)

	value, ok := result["whatever"].([]any)

	if !ok {
		t.Fatal("Value must be an array")
	}

	value2, ok := value[0].(string)

	if !ok {
		t.Fatal("Nested Value must be a string")
	}

	if value2 != "there" {
		t.Fatal("Nested value must be set to \"there\" (was: \"" + fmt.Sprint(value2) + "\"")
	}
}

func TestSetJsonArrayArrayFieldIndexWithArray(t *testing.T) {
	jsonExample := "{\"whatever\":[[\"whatever2\",\"hi\"]]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever:0", "[\"whatever3\",\"there\"]")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)

	value, ok := result["whatever"].([]any)

	if !ok {
		t.Fatal("Value must be an array")
	}

	value2, ok := value[0].([]any)

	if !ok {
		t.Fatal("Nested Value must be a map")
	}

	if value2[0] != "whatever3" {
		t.Fatal("Nested value must be set to \"whatever3\" (was: \"" + fmt.Sprint(value2) + "\"")
	}
}

func TestSetJsonArrayArrayFieldIndexWithString(t *testing.T) {
	jsonExample := "{\"whatever\":[[\"whatever2\",\"hi\"]]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever:0", "there")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = json.Unmarshal([]byte((*writer.Output)["/etc/config.json"]), &result)

	value, ok := result["whatever"].([]any)

	if !ok {
		t.Fatal("Value must be an array")
	}

	value2, ok := value[0].(string)

	if !ok {
		t.Fatal("Nested Value must be a string")
	}

	if value2 != "there" {
		t.Fatal("Nested value must be set to \"there\" (was: \"" + fmt.Sprint(value2) + "\"")
	}
}

func TestSetJsonMissingNestedField(t *testing.T) {
	jsonExample := "{\"whatever\":\"value\"}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: &map[string]string{
			"/etc/config.json": jsonExample,
		},
	}
	manipulator := JsonManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: JsonUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.json") {
		t.Fatal("Must be able to manipulate JSON files")
	}

	err := manipulator.SetValue("/etc/config.json", "whatever:whatever2", "newvalue")

	if err == nil {
		t.Fatal("Should have failed to perform replacement")
	}
}

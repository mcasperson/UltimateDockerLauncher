package manipulators

import (
	"fmt"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/readers"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/writers"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestYamlInvalidFile(t *testing.T) {
	yamlExample := "whatever: \"value\""
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if manipulator.CanManipulate("/etc/config.doesnotexist") {
		t.Fatal("This should have failed")
	}
}

func TestYamlInvalidJson(t *testing.T) {
	yamlExample := "blah: hi\n- hi"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}
}

func TestYamlSetInvalidFile(t *testing.T) {
	yamlExample := "whatever: \"value\""
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	err := manipulator.SetValue("/etc/config.doesnotexist", "whatever", "newvalue")

	if err == nil {
		t.Fatal("This should have failed")
	}
}

func TestYamlSetInvalidJson(t *testing.T) {
	yamlExample := "blah: hi\n- hi"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever", "newvalue")

	if err == nil {
		t.Fatal("This should have failed")
	}
}

func TestYamlSetStringField(t *testing.T) {
	yamlExample := "whatever: \"value\""
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever", "newvalue")

	if err != nil {
		t.Fatal("Failed to manipulate YAML file")
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)

	value, ok := result["whatever"].(string)

	if !ok {
		t.Fatal("Value must be a string")
	}

	if value != "newvalue" {
		t.Fatal("Value must be set to \"newvalue\" (was: \"" + value + "\"")
	}
}

func TestYamlSetNumberField(t *testing.T) {
	yamlExample := "{\"whatever\":5}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever", "6")

	if err != nil {
		t.Fatal("Failed to manipulate YAML file")
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)

	value, ok := result["whatever"].(int)

	if !ok {
		t.Fatal("Value must be a int")
	}

	if value != 6 {
		t.Fatal("Value must be set to \"6\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestYamlSetNumberFieldWithString(t *testing.T) {
	yamlExample := "{\"whatever\":5}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever", "newvalue")

	if err != nil {
		t.Fatal("Failed to manipulate YAML file")
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)

	value, ok := result["whatever"].(string)

	if !ok {
		t.Fatal("Value must be a string")
	}

	if value != "newvalue" {
		t.Fatal("Value must be set to \"newvalue\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestYamlSetBoolField(t *testing.T) {
	yamlExample := "{\"whatever\":true}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever", "false")

	if err != nil {
		t.Fatal("Failed to manipulate YAML file")
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)

	value, ok := result["whatever"].(bool)

	if !ok {
		t.Fatal("Value must be a bool")
	}

	if value {
		t.Fatal("Value must be set to \"false\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestYamlSetBoolFieldWithString(t *testing.T) {
	yamlExample := "{\"whatever\":true}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever", "newvalue")

	if err != nil {
		t.Fatal("Failed to manipulate YAML file")
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)

	value, ok := result["whatever"].(string)

	if !ok {
		t.Fatal("Value must be a string")
	}

	if value != "newvalue" {
		t.Fatal("Value must be set to \"newvalue\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestYamlSetObjectField(t *testing.T) {
	yamlExample := "whatever:\n  whatever2: true"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever", "whatever3: 6")

	if err != nil {
		t.Fatal("Failed to manipulate YAML file")
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)

	value, ok := result["whatever"].(map[string]any)

	if !ok {
		t.Fatal("Value must be a map")
	}

	value2, ok := value["whatever3"].(int)

	if !ok {
		t.Fatal("Nested Value must be a int")
	}

	if value2 != 6 {
		t.Fatal("Nested value must be set to \"6\" (was: \"" + fmt.Sprint(value2) + "\"")
	}
}

func TestYamlSetObjectFieldWithString(t *testing.T) {
	yamlExample := "{\"whatever\":{\"whatever2\":true}}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever", "newvalue")

	if err != nil {
		t.Fatal("Failed to manipulate YAML file")
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)

	value, ok := result["whatever"].(string)

	if !ok {
		t.Fatal("Value must be a string")
	}

	if value != "newvalue" {
		t.Fatal("Nested value must be set to \"newvalue\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestYamlSetArrayField(t *testing.T) {
	yamlExample := "whatever: [\"hi\"]"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever", "[\"there\"]")

	if err != nil {
		t.Fatal("Failed to manipulate YAML file")
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)

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

func TestYamlSetArrayFieldWithString(t *testing.T) {
	yamlExample := "{\"whatever\":[\"hi\"]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever", "there")

	if err != nil {
		t.Fatal("Failed to manipulate YAML file")
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)

	value, ok := result["whatever"].(string)

	if !ok {
		t.Fatal("Value must be an string")
	}

	if value != "there" {
		t.Fatal("Nested value must be set to \"there\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestYamlSetArrayFieldWithArray(t *testing.T) {
	yamlExample := "{\"whatever\":[[\"hi\"]]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever:0", "[\"there\"]")

	if err != nil {
		t.Fatal("Failed to manipulate YAML file")
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)

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

func TestYamlSetNewField(t *testing.T) {
	yamlExample := "whatever: \"value\""
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever2", "newvalue")

	if err != nil {
		t.Fatal("Failed to perform replacement")
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)
	value, ok := result["whatever2"].(string)

	if !ok {
		t.Fatal("Value must be a string")
	}

	if value != "newvalue" {
		t.Fatal("New value must be set to \"newvalue\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestYamlSetNewNumberField(t *testing.T) {
	yamlExample := "whatever: \"value\""
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever2", "5")

	if err != nil {
		t.Fatal("Failed to perform replacement")
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)
	value, ok := result["whatever2"].(string)

	if !ok {
		t.Fatal("Value must be a string")
	}

	if value != "5" {
		t.Fatal("New value must be set to \"5\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestYamlSetNewBooleanField(t *testing.T) {
	yamlExample := "whatever: \"value\""
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever2", "true")

	if err != nil {
		t.Fatal("Failed to perform replacement")
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)
	value, ok := result["whatever2"].(string)

	if !ok {
		t.Fatal("Value must be a string")
	}

	if value != "true" {
		t.Fatal("New value must be set to \"true\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestYamlSetArrayFieldIndex(t *testing.T) {
	yamlExample := "{\"whatever\":[\"hi\"]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever:0", "there")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)

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

func TestYamlSetNumberArrayFieldIndexWithNumber(t *testing.T) {
	yamlExample := "{\"whatever\":[10]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever:0", "20")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)

	value, ok := result["whatever"].([]any)

	if !ok {
		t.Fatal("Value must be an array")
	}

	value2, ok := value[0].(int)

	if !ok {
		t.Fatal("Nested Value must be a int")
	}

	if value2 != 20 {
		t.Fatal("Nested value must be set to 20 (was: \"" + fmt.Sprint(value2) + "\"")
	}
}

func TestYamlSetArrayFieldIndexOutOfBounds(t *testing.T) {
	yamlExample := "{\"whatever\":[\"hi\"]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever:10", "there")

	if err == nil {
		t.Fatal("This should have failed")
	}
}

func TestYamlSetArrayFieldAgainstObject(t *testing.T) {
	yamlExample := "{\"whatever\":{\"hi\":\"there\"}}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever:10", "there")

	if err == nil {
		t.Fatal("This should have failed")
	}
}

func TestYamlSetArrayFieldDoubleIndex(t *testing.T) {
	yamlExample := "{\"whatever\":[\"hi\"]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever:0:0", "there")

	if err == nil {
		t.Fatal("This should have failed")
	}
}

func TestYamlSetArrayFieldIndexWithNumber(t *testing.T) {
	yamlExample := "{\"whatever\":[\"hi\"]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever:0", "10")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)

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

func TestYamlSetArrayFieldIndexWithBool(t *testing.T) {
	yamlExample := "{\"whatever\":[\"hi\"]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever:0", "true")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)

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

func TestYamlSetIntArrayFieldIndexWithInt(t *testing.T) {
	yamlExample := "{\"whatever\":[10]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever:0", "20")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)

	value, ok := result["whatever"].([]any)

	if !ok {
		t.Fatal("Value must be an array")
	}

	value2, ok := value[0].(int)

	if !ok {
		t.Fatal("Nested Value must be a number")
	}

	if value2 != 20 {
		t.Fatal("Nested value must be set to 20 (was: \"" + fmt.Sprint(value2) + "\"")
	}
}

func TestYamlSetIntArrayFieldIndexWithString(t *testing.T) {
	yamlExample := "{\"whatever\":[10]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever:0", "blah")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)

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

func TestYamlSetBoolArrayFieldIndexWithBool(t *testing.T) {
	yamlExample := "{\"whatever\":[true]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever:0", "false")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)

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

func TestYamlSetBoolArrayFieldIndexWithString(t *testing.T) {
	yamlExample := "{\"whatever\":[true]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever:0", "blah")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)

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

func TestYamlSetObjectArrayFieldIndexWithObject(t *testing.T) {
	yamlExample := "{\"whatever\":[{\"whatever2\":\"hi\"}]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever:0", "{\"whatever3\":\"there\"}")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)

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

func TestYamlSetObjectArrayFieldIndexWithString(t *testing.T) {
	yamlExample := "{\"whatever\":[{\"whatever2\":\"hi\"}]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever:0", "there")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)

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

func TestYamlSetArrayArrayFieldIndexWithArray(t *testing.T) {
	yamlExample := "{\"whatever\":[[\"whatever2\",\"hi\"]]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever:0", "[\"whatever3\",\"there\"]")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)

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

func TestYamlSetArrayArrayFieldIndexWithString(t *testing.T) {
	yamlExample := "{\"whatever\":[[\"whatever2\",\"hi\"]]}"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever:0", "there")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = yaml.Unmarshal([]byte(writer.Output["/etc/config.yaml"]), &result)

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

func TestYamlSetMissingNestedField(t *testing.T) {
	yamlExample := "whatever: \"value\""
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.yaml": yamlExample,
		},
	}
	manipulator := YamlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: CommonMapManipulator{
			Unmarshaller: YamlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.yaml") {
		t.Fatal("Must be able to manipulate YAML files")
	}

	err := manipulator.SetValue("/etc/config.yaml", "whatever:whatever2", "newvalue")

	if err == nil {
		t.Fatal("Should have failed to perform replacement")
	}
}

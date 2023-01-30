package tomlmanipulators

import (
	"fmt"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/manipulators"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/readers"
	"github.com/mcasperson/UltimateDockerLauncher/cmd/internal/writers"
	"github.com/pelletier/go-toml/v2"
	"testing"
)

func TestTomlInvalidFile(t *testing.T) {
	tamlExample := "whatever= \"value\""
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if manipulator.CanManipulate("/etc/config.doesnotexist") {
		t.Fatal("This should have failed")
	}
}

func TestTomlInvalidToml(t *testing.T) {
	tamlExample := "blah: hi\n- hi"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}
}

func TestTomlSetInvalidFile(t *testing.T) {
	tamlExample := "whatever= \"value\""
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	err := manipulator.SetValue("/etc/config.doesnotexist", "whatever", "newvalue")

	if err == nil {
		t.Fatal("This should have failed")
	}
}

func TestTomlSetInvalidToml(t *testing.T) {
	tamlExample := "blah: hi\n- hi"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever", "newvalue")

	if err == nil {
		t.Fatal("This should have failed")
	}
}

func TestTomlSetStringField(t *testing.T) {
	tamlExample := "whatever= \"value\""
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever", "newvalue")

	if err != nil {
		t.Fatal("Failed to manipulate TOML file: " + err.Error())
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)

	value, ok := result["whatever"].(string)

	if !ok {
		t.Fatal("Value must be a string")
	}

	if value != "newvalue" {
		t.Fatal("Value must be set to \"newvalue\" (was: \"" + value + "\"")
	}
}

func TestTomlSetNumberField(t *testing.T) {
	tamlExample := "\"whatever\"= 5"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever", "6")

	if err != nil {
		t.Fatal("Failed to manipulate TOML file")
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)

	value, ok := result["whatever"].(float64)

	if !ok {
		t.Fatal("Value must be a float")
	}

	if value != 6 {
		t.Fatal("Value must be set to \"6\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestTomlSetNumberFieldWithString(t *testing.T) {
	tamlExample := "\"whatever\"= 5"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever", "newvalue")

	if err != nil {
		t.Fatal("Failed to manipulate TOML file")
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)

	value, ok := result["whatever"].(string)

	if !ok {
		t.Fatal("Value must be a string")
	}

	if value != "newvalue" {
		t.Fatal("Value must be set to \"newvalue\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestTomlSetBoolField(t *testing.T) {
	tamlExample := "\"whatever\"= true"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever", "false")

	if err != nil {
		t.Fatal("Failed to manipulate TOML file")
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)

	value, ok := result["whatever"].(bool)

	if !ok {
		t.Fatal("Value must be a bool")
	}

	if value {
		t.Fatal("Value must be set to \"false\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestTomlSetBoolFieldWithString(t *testing.T) {
	tamlExample := "\"whatever\"= true"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever", "newvalue")

	if err != nil {
		t.Fatal("Failed to manipulate TOML file")
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)

	value, ok := result["whatever"].(string)

	if !ok {
		t.Fatal("Value must be a string")
	}

	if value != "newvalue" {
		t.Fatal("Value must be set to \"newvalue\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestTomlSetObjectField(t *testing.T) {
	tamlExample := "[whatever]\nwhatever2 = true"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever", "whatever3 = 6")

	if err != nil {
		t.Fatal("Failed to manipulate TOML file")
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)

	value, ok := result["whatever"].(map[string]any)

	if !ok {
		t.Fatal("Value must be a map")
	}

	value2, ok := value["whatever3"].(int64)

	if !ok {
		t.Fatal("Nested Value must be a int")
	}

	if value2 != 6 {
		t.Fatal("Nested value must be set to \"6\" (was: \"" + fmt.Sprint(value2) + "\"")
	}
}

func TestTomlSetObjectFieldWithString(t *testing.T) {
	tamlExample := "[whatever]\nwhatever2= true"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever", "newvalue")

	if err != nil {
		t.Fatal("Failed to manipulate TOML file")
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)

	value, ok := result["whatever"].(string)

	if !ok {
		t.Fatal("Value must be a string")
	}

	if value != "newvalue" {
		t.Fatal("Nested value must be set to \"newvalue\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestTomlSetArrayField(t *testing.T) {
	tamlExample := "whatever= [\"hi\"]"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever", "[\"there\"]")

	if err != nil {
		t.Fatal("Failed to manipulate TOML file")
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)

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

func TestTomlSetArrayFieldWithString(t *testing.T) {
	tamlExample := "whatever= [\"hi\"]"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever", "there")

	if err != nil {
		t.Fatal("Failed to manipulate TOML file")
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)

	value, ok := result["whatever"].(string)

	if !ok {
		t.Fatal("Value must be an string")
	}

	if value != "there" {
		t.Fatal("Nested value must be set to \"there\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestTomlSetArrayFieldWithArray(t *testing.T) {
	tamlExample := "whatever = [[\"hi\"]]"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever:0", "[\"there\"]")

	if err != nil {
		t.Fatal("Failed to manipulate TOML file")
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)

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

func TestTomlSetNewField(t *testing.T) {
	tamlExample := "whatever= \"value\""
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever2", "newvalue")

	if err != nil {
		t.Fatal("Failed to perform replacement")
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)
	value, ok := result["whatever2"].(string)

	if !ok {
		t.Fatal("Value must be a string")
	}

	if value != "newvalue" {
		t.Fatal("New value must be set to \"newvalue\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestTomlSetNewNumberField(t *testing.T) {
	tamlExample := "whatever= \"value\""
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever2", "5")

	if err != nil {
		t.Fatal("Failed to perform replacement")
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)
	value, ok := result["whatever2"].(string)

	if !ok {
		t.Fatal("Value must be a string")
	}

	if value != "5" {
		t.Fatal("New value must be set to \"5\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestTomlSetNewBooleanField(t *testing.T) {
	tamlExample := "whatever= \"value\""
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever2", "true")

	if err != nil {
		t.Fatal("Failed to perform replacement")
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)
	value, ok := result["whatever2"].(string)

	if !ok {
		t.Fatal("Value must be a string")
	}

	if value != "true" {
		t.Fatal("New value must be set to \"true\" (was: \"" + fmt.Sprint(value) + "\"")
	}
}

func TestTomlSetArrayFieldIndex(t *testing.T) {
	tamlExample := "whatever = [\"hi\"]"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever:0", "there")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)

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

func TestTomlSetNumberArrayFieldIndexWithNumber(t *testing.T) {
	tamlExample := "whatever = [10]"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever:0", "20")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)

	value, ok := result["whatever"].([]any)

	if !ok {
		t.Fatal("Value must be an array")
	}

	value2, ok := value[0].(float64)

	if !ok {
		t.Fatal("Nested Value must be a int")
	}

	if value2 != 20 {
		t.Fatal("Nested value must be set to 20 (was: \"" + fmt.Sprint(value2) + "\"")
	}
}

func TestTomlSetArrayFieldIndexOutOfBounds(t *testing.T) {
	tamlExample := "whatever= [\"hi\"]"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever:10", "there")

	if err == nil {
		t.Fatal("This should have failed")
	}
}

func TestTomlSetArrayFieldAgainstObject(t *testing.T) {
	tamlExample := "[whatever]\nhi= \"there\""
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever:10", "there")

	if err == nil {
		t.Fatal("This should have failed")
	}
}

func TestTomlSetArrayFieldDoubleIndex(t *testing.T) {
	tamlExample := "whatever= [\"hi\"]"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever:0:0", "there")

	if err == nil {
		t.Fatal("This should have failed")
	}
}

func TestTomlSetArrayFieldIndexWithNumber(t *testing.T) {
	tamlExample := "whatever = [\"hi\"]"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever:0", "10")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)

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

func TestTomlSetArrayFieldIndexWithBool(t *testing.T) {
	tamlExample := "whatever = [\"hi\"]"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever:0", "true")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)

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

func TestTomlSetIntArrayFieldIndexWithInt(t *testing.T) {
	tamlExample := "whatever = [10]"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever:0", "20")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)

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

func TestTomlSetIntArrayFieldIndexWithString(t *testing.T) {
	tamlExample := "whatever = [10]"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever:0", "blah")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)

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

func TestTomlSetBoolArrayFieldIndexWithBool(t *testing.T) {
	tamlExample := "whatever = [true]"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever:0", "false")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)

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

func TestTomlSetBoolArrayFieldIndexWithString(t *testing.T) {
	tamlExample := "whatever = [true]"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever:0", "blah")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)

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

func TestTomlSetObjectArrayFieldIndexWithObject(t *testing.T) {
	tamlExample := "[[whatever]]\nwhatever2= \"hi\""
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever:0", "{\"whatever3\":\"there\"}")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)

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

func TestTomlSetObjectArrayFieldIndexWithString(t *testing.T) {
	tamlExample := "[[whatever]]\nwhatever2=\"hi\""
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever:0", "there")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)

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

func TestTomlSetArrayArrayFieldIndexWithArray(t *testing.T) {
	tamlExample := "whatever = [[\"whatever2\",\"hi\"]]"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever:0", "[\"whatever3\",\"there\"]")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)

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

func TestTomlSetArrayArrayFieldIndexWithString(t *testing.T) {
	tamlExample := "whatever = [[\"whatever2\",\"hi\"]]"
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever:0", "there")

	if err != nil {
		t.Fatal("Failed to manipulate JSON file: " + err.Error())
	}

	var result map[string]any
	err = toml.Unmarshal([]byte(writer.Output["/etc/config.toml"]), &result)

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

func TestTomlSetMissingNestedField(t *testing.T) {
	tamlExample := "whatever= \"value\""
	writer := writers.StringWriter{}
	reader := readers.StringReader{
		Files: map[string]string{
			"/etc/config.toml": tamlExample,
		},
	}
	manipulator := TomlManipulator{
		Writer: &writer,
		Reader: reader,
		MapManipulator: manipulators.CommonMapManipulator{
			Unmarshaller: TomlUnmarshaller{},
		},
	}

	if !manipulator.CanManipulate("/etc/config.toml") {
		t.Fatal("Must be able to manipulate TOML files")
	}

	err := manipulator.SetValue("/etc/config.toml", "whatever:whatever2", "newvalue")

	if err == nil {
		t.Fatal("Should have failed to perform replacement")
	}
}

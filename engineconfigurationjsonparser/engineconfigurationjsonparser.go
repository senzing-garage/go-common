/*
Package engineconfigurationjsonparser is used to generate the JSON document used to configure a Senzing client.
*/
package engineconfigurationjsonparser

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// EngineConfigurationJsonParserImpl is the default implementation of the EngineConfigurationJsonParser interface.
type EngineConfigurationJsonParserImpl struct {
	EngineConfigurationJson string
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func contains(haystack []string, needle string) bool {
	for _, value := range haystack {
		if value == needle {
			return true
		}
	}
	return false
}

func isJson(unknownString string) bool {
	unknownStringUnescaped, err := strconv.Unquote(unknownString)
	if err != nil {
		unknownStringUnescaped = unknownString
	}
	var jsonString json.RawMessage
	return json.Unmarshal([]byte(unknownStringUnescaped), &jsonString) == nil
}

// ----------------------------------------------------------------------------
// Constructor  methods
// ----------------------------------------------------------------------------

func New(engineConfigurationJson string) (EngineConfigurationJsonParser, error) {
	var err error = nil
	if !isJson(engineConfigurationJson) {
		return nil, fmt.Errorf("incorrect JSON syntax in %s", engineConfigurationJson)
	}
	result := &EngineConfigurationJsonParserImpl{
		EngineConfigurationJson: engineConfigurationJson,
	}
	return result, err
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The GetConfigPath method returns the PIPELINE.CONFIGPATH value of _ENGINE_CONFIGURATION_JSON.

Input
  - ctx: A context to control lifecycle.

Output
  - A string containing the value of a PIPELINE.CONFIGPATH.
*/
func (parser *EngineConfigurationJsonParserImpl) GetConfigPath(ctx context.Context) (string, error) {
	engineConfiguration := &EngineConfiguration{}

	err := json.Unmarshal([]byte(parser.EngineConfigurationJson), &engineConfiguration)
	if err != nil {
		return "", err
	}
	return engineConfiguration.Pipeline.ConfigPath, err
}

/*
The GetConfigPath method returns the PIPELINE.CONFIGPATH value of _ENGINE_CONFIGURATION_JSON.

Input
  - ctx: A context to control lifecycle.

Output
  - A string containing the value of a PIPELINE.CONFIGPATH.
*/
func (parser *EngineConfigurationJsonParserImpl) GetDatabaseUrls(ctx context.Context) ([]string, error) {
	var result []string

	engineConfiguration := &EngineConfiguration{}
	err := json.Unmarshal([]byte(parser.EngineConfigurationJson), &engineConfiguration)
	if err != nil {
		return result, err
	}
	result = append(result, engineConfiguration.Sql.Connection)

	// Handle multi-database case.

	backend := engineConfiguration.Sql.Backend
	if len(backend) > 0 && backend != "SQL" {
		var dictionary map[string]interface{}
		var databaseJsonKeys []string
		err = json.Unmarshal([]byte(parser.EngineConfigurationJson), &dictionary)
		if err != nil {
			return result, err
		}

		// Determine JSON keys for database definitions.

		backendMap := dictionary[backend]
		for _, value := range backendMap.(map[string]interface{}) {
			valueString := value.(string)
			if !contains(databaseJsonKeys, valueString) {
				databaseJsonKeys = append(databaseJsonKeys, valueString)
			}
		}

		// Add each database.

		for _, databaseJsonKey := range databaseJsonKeys {
			databaseJson := dictionary[databaseJsonKey].(map[string]interface{})
			databaseName := databaseJson["DB_1"].(string)
			if !contains(result, databaseName) {
				result = append(result, databaseName)
			}
		}
	}

	// TODO:  Implement multi-database list.

	return result, err
}

/*
The GetResourcePath method returns the PIPELINE.RESOURCEPATH value of _ENGINE_CONFIGURATION_JSON.

Input
  - ctx: A context to control lifecycle.

Output
  - A string containing the value of a PIPELINE.RESOURCEPATH.
*/
func (parser *EngineConfigurationJsonParserImpl) GetResourcePath(ctx context.Context) (string, error) {
	engineConfiguration := &EngineConfiguration{}
	err := json.Unmarshal([]byte(parser.EngineConfigurationJson), &engineConfiguration)
	if err != nil {
		return "", err
	}
	return engineConfiguration.Pipeline.ResourcePath, err
}

/*
The GetSupportPath method returns the PIPELINE.SUPPORTPATH value of _ENGINE_CONFIGURATION_JSON.

Input
  - ctx: A context to control lifecycle.

Output
  - A string containing the value of a PIPELINE.SUPPORTPATH.
*/
func (parser *EngineConfigurationJsonParserImpl) GetSupportPath(ctx context.Context) (string, error) {
	engineConfiguration := &EngineConfiguration{}
	err := json.Unmarshal([]byte(parser.EngineConfigurationJson), &engineConfiguration)
	if err != nil {
		return "", err
	}
	return engineConfiguration.Pipeline.SupportPath, err
}

package api

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"toyHome/database"
	"toyHome/entities"
)

var RequestPayload entities.RequestPayload
var Payload entities.Payload

func TestResponseDiscovery(t *testing.T) {
	database.InitForTest()
	Payload.AccessToken = "test"
	rawMesg, _ := json.Marshal(Payload)
	response, err := ResponseDiscovery(rawMesg)
	assert.Nil(t, err)
	assert.NotEqual(t, len(response), 0)
	t.Logf("%v\n", string(response))
}

func TestResponseRequest(t *testing.T) {
	database.InitForTest()
	Payload.AccessToken = "test"
	rawMesg, _ := json.Marshal(Payload)
	response, err := ResponseRequest(rawMesg)
	assert.Nil(t, err)
	assert.NotEqual(t, len(response), 0)
	t.Logf("%v\n", string(response))
}

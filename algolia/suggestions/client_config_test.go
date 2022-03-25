package suggestions

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig_Unmarshal(t *testing.T) {
	var c IndexConfiguration
	data := []byte(`{"languages":["ja"]}`)
	err := json.Unmarshal(data, &c)
	require.NoError(t, err)
	require.False(t, c.Languages.IsBool)
	require.Equal(t, []string{"ja"}, c.Languages.StringArray)

	c = IndexConfiguration{}
	data = []byte(`{"languages":false}`)
	err = json.Unmarshal(data, &c)
	require.NoError(t, err)
	require.False(t, c.Languages.Bool)

	c = IndexConfiguration{}
	data = []byte(`{}`)
	err = json.Unmarshal(data, &c)
	require.NoError(t, err)
	require.False(t, c.Languages.IsBool)
}

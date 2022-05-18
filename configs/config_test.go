package configs

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	v := viper.New()
	t.Run("Load does not return error", func(t *testing.T) {
		// wantErr := error.Error(viper.ConfigFileNotFoundError{})
		v.AddConfigPath("/etc/viper")
		err := v.ReadInConfig()
		assert.NoError(t, err)
	})

}

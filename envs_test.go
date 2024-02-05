package envs

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvs(t *testing.T) {

	os.Setenv("FLAG_INCORRECT_CASING", "test")

	flag_incorrect_casing := flag.String("flag_incorrect_casing", "default", "")

	flag.Parse()

	err := GetAllFlags()
	assert.NoError(t, err)

	assert.Equal(t, "test", *flag_incorrect_casing)

}

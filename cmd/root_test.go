package cmd

import (
	"os/user"
	"testing"

	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestCanMvFile(t *testing.T) {
	defer viper.Reset()
	appFs = afero.NewMemMapFs()
	appFs.Mkdir("testingblah", 0755)
	moveFile()
	exists, _ := afero.DirExists(appFs, "testingnot")
	assert.Equal(t, exists, true)
}

func TestCanReadConfig(t *testing.T) {
	defer viper.Reset()
	appFs = afero.NewMemMapFs()
	usr, err := user.Current()
	if err != nil {
		t.Error()
	}
	var yamlExample = []byte(`
movies:
  exts:
  - mkv
  - avi
  patterns:
  - /[sS]([0-9]+)[eE]([0-9]+).+(720p|1080p)?/
  result: S$1E$2
`)
	afero.WriteFile(appFs, usr.HomeDir+"/.massnomer.yaml", yamlExample, 0755)
	initConfig()
	assert.Equal(t, viper.GetStringSlice("movies.exts")[0], "mkv")
	assert.Equal(t, viper.GetString("movies.result"), "S$1E$2")
}

func TestHasDefaultProfiles(t *testing.T) {
	defer viper.Reset()

	initConfig()
	assert.Equal(t, viper.GetStringSlice("shows.exts")[0], "mkv")
	assert.Equal(t, viper.GetStringSlice("shows.patterns")[0], "/[sS]([0-9]+)[eE]([0-9]+).+(720p|1080p)?/")
}

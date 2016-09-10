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
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)
	afero.WriteFile(appFs, usr.HomeDir+"/.massnomer.yaml", yamlExample, 0755)
	initConfig()
	viper.Get("name")
	assert.Equal(t, viper.Get("name"), "steve")
	assert.Equal(t, viper.GetStringSlice("hobbies")[0], "skateboarding")
}

package cmd

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestCanMvFile(t *testing.T) {
	defer viper.Reset()
	appFs = afero.NewMemMapFs()
	// afs := &Afero{Fs: appFs}
	appFs.Mkdir("testingblah", 0755)
	moveFile()
	exists, _ := afero.DirExists(appFs, "testingnot")
	assert.Equal(t, exists, true)
	//	viper.Set("")
}

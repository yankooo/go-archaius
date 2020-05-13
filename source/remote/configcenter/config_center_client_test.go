package configcenter_test

import (
	"github.com/yankooo/go-archaius/source/remote"
	"github.com/yankooo/go-archaius/source/remote/configcenter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfigCenter(t *testing.T) {
	c, err := configcenter.NewConfigCenter(remote.Options{
		ServerURI: "http://",
		Labels:    map[string]string{remote.LabelApp: "default"}})
	assert.NoError(t, err)
	assert.Equal(t, "default", c.Options().Labels[remote.LabelApp])
}

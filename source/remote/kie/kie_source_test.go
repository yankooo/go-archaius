package kie

import (
	"testing"

	"github.com/yankooo/go-archaius"
	"github.com/yankooo/go-archaius/source/remote"
	"github.com/stretchr/testify/assert"
)

func TestNewKieSource(t *testing.T) {
	opts := &archaius.RemoteInfo{
		DefaultDimension: map[string]string{
			remote.LabelApp:     "default",
			remote.LabelService: "cart",
		},
		TenantName: "default",
		URL:        "http://",
	}
	_, err := NewKieSource(opts)
	assert.NoError(t, err)
}

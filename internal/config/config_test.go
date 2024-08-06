package config

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	fmt.Println(Cfg.Oss.Endpoint)
	if Cfg.Oss.Endpoint == "" {
		t.Errorf("endpoint is empty")
		t.Failed()
		return
	}
	t.Logf("endpoint: %s", Cfg.Oss.Endpoint)
}

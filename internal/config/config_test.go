package config

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	fmt.Println(Cfg.Oss.Supplier)
	fmt.Println(Cfg.Oss.AccessKey)
}

package config

import (
	"github.com/Dizzrt/DAuth/backend/common/config/config_base"
	"github.com/Dizzrt/DAuth/backend/common/dlog"
)

const (
	_KEY_SERVER_PORT = "server.port"
)

func ServerGetPort() uint16 {
	cfg := config_base.Instance()

	var port uint16
	if !cfg.V.IsSet(_KEY_SERVER_PORT) {
		port = 80

		cfg.V.Set(_KEY_SERVER_PORT, port)
		err := cfg.V.WriteConfig()
		if err != nil {
			dlog.Errorf("set default server port failed with error: %v", err)
		}
	} else {
		port = cfg.V.GetUint16(_KEY_SERVER_PORT)
	}

	return port
}

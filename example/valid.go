package example

import (
	"go.uber.org/zap"
)

type validVoid struct {
	smth string
}

func valid() {
	void := validVoid{
		smth: "smth",
	}
	zap.L().Info("info - ", zap.String("void smth", void.smth))
}

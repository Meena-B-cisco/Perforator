package logfield

import (
	"github.com/yandex/perforator/library/go/core/log"
	"github.com/yandex/perforator/perforator/pkg/linux"
)

func Pid(pid linux.ProcessID) log.Field {
	return log.UInt32("pid", uint32(pid))
}

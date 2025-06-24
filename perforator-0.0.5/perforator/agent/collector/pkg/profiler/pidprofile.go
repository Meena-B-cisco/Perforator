package profiler

import (
	"github.com/yandex/perforator/perforator/agent/collector/pkg/machine"
	"github.com/yandex/perforator/perforator/pkg/linux"
)

////////////////////////////////////////////////////////////////////////////////

type trackedProcess struct {
	pid     linux.ProcessID
	builder *multiProfileBuilder
}

func newTrackedProcess(
	pid linux.ProcessID,
	labels map[string]string,
	bpf *machine.BPF,
) (*trackedProcess, error) {
	err := bpf.AddTracedProcess(pid)
	if err != nil {
		return nil, err
	}

	profiler := &trackedProcess{
		pid:     pid,
		builder: newMultiProfileBuilder(labels),
	}
	return profiler, nil
}

////////////////////////////////////////////////////////////////////////////////

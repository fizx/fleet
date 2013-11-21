package job

import (
	"errors"
	"fmt"
	"strings"

	"github.com/coreos/coreinit/machine"
)

type Job struct {
	Name    string      `json:"name"`
	State   *JobState   `json:"state"`
	Payload *JobPayload `json:"payload"`
}

// JobPayload should be specific to the type of target it is
// running on. Currently, coreinit only supports the 'systemd' type
type JobPayload struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func NewJob(name string, state *JobState, payload *JobPayload) *Job {
	return &Job{name, state, payload}
}

func NewJobPayload(payloadType string, value string) (*JobPayload, error) {
	switch payloadType {
		case "systemd-service":
		case "systemd-socket":
			break
		default:
			return nil, errors.New("Invalid JobPayload type argument")
	}

	return &JobPayload{payloadType, value}, nil
}

func NewJobPayloadFromSystemdUnit(name string, contents string) (*JobPayload, error) {
	var payloadType string

	if strings.HasSuffix(name, ".service") {
		payloadType = "systemd-service"
	} else if strings.HasSuffix(name, ".socket") {
		payloadType = "systemd-socket"
	} else {
		return nil, errors.New(fmt.Sprintf("Unrecognized systemd unit %s", name))
	}

	return NewJobPayload(payloadType, contents)
}

// JobState should be generated by the target that is actually running
// the job. This will be extended with other data in the future (i.e. ports,
// memory usage, etc).
type JobState struct {
	State   string           `json:"state"`
	Machine *machine.Machine `json:"machine"`
}

func NewJobState(state string, machine *machine.Machine) *JobState {
	return &JobState{state, machine}
}
package probes

import corev1 "k8s.io/api/core/v1"

type Modification func(*corev1.Probe)

func Apply(funcs ...Modification) Modification {
	return func(probe *corev1.Probe) {
		for _, f := range funcs {
			f(probe)
		}
	}
}

func New(funcs ...Modification) corev1.Probe {
	probe := corev1.Probe{}
	for _, f := range funcs {
		f(&probe)
	}
	return probe
}

func WithExecCommand(cmd []string) Modification {
	return func(probe *corev1.Probe) {
		if probe.Handler.Exec == nil {
			probe.Handler.Exec = &corev1.ExecAction{}
		}
		probe.Handler.Exec.Command = cmd
	}
}

func WithFailureThreshold(failureThreshold int) Modification {
	return func(probe *corev1.Probe) {
		probe.FailureThreshold = int32(failureThreshold)
	}
}

func WithInitialDelaySeconds(initialDelaySeconds int) Modification {
	return func(probe *corev1.Probe) {
		probe.InitialDelaySeconds = int32(initialDelaySeconds)
	}
}
func WithSuccessThreshold(successThreshold int) Modification {
	return func(probe *corev1.Probe) {
		probe.SuccessThreshold = int32(successThreshold)
	}
}
func WithPeriodSeconds(periodSeconds int) Modification {
	return func(probe *corev1.Probe) {
		probe.PeriodSeconds = int32(periodSeconds)
	}
}
func WithTimeoutSeconds(initialDelaySeconds int) Modification {
	return func(probe *corev1.Probe) {
		probe.InitialDelaySeconds = int32(initialDelaySeconds)
	}
}

func WithHandler(handler corev1.Handler) Modification {
	return func(probe *corev1.Probe) {
		probe.Handler = handler
	}
}
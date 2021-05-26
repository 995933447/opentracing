package tracer_config

type TracerConfig interface {
	SetOption(string, interface{}) *TracerConfig
}

package config

const (
	defaultFile = "/etc/mnemonic.conf"
	defaultDir  = "/mnemonic"
)

// Config represents the configuration of the mnemonic tool.
type Config struct {
	file string
	dir  string
}

// New returns a new Config.
func New(file, dir string) *Config {
	cfg := &Config{
		file: file,
		dir:  dir,
	}
	return cfg
}

// DefaultConfig returns the default mnemonic configuration.
func DefaultConfig() *Config {
	return &Config{
		file: defaultFile,
		dir:  defaultDir,
	}
}

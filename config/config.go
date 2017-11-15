package config

// Store provides a configuration store interface. Services can retrieve
// their configuration using the service names.
type Store interface {
	// Open creates the resources like db connections or files required by the store.
	Open() error

	// Close releases any underlying resources used by the store.
	Close()

	// Get returns the configuration for the specified service or error if the
	// configuration is not found in the store.
	Get(name string, config interface{}) error
}

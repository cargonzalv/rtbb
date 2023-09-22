package serializer

// Dependency injection Serializer provider.
func ProvideJsonSerializer() JsonSerializer {
	return &service{}
}

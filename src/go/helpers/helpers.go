package helpers

func WithDefault(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

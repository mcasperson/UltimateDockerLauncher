package customerror

// UdlError captures an error and the env var that triggered it
type UdlError struct {
	EnvVar string
	Err    error
}

func (e *UdlError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}

	return ""
}

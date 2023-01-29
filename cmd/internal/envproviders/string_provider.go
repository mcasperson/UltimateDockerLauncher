package envproviders

type StringProvider struct {
	Vars map[string]string
}

func (e StringProvider) GetEnvVar(name string) string {
	if e.Vars == nil {
		return ""
	}

	value, ok := e.Vars[name]
	if ok {
		return value
	}
	return ""
}

func (e StringProvider) GetAllEnvVars() []string {
	retValue := []string{}

	if e.Vars != nil {
		for k, v := range e.Vars {
			retValue = append(retValue, k+"="+v)
		}
	}

	return retValue
}

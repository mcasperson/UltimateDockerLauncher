package prefixes

// EnvVarPrefixes lists the known prefixes to scan for. For example, Azure web apps have a set of prefixes
// applied to each environment variable.
// https://learn.microsoft.com/en-us/azure/app-service/reference-app-settings?tabs=kudu%2Cdotnet#variable-prefixes
var EnvVarPrefixes = []string{"", "APPSETTING_"}

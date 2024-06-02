package types

type VSCodeSettings map[string]interface{}

type VSCodeExtensions struct {
	Recommendations []string `json:"recommendations"`
}

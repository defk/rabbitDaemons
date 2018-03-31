package structures

type Task struct {
	Alias string `json:"alias"`
	Content struct {
		GUID    string                 `json:"guid"`
		Version string                 `json:"version"`
		Data    map[string]interface{} `json:"data"`
	} `json:"content"`
}

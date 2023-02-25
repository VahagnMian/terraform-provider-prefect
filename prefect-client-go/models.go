package prefect_client

type WorkQueue struct {
	ID               string `json:"id,omitempty"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	IsPaused         bool   `json:"is_paused"`
	ConcurrencyLimit int    `json:"concurrency_limit"`
}

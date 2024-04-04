package query

type JobResponse struct {
	Status    string `json:"pending"`
	JobNumber string `json:"job_number"`
}

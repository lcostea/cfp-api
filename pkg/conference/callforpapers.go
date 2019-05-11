package conference

import "time"

//CallForPapers represents a call for papers/requests done by conference organisers in order to get talks
type CallForPapers struct {
	ConferenceName string
	ConferenceID   string
	URL            string
	Deadline       time.Time
	Description    string
	Starts         time.Time
}

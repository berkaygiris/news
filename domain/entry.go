package domain

type Entry struct {
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
	URL       string `json:"url"`
	Comment   string `json:"comment"`
	Type      string `json:"type"`
	Date      string `json:"date"`
}

func (e Entry) GroupKey() string {
	if len(e.Date) >= 7 {
		return e.Date[:7]
	}
	return e.Date
}

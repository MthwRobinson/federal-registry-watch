package models

type Result struct {
	Title                  string `json:"title"`
	Type                   string `json:"type"`
	Abstract               string `json:"abstract"`
	DocumentNumber         string `json:"document_number"`
	HTMLURL                string `json:"html_url"`
	PdfURL                 string `json:"pdf_url"`
	PublicInspectionPdfURL string `json:"public_inspection_pdf_url"`
	PublicationDate        string `json:"publication_date"`
	Agencies               []struct {
		RawName  string      `json:"raw_name"`
		Name     string      `json:"name"`
		ID       int         `json:"id"`
		URL      string      `json:"url"`
		JSONURL  string      `json:"json_url"`
		ParentID interface{} `json:"parent_id"`
		Slug     string      `json:"slug"`
	} `json:"agencies"`
	Excerpts string `json:"excerpts"`
}

type RegisterResults struct {
	Count       int      `json:"count"`
	Description string   `json:"description"`
	TotalPages  int      `json:"total_pages"`
	NextPageURL string   `json:"next_page_url"`
	Results     []Result `json:"results"`
}

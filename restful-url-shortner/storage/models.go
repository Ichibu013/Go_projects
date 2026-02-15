package storage

type URLMapping struct {
	OriginalURL string `json:"original_url"`
	ShortCode   string `json:": short_code"`
}

// In - memory Database
// Key = ShortCode, Value = OriginalURL
var urlStore = make(map[string]string)

// @Setter of SpringBoot
func (r *Repository) SaveURL(shortCode, originalURL string) error {
	query := `INSERT INTO urls (shortCode, originalURL) VALUES ($1, $2)`
	_, err := r.db.Exec(query, shortCode, originalURL)
	return err
}

func (r *Repository) GetURL(shortCode string) (string, bool) {
	var originalURL string
	query := `SELECT original_url FROM urls WHERE short_code = $1`
	err := r.db.QueryRow(query, shortCode).Scan(&originalURL)
	if err != nil {
		return "", false
	}
	return originalURL, true
}

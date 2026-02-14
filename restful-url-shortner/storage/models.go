package storage

type URLMapping struct {
	OriginalURL string `json:"original_url"`
	ShortCode   string `json:": short_code"`
}

// In - memory Database
// Key = ShortCode, Value = OriginalURL
var urlStore = make(map[string]string)

// @Setter of SpringBoot
func SaveUrlMapping(shortCode, originalUrl string) {
	urlStore[shortCode] = originalUrl
}

// @Getter of Springboot
func GetUrlMapping(shortCode string) (string, bool) {
	url, found := urlStore[shortCode]
	return url, found
}

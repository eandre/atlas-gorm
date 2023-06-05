package url

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	"encore.dev/storage/sqldb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type URL struct {
	ID         string // short-form URL id
	URL        string // complete URL, in long form
	VisitCount int    // number of times visited
}

type ShortenParams struct {
	URL string // the URL to shorten
}

// Shorten shortens a URL.
//
//encore:api public method=POST path=/url
func (s *Service) Shorten(ctx context.Context, p *ShortenParams) (*URL, error) {
	id, err := generateID()
	if err != nil {
		return nil, err
	}
	url := &URL{ID: id, URL: p.URL}
	if err := s.db.Create(url).Error; err != nil {
		return nil, err
	}
	return url, nil
}

// Get retrieves the original URL for the id.
//
//encore:api public method=GET path=/url/:id
func (s *Service) Get(ctx context.Context, id string) (*URL, error) {
	var url URL
	err := s.db.
		Model(&url).
		Where("id = ?", id).
		Update("visit_count", gorm.Expr("visit_count + 1")).
		First(&url).
		Error
	return &url, err
}

type ListResponse struct {
	URLs []*URL
}

// List retrieves all URLs.
//
//encore:api public method=GET path=/url
func (s *Service) List(ctx context.Context) (*ListResponse, error) {
	var urls []*URL
	if err := s.db.Find(&urls).Error; err != nil {
		return nil, err
	}
	return &ListResponse{URLs: urls}, nil
}

// generateID generates a random short ID.
func generateID() (string, error) {
	var data [6]byte // 6 bytes of entropy
	if _, err := rand.Read(data[:]); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(data[:]), nil
}

//encore:service
type Service struct {
	db *gorm.DB
}

// initService initializes the site service.
// It is automatically called by Encore on service startup.
//
//lint:ignore U1000 called by Encore
func initService() (*Service, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: urldb.Stdlib(),
	}))
	if err != nil {
		return nil, err
	}
	return &Service{db: db}, nil
}

var urldb = sqldb.NewDatabase("url", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})

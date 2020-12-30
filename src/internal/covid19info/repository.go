package covid19info

import "log"

type Repository interface {
	// Get returns the album with the specified album ID.
	/*
		Get(ctx context.Context, id string) (entity.Album, error)
		// Count returns the number of albums.
		Count(ctx context.Context) (int, error)
		// Query returns the list of albums with the given offset and limit.
		Query(ctx context.Context, offset, limit int) ([]entity.Album, error)
		// Create saves a new album in the storage.
		Create(ctx context.Context, album entity.Album) error
		// Update updates the album with given ID in the storage.
		Update(ctx context.Context, album entity.Album) error
		// Delete removes the album with given ID from the storage.
		Delete(ctx context.Context, id string) error
	*/
}

// repository persists albums in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new album repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

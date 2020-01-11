package auth

// TokenFetcher fetches tokens from an authentication entity.
type TokenFetcher interface {
	Fetch() (string, error)
}

func Fetch(fetcher TokenFetcher) (string, error) {
	return fetcher.Fetch()
}

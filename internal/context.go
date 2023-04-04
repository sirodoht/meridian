package internal

type ContextKey int

const (
	KeyIdentity        ContextKey = iota
	KeyIsAuthenticated ContextKey = iota
)

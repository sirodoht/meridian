package internal

type ContextKey int

const (
	KeyUsername        ContextKey = iota
	KeyIsAuthenticated ContextKey = iota
)

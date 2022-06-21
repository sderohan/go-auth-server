package jwtauth

type IJWTAuth interface {
	GenerateToken()
	ValidateToken()
}

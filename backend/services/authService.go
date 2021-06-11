package auth

type AuthService interface {
  GenerateToken(userId int) (string, error)
  TokenValidation(encodedToken string) (*jwt.Token, error)
}
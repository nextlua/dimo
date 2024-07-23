package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"log/slog"
	"nextlua/dimo/internal/config"
	"nextlua/dimo/internal/models"
	"nextlua/dimo/internal/repository"
)

// compile-time proofs of service interface implementation
var _ Service = (*service)(nil)

// Service defines behaviors of sample service
type Service interface {
	Auth(ctx *gin.Context, req models.AuthInput) (int, models.AuthResponse)
	AuthRedirect(ctx *gin.Context, req models.AuthRedirectInput) (int, models.AuthRedirectResponse)
	Token(ctx *gin.Context, req models.TokenInput) (int, models.TokenResponse)
}

// Service represents service
type service struct {
	logger       *slog.Logger
	repo         *repository.Repository
	cacheManager *cache.Cache
	envVariables *config.EnvVars
}

// NewService creates and returns service
func NewService(
	logger *slog.Logger,
	repo *repository.Repository,
	cacheManager *cache.Cache,
	envVariables *config.EnvVars) Service {
	return &service{
		logger:       logger,
		repo:         repo,
		cacheManager: cacheManager,
		envVariables: envVariables,
	}
}

func (s *service) log(ctx context.Context, level slog.Level, message string, additionalParams map[string]interface{}) {
	attrs := make([]interface{}, 0, len(additionalParams))
	for key, value := range additionalParams {
		attrs = append(attrs, key, value)
	}

	s.logger.Log(ctx, level, message, attrs...)
}

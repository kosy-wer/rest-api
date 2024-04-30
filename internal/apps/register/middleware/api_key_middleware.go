package middleware
                                                                      
import (
        "net/http"
        "rest_api/internal/apps/register/helper"
        "rest_api/internal/apps/register/model/web"
)

type ApiKeyMiddleware struct {
    cfg            conf.Config
    logger         logging.Logger
    apiKeyHeader   string
    decodedAPIKeys map[string][]byte
}

func NewApiKeyMiddleware(cfg conf.Config, logger logging.Logger) (*ApiKeyMiddleware, error) {
    apiKeyHeader := cfg.APIKeyHeader
    apiKeys := cfg.APIKeys

    decodedAPIKeys := make(map[string][]byte)
    for name, value := range apiKeys {
        decodedKey, err := hex.DecodeString(value)
        if err != nil {
            return nil, err
        }
        decodedAPIKeys[name] = decodedKey
    }

    return &ApiKeyMiddleware{
        cfg:            cfg,
        logger:         logger,
        apiKeyHeader:   apiKeyHeader,
        decodedAPIKeys: decodedAPIKeys,
    }, nil
}

func (mw *ApiKeyMiddleware) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()

        apiKey, err := bearerToken(r, mw.apiKeyHeader)
        if err != nil {
            mw.logger.Errorw("request failed API key authentication", "error", err)
            RespondError(w, http.StatusUnauthorized, "invalid API key")
            return
        }

        if _, ok := apiKeyIsValid(apiKey, mw.decodedAPIKeys); !ok {
            hostIP, _, err := net.SplitHostPort(r.RemoteAddr)
            if err != nil {
                mw.logger.Errorw("failed to parse remote address", "error", err)
                hostIP = r.RemoteAddr
            }
            mw.logger.Errorw("no matching API key found", "remoteIP", hostIP)

            RespondError(w, http.StatusUnauthorized, "invalid api key")
            return
        }

        next.ServeHTTP(w, r.WithContext(ctx))
    })
}


package util

import (
	"strconv"

	"github.com/alinz/releasifier/web/constants"
	"github.com/dgrijalva/jwt-go"
	"github.com/pressly/chi"
	"golang.org/x/net/context"
)

//GetUserIDFromContext returns user id from context
func GetUserIDFromContext(ctx context.Context) (int64, error) {
	token := ctx.Value(constants.CtxKeyJwtToken).(*jwt.Token)
	userIDStr := token.Claims["user_id"].(string)
	userID, err := strconv.ParseInt(userIDStr, 10, 64)

	if err != nil {
		return -1, err
	}

	return userID, nil
}

//GetParamValueAsID accepts the context and try to parse an id into int64
func GetParamValueAsID(ctx context.Context, param string) (int64, error) {
	value := chi.URLParams(ctx)[param]
	id, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		return -1, err
	}

	return id, nil
}

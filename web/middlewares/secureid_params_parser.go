package middlewares

import (
	"fmt"
	"net/http"

	"github.com/alinz/releasifier/data"
	"github.com/alinz/releasifier/lib/utils"
	"github.com/pressly/chi"
	"golang.org/x/net/context"
)

//SecureIDParamsParser accepts list of params and decrypt them as a valid id.
//it returns a http error 400 if one of the id is not correctly formated.
func SecureIDParamsParser(names ...string) func(chi.Handler) chi.Handler {
	return func(next chi.Handler) chi.Handler {
		return chi.HandlerFunc(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
			params := ctx.Value(chi.URLParamsCtxKey).(map[string]string)

			for _, key := range names {
				secureID, err := data.DecryptSecureID(params[key])
				if err != nil {
					utils.Respond(w, 400, fmt.Errorf("param '%s' is not a valid id", key))
					return
				}
				params[key] = secureID.String()
			}

			next.ServeHTTPC(ctx, w, r)
		})
	}
}

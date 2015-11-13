package app

import (
	"net/http"

	"github.com/alinz/releasifier/lib/utils"
	"golang.org/x/net/context"
)

func test(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	utils.Respond(w, 200, "wow")
}

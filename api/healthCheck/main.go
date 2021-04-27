package healthCheck

import (
	"net/http"

	"secretSanta/utils"
)

var GetHealth = func(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, http.StatusOK, "Ah Ah Ah Ah Stayin Alive!!", nil)
}

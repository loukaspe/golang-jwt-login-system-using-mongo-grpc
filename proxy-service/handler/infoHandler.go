package handler

import (
	"encoding/json"
	"github.com/loukaspe/auth/proxy/helper"
	"github.com/sirupsen/logrus"
	"net/http"
)

type InfoHandler struct {
	claimsRetriever helper.ContextJwtClaimsRetrieverInterface
}

func NewInfoHandler(
	claimsRetriever helper.ContextJwtClaimsRetrieverInterface,
) *InfoHandler {
	return &InfoHandler{claimsRetriever: claimsRetriever}
}

func (handler *InfoHandler) InfoController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jwtClaims := r.Context().Value(r.Header.Get("Authorization"))
	info, err := handler.claimsRetriever.ExtractInfo(jwtClaims)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"requestToken": r.Header.Get("Authorization"),
			"errorMessage": err.Error(),
		}).Error("Error getting jwt claims")

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("bad jwt token")
		return
	}

	w.Write([]byte(info))
	w.WriteHeader(http.StatusOK)
}

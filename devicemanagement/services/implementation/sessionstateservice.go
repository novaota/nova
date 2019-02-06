package implementation

import (
	"encoding/gob"
	"net/http"

	"nova/devicemanagement/services"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
)

// constants
var AuthenticationCookieName = "novaauth"
var UserCookieNode = "user"

type sessionStateService struct {
	cookieStore *sessions.CookieStore
}

func NewSessionStateService() *sessionStateService {
	gob.Register(&services.UserToken{})
	return &sessionStateService{cookieStore: generateCookieStore()}
}

// https://curtisvermeeren.github.io/2018/05/13/Golang-Gorilla-Sessions
func generateCookieStore() *sessions.CookieStore {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	store := sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	store.Options = &sessions.Options{
		MaxAge:   60 * 15,
		HttpOnly: true,
	}

	return store
}
func (service *sessionStateService) RestoreFromRequest(r *http.Request) (*services.UserToken, error) {
	session, err := service.cookieStore.Get(r, AuthenticationCookieName)
	if err != nil {
		return nil, err
	}

	val := session.Values[UserCookieNode]

	token := &services.UserToken{}
	var ok bool
	token, ok = val.(*services.UserToken)

	if !ok {
		return nil, errors.New("UserToken is not present")
	}

	return token, nil
}

func (service *sessionStateService) SetToResponse(r *http.Request, w *http.ResponseWriter, token services.UserToken) error {

	// Get a session. Get() always returns a session, even if empty.
	session, err := service.cookieStore.Get(r, AuthenticationCookieName)
	if err != nil {
		return err
	}

	session.Values[UserCookieNode] = token
	session.Save(r, *w)
	return nil
}

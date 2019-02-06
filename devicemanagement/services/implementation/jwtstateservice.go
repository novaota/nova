package implementation
//
//import (
//	"crypto/rsa"
//	"io/ioutil"
//	"net/http"
//
//	"errors"
//
//	"github.com/dgrijalva/jwt-go"
//	"updateservice/datastorage"
//	"updateservice/services"
//)
//
//type jwtStateService struct {
//	datastorage.RepositoryFactory
//	verifyKey *rsa.PublicKey
//	signKey   *rsa.PrivateKey
//}
//
//type CertificateSettings struct {
//	Certificate string
//	Key         string
//}
//
//func NewJsonWebTokenService(repositoryFactory datastorage.RepositoryFactory, certSettings CertificateSettings) *jwtStateService {
//
//	result := &jwtStateService{}
//
//	signBytes, err := ioutil.ReadFile(certSettings.Key)
//	if err != nil {
//		panic("...")
//	}
//
//	result.signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
//	if err != nil {
//		panic("...")
//	}
//
//	verifyBytes, err := ioutil.ReadFile(certSettings.Certificate)
//	if err != nil {
//		panic("...")
//	}
//
//	result.verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
//	if err != nil {
//		panic("...")
//	}
//
//	return result
//}
//
//var AuthentificationCookie string = "state"
//
//var signingMethod jwt.SigningMethod = jwt.GetSigningMethod("RS256")
//
//func (service *jwtStateService) RestoreFromRequest(r *http.Request) (*services.UserToken, error) {
//	str := r.Header.Get(AuthentificationCookie)
//
//	if len(str) == 0 {
//		return nil, errors.New("No Authentification Header found")
//	}
//
//	authCookie, err := r.Cookie(AuthentificationCookie)
//
//	if err != nil {
//		return nil, err
//	}
//
//	tokenString := authCookie.Value
//
//	result := services.UserToken{}
//
//	token, err := jwt.ParseWithClaims(tokenString,// &result, func(token *jwt.Token) (interface{}, error) {
//		// since we only use the one private key to sign the tokens,
//		// we also only use its public counter part to verify
//		return service.verifyKey, nil
//	})
//
//	if err != nil {
//		return nil, err
//	}
//
//	if !token.Valid {
//		return nil, errors.New("Invalid Token")
//	}
//
//	return &result, nil
//}
//
//func (service *jwtStateService) SetToResponse(w *http.ResponseWriter, token services.UserToken) error {
//
//	jwt := jwt.NewWithClaims(signingMethod, token)
//
//	tokenString, err := jwt.SignedString(service.signKey)
//	if err != nil {
//		return err
//	}
//
//	cookie := &http.Cookie{Name: AuthentificationCookie, Value: tokenString}
//	http.SetCookie(*w, cookie)
//
//	return nil
//}
//
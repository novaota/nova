package updatepackage

// Das ist alles nur geklaut, oooo
// https://www.safaribooksonline.com/library/view/security-with-go/9781788627917/370fc9d6-c0cd-4b96-9533-b5cca1518ecd.xhtml
import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"nova/shared"
)

// Cryptographically sign a message= creating a digital signature
// of the original message. Uses SHA-256 hashing.
func signMessage(privateKey *rsa.PrivateKey, message []byte) ([]byte, error) {
	hashed := sha256.Sum256(message)

	signature, err := rsa.SignPKCS1v15(
		rand.Reader,
		privateKey,
		crypto.SHA256,
		hashed[:],
	)
	if err != nil {
		return nil, err
	}

	return signature, nil
}

//TODO: Too long error fixen
// Load the message that will be signed from file
func loadMessageFromFile(messageFilename string) ([]byte, error) {
	fileData, err := ioutil.ReadFile(messageFilename)
	if err != nil {
		return nil, err
	}
	return fileData, nil
}

// Load the RSA private key from a PEM encoded file
func loadPrivateKeyFromPemFile(privateKeyFilename string) (*rsa.PrivateKey, error) {
	// Quick load file to memory
	fileData, err := ioutil.ReadFile(privateKeyFilename)
	if err != nil {
		return nil, err
	}

	// Get the block data from the PEM encoded file
	block, _ := pem.Decode(fileData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("Unable to load a valid private key.")
	}

	// Parse the bytes and put it in to a proper privateKey struct
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.New("Error loading private key.")
	}

	return privateKey, nil
}
// Returns bool whether signature was verified

func verifySignature(
	signature []byte,
	message []byte,
	publicKey *rsa.PublicKey) (bool, error) {

	hashedMessage := sha256.Sum256(message)

	err := rsa.VerifyPKCS1v15(
		publicKey,
		crypto.SHA256,
		hashedMessage[:],
		signature,
	)

	if err != nil {
		return false, err
	}
	return true, nil
}

// Load file to memory
func loadFile(filename string) ([]byte, error) {
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return fileData, nil
}

// Load a public RSA key from a PEM encoded file
func loadPublicKeyFromPemFile(publicKeyFilename string) (*rsa.PublicKey, error) {
	// Quick load file to memory
	fileData, err := ioutil.ReadFile(publicKeyFilename)
	if err != nil {
		return nil, err
	}

	// Get the block data from the PEM encoded file
	block, _ := pem.Decode(fileData)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("Unable to load valid public key.")
	}

	// Parse the bytes and store in a public key format
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.New("Error loading public key. ")
	}

	return publicKey.(*rsa.PublicKey), nil // Cast interface to PublicKey
}

type certificateSignService struct {
	shared.CertificateSettings
}

func NewCertificateSignService(cert string, key string) *certificateSignService {
	return &certificateSignService{
		CertificateSettings: shared.CertificateSettings{
			CAKey:         key,
			CACertificate: cert,
		},
	}
}

func (service *certificateSignService) IsSigned(signature string, file string) (bool, error) {
	// Load all the files from disk
	publicKey, err := loadPublicKeyFromPemFile(service.CACertificate)

	if err != nil {
		return false, err
	}

	decodedSign, err := base64.StdEncoding.DecodeString(signature)

	if err != nil {
		return false, err
	}

	message, err := loadFile(file)

	if err != nil {
		return false, err
	}
	// Verify signature
	isValid, err := verifySignature(decodedSign, message, publicKey)
	return isValid, err
}

func (service *certificateSignService) SignFile(file string) (string, error) {
	message, err := loadMessageFromFile(file)

	if err != nil {
		return "", err
	}

	privateKey, err := loadPrivateKeyFromPemFile(service.CAKey)

	if err != nil {
		return "", err
	}

	signBytes, err := signMessage(privateKey, message)

	if err != nil {
		return "", err
	}

	signB64 := base64.StdEncoding.EncodeToString(signBytes)

	return signB64, nil
}

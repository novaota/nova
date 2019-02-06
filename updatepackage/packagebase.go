package updatepackage

import (
	"fmt"
	"io/ioutil"

	"nova/shared"
)

type PackageBase struct {
	FileChopService     FileChopService
	SignService         SignService
	CompressionService  CompressionService
	CertificateSettings shared.CertificateSettings
	WorkspaceDir        string
}

func NewPackageBase(certSettings shared.CertificateSettings) *PackageBase {
	result := &PackageBase{}
	result.CertificateSettings = certSettings

	tempDir, err := ioutil.TempDir("", "NOVA")

	if err != nil {
		panic(fmt.Sprintf("Could not create dir: %v", tempDir))
	}

	result.SetWorkspace(tempDir)
	result.SetupDefaultServices()
	return result
}

func (pb *PackageBase) SetWorkspace(workspace string) {
	pb.WorkspaceDir = workspace
}

func (pb *PackageBase) SetupDefaultServices() {
	pb.SignService = NewCertificateSignService(pb.CertificateSettings.CACertificate, pb.CertificateSettings.CAKey)
	pb.FileChopService = NewStraightFileChopService()
	pb.CompressionService = NewZipCompressionService()
}


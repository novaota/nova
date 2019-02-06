package updatepackage

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
	"nova/shared"
)

type Unpacker struct {
	*PackageBase
}

func NewUnpacker(certSettings shared.CertificateSettings) *Unpacker {
	return &Unpacker{
		PackageBase: NewPackageBase(certSettings),
	}
}

func (up *Unpacker) Unpack(deployment UpdatePackageDeployment, outputDir string) (*UpdatePackage, error) {
	tempDestFile := outputDir + "pack.pack"
	defer os.Remove(tempDestFile)

	if err := up.joinFiles(deployment, outputDir); err != nil {
		log.Fatalf("Unpacker: Could not unpack %v\n", err.Error())
	}

	files, err := up.decompressFiles(tempDestFile, outputDir)

	if err != nil {
		log.Fatalf("Unpacker: Could not unpack %v\n", err.Error())
	}

	configFile, err := up.getConfig(files)

	if err != nil {
		log.Fatalf("Unpacker: Could not unpack %v\n", err.Error())
	}

	result := &UpdatePackage{}
	err = shared.DefaultSerializer.DeserializeFromFile(configFile, result)
	return result, err
}

func (up *Unpacker) getConfig(files []string) (string, error){
	for _, file := range files {
		if strings.Contains(file, "config.json") {
			return file, nil
		}
	}

	return "", errors.New("No config file")
}

func (up *Unpacker) joinFiles(deployment UpdatePackageDeployment, outputFile string) error {
	// check signatures
	segments := deployment.Segments

	if err := up.checkSegmentSignatures(segments); err != nil {
		return err
	}

	filenames := up.generateSegmentFileNames(segments)
	up.FileChopService.JoinFile(filenames, outputFile)

	isSigned, _ := up.SignService.IsSigned(deployment.Signature, outputFile)
	if !isSigned {
		os.Remove(outputFile)
		return errors.New("Package Signature doesnt Match...")
	}

	return nil
}

func (up *Unpacker) decompressFiles(src string, dest string) ([]string, error){
	return up.CompressionService.Decompress(src, dest)
}

func (up *Unpacker) generateSegmentFileNames(segments []Segment) []string {
	result := []string{}

	for _, segment := range segments {
		result = append(result, segment.Path)
	}

	return result
}

func (up *Unpacker) checkSegmentSignatures(segments []Segment) error {
	for _, segment := range segments {
		isSigned, _ := up.SignService.IsSigned(segment.Signature, segment.Path)
		if !isSigned{
			return errors.New(fmt.Sprintf("Unpacker: Segment %v Signature is invalid", segment.Path))
		}
	}
	return nil
}
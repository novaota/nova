package updatepackage

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"

	"nova/shared"
)

type Packer struct {
	*PackageBase
}

type UpdatePackageWorkspace struct {
	Dir         string
	ResourceDir string
}

func NewPacker(certSettings shared.CertificateSettings) *Packer {
	return &Packer{
		PackageBase: NewPackageBase(certSettings),
	}
}

func (pc *Packer) Pack(pack *UpdatePackage, outputDir string) (*UpdatePackageDeployment, error) {
	workspace, err := pc.createWorkspace(pack)

	if err != nil {
		return nil, errors.New("Could not create workspace, exiting")
	}

	configFile, err := pc.createPackageJson(pack, *workspace)
	if err != nil {
		return nil, errors.New("Could not create config, exiting")
	}

	files := append(pack.Files, configFile)

	compressedFilename, err := pc.compressFiles(files, workspace, pack)

	signature, err := pc.SignService.SignFile(compressedFilename)

	if err != nil {
		return nil, errors.New("Could not compress files")
	}

	chops, err := pc.chopFiles(compressedFilename, outputDir)

	segments := pc.createSegments(chops)

	result := &UpdatePackageDeployment{
		Segments:  segments,
		Signature: signature,
	}

	configPath := outputDir + pack.Name + ".json"

	shared.DefaultSerializer.SerializeToFile(result, configPath)
	return result, nil
}

func (pc *Packer) createSegments(files []string) []Segment {
	result := []Segment{}

	for i, file := range files {
		signature, _ := pc.SignService.SignFile(file)
		segment := Segment{
			Path:      file,
			Index:     i,
			Signature: signature,
		}
		result = append(result, segment)
	}

	return result
}

func (pc *Packer) compressFiles(files []string, workspace *UpdatePackageWorkspace, pack *UpdatePackage) (string, error) {
	destFilename := path.Join(workspace.Dir, pack.Name + ".pack")
	err := pc.CompressionService.Compress(destFilename, files)

	if err != nil {
		return "", err
	}

	return destFilename, nil
}

func (pc *Packer) createWorkspace(pack *UpdatePackage) (*UpdatePackageWorkspace, error) {
	workspace := &UpdatePackageWorkspace{}

	workspace.Dir = path.Join(pc.WorkspaceDir, pack.Name)
	workspace.ResourceDir = path.Join(workspace.Dir, "resources")

	err := os.Mkdir(workspace.Dir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	err = os.Mkdir(workspace.ResourceDir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return workspace, nil
}

func (pc *Packer) createPackageJson(pack *UpdatePackage, workspace UpdatePackageWorkspace) (string, error) {
	content, err := json.Marshal(pack)

	if err != nil {
		return "", err
	}

	filename := path.Join(workspace.Dir, "config.json")

	err = ioutil.WriteFile(filename, content, os.ModePerm)

	if err != nil {
		return "", err
	}

	return filename, nil
}

func (pc *Packer) chopFiles(source string, outputDir string) ([]string, error) {
	return pc.FileChopService.ChopFile(source, outputDir, Megabyte)
}
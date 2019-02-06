package update

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	"github.com/pkg/errors"
	"nova/segmentdownloader"
	"nova/shared"
	"nova/updatepackage"
)

func NewUpdateHandler() *handler {
	return &handler{}
}

type handler struct {
	parameters Command
	downloader segmentdownloader.SegmentDownloader
	unpacker updatepackage.Unpacker
	dockerService DockerContainerService
	tempDir string
}


func (h *handler) GetIdentifier() string {
	return Identifier
}

func (h *handler) SetParameters(command interface{}) {
	parameters, ok := command.(Command)

	if !ok {
		log.Fatalf("Could not parse parameters")
	}

	h.parameters = parameters
}

func (h *handler) SetByteParameters(data []byte) {

}

func (h *handler) Do() error {
	deploymentPackage := h.extractDeploymentPackage()
	segments := h.download(deploymentPackage)

	files := h.extract(deploymentPackage.Signature, segments)

	dockerFile := h.getDockerTar(files)

	h.replace(dockerFile)
	return nil
}

func (h *handler) getDockerTar(files []string) string {
	for _, file := range files {
		if strings.HasSuffix(file, ".tar") {
			return file
		}
	}
	return ""
}

func (h *handler) Undo() error {
  //TODO: Implement
  return errors.New("This method is not implemented")
}

func (h *handler) IsExecutable() bool {
	return true
}

func (h *handler) download(deployPackage updatepackage.UpdatePackageDeployment) []updatepackage.Segment {
	return h.downloader.DownloadPackage(deployPackage, h.tempDir)
}

func (h *handler) extractDeploymentPackage() updatepackage.UpdatePackageDeployment {
	url := h.parameters.PayloadURL
	content, err := shared.DefaultHTTPUtil.GetBytes(url)

	if err != nil {
		log.Fatalf("DockerCommandHander: Could not download payload url: %v", url)
	}

	result := &updatepackage.UpdatePackageDeployment{}

	err = json.Unmarshal(content, result)

	if err != nil {
		log.Fatalf("DockerCommandHander: Could not parse payload url: %v", url)
	}

	return *result
}

func (h *handler) buildTempDir() {
	h.tempDir, _ = ioutil.TempDir("", "Update")
}

func (h *handler) extract(signature string, segments []updatepackage.Segment) []string {
	pack := updatepackage.UpdatePackageDeployment{
		Signature:signature,
		Segments:segments,
	}
	h.unpacker.Unpack(pack, h.tempDir)
}

func (h *handler) replace(container string) {
	containerId := "carservice"
	h.dockerService.Stop(containerId)
	h.dockerService.Remove(containerId)
	h.dockerService.Install(container)
	h.dockerService.Start(containerId)
}
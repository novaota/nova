package update

import (
	"fmt"
	"log"
	"os/exec"
)

type DockerContainerService struct {

}

func (s *DockerContainerService) Start(id string) {
	s.Command("run -p 8080:8080 %v", id)
}

func (s *DockerContainerService ) Stop(id string) {
	s.Command("stop $(docker ps -aq)")
}

func (s *DockerContainerService ) Remove(id string) {
	s.Command("rmi $(docker images -q)")
}

func (s *DockerContainerService ) Install(source string) {
	s.Command("import")
}

func (s *DockerContainerService) Command(format string, a ...interface{})  {
	args := fmt.Sprintf(format, a)
	cmd := exec.Command("docker", args)
	err := cmd.Run()

	if err != nil {
		log.Fatalf("Could not execute command %v \n", args)
	}
}
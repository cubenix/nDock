package helpers

import (
	"context"
	"log"

	"github.com/gauravgahlot/watchdock/pb"
)

// GetContainer get container details for a container ID
func GetContainer(c pb.ContainerServiceClient) {
	res, err := c.GetContainer(context.Background(), &pb.GetContainerRequest{Id: "container-id"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}

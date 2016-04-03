package consul

import (
	"log"
	"os"
	"strconv"

	"github.com/hashicorp/consul/api"
)

func Register() *api.Client {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Printf("Error connecting on Consul, %v", err)
		return nil
	}

	registerThis(client)
	log.Printf("Registered on Consul.")
	return client
}

func registerThis(client *api.Client) {
	hostname, _ := os.Hostname()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	portNumber, _ := strconv.Atoi(port)
	client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      "transcribe-" + hostname,
		Name:    "transcribe",
		Port:    portNumber,
		Address: "http://" + hostname,
	})
}

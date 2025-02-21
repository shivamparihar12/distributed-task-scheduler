package main

import (
	"flag"

	"github.com/shivam/distributed-task-scheduler/pkg/worker"
)

var (
	serverPort      = flag.String("worker_port", "", "Port on which the Worker serves requests.")
	coordinatorPort = flag.String("coordinator", ":8080", "Network address of the Coordinator.")
)

func main() {
	flag.Parse()

	worker := worker.NewServer(*serverPort, *coordinatorPort)
	worker.Start()
}

package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	agent_ssrf "github.com/n1k1x86/rasp-agents/ssrf"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	agent, err := agent_ssrf.NewSSRFClient(ctx, "localhost", "50051", "localhost:8000", time.Second*30)
	if err != nil {
		log.Fatal(err)
	}
	agentName := "ssrf agent"
	serviceID := "68cce5d44225685132ecf84c"

	err = agent.RegAgent(agentName, serviceID)

	log.Println("AGENT ID = ", agent.AgentID)
	if err != nil {
		log.Println(err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	<-sig
	cancel()
	time.Sleep(time.Second * 10)
}

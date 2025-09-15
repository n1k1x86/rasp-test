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

	agent, err := agent_ssrf.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	serviceName := "test service"
	serviceDescription := "lets check some description"
	agentName := "ssrf agent"

	hostRules := []string{"metadata.google.internal"}
	ipRules := []string{"127.0.0.1", "192.168.0.0"}
	regexpRules := []string{"(127|10|192\\.168|172\\.(1[6-9]|2\\d|3[01])|169\\.254)\\.[0-9]{1,3}\\.[0-9]{1,3}"}

	updateURL := "test url update"
	err = agent.RegAgent(hostRules, ipRules, regexpRules, serviceName,
		serviceDescription, agentName, updateURL)

	log.Println("AGENT ID = ", agent.GetAgentID())
	if err != nil {
		log.Println(err)
	}

	err = agent.DeactivateAgent(serviceName, agentName)
	if err != nil {
		log.Println(err)
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	<-sig
	cancel()
	time.Sleep(time.Second * 10)
}

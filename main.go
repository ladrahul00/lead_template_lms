package main

import (
	"github.com/wolf00/lead_template_lms/handler"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	lead_template "github.com/wolf00/lead_template_lms/proto/lead_template"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.lead_template"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	lead_template.RegisterLeadTemplateHandler(service.Server(), new(handler.LeadTemplateRequestHandler))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

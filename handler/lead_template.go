package handler

import (
	"context"

	lead_template "github.com/wolf00/lead_template_lms/proto/lead_template"
	"github.com/wolf00/lead_template_lms/services"
)

// LeadTemplateRequestHandler is handler for lead template requests
type LeadTemplateRequestHandler struct {
	services.LeadTemplateService
}

// Create is for creating a new lead template
func (e *LeadTemplateRequestHandler) Create(ctx context.Context, req *lead_template.NewLeadTemplateRequest, rsp *lead_template.NewLeadTemplateResponse) error {
	return e.LeadTemplateService.CreateLeadTemplate(ctx, req, rsp)
}

// Get is for geting lead template by id
func (e *LeadTemplateRequestHandler) Get(ctx context.Context, req *lead_template.LeadTemplateByIdRequest, rsp *lead_template.LeadTemplateResponse) error {
	return e.LeadTemplateService.LeadTeamplateByID(ctx, req, rsp)
}

// All is for getting all availabe lead templates
func (e *LeadTemplateRequestHandler) All(ctx context.Context, req *lead_template.AllLeadTemplateRequest, rsp *lead_template.LeadTemplateListResponse) error {
	return e.LeadTemplateByStatus(ctx, req, rsp)
}

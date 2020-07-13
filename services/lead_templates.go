package services

import (
	"context"
	"fmt"

	"github.com/wolf00/lead_template_lms/db"
	"github.com/wolf00/lead_template_lms/db/models"

	lead_template "github.com/wolf00/lead_template_lms/proto/lead_template"

	log "github.com/micro/go-micro/v2/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// LeadTemplateService hanldes
type LeadTemplateService struct {
}

// LeadTeamplateByID service
func (c *LeadTemplateService) LeadTeamplateByID(ctx context.Context, req *lead_template.LeadTemplateByIdRequest, rsp *lead_template.LeadTemplateResponse) error {
	leadTemplate, err := leadTeamplateByID(ctx, req.TemplateId)
	if err != nil {
		log.Error(err)
		return err
	}
	respLeadTemplate := c.leadTemplateToLeadTemplateResponse(leadTemplate)
	rsp.Name = respLeadTemplate.Name
	rsp.KeyValueTypes = respLeadTemplate.KeyValueTypes
	return nil
}

// CreateLeadTemplate creates a new template
func (c *LeadTemplateService) CreateLeadTemplate(ctx context.Context, req *lead_template.NewLeadTemplateRequest, rsp *lead_template.NewLeadTemplateResponse) error {
	newLeadTemplate := models.LeadTemplate{}
	keyValueTypes := []models.TemplateKeyValueTypes{}
	for i := 0; i < len(req.KeyValueTypes); i++ {
		// TO_DO: Add validations for supported value types
		keyValueTypes = append(keyValueTypes, models.TemplateKeyValueTypes{Key: req.KeyValueTypes[i].Key, ValueType: req.KeyValueTypes[i].ValueType})
	}
	newLeadTemplate.Name = req.Name
	newLeadTemplate.KeyValueTypes = keyValueTypes
	_, err := createLeadTemplate(ctx, newLeadTemplate)
	if err != nil {
		rsp.Message = "failed to create the new template"
		rsp.Status = false
		log.Error(err)
		return err
	}
	rsp.Message = "template added successfully"
	rsp.Status = true
	return nil
}

// LeadTemplateByStatus for all leads by status
func (c *LeadTemplateService) LeadTemplateByStatus(ctx context.Context, req *lead_template.AllLeadTemplateRequest, rsp *lead_template.LeadTemplateListResponse) error {
	leadTemplates, err := leadTemplateByStatus(ctx, req.Status)
	if err != nil {
		return err
	}
	respLeadTemplates := []*lead_template.LeadTemplateResponse{}
	for i := 0; i < len(leadTemplates); i++ {
		respLeadTemplates = append(respLeadTemplates, c.leadTemplateToLeadTemplateResponse(leadTemplates[i]))
	}
	rsp.LeadTemplates = respLeadTemplates
	return nil
}

func (c *LeadTemplateService) leadTemplateToLeadTemplateResponse(leadTemplate models.LeadTemplate) *lead_template.LeadTemplateResponse {
	respLeadTemplate := lead_template.LeadTemplateResponse{}
	leadTemplateKeyValueTypes := []*lead_template.LeadTemplateResponse_KeyValue{}
	for lti := 0; lti < len(leadTemplate.KeyValueTypes); lti++ {
		leadTemplateKeyValueTypes = append(leadTemplateKeyValueTypes, &lead_template.LeadTemplateResponse_KeyValue{Key: leadTemplate.KeyValueTypes[lti].Key, ValueType: leadTemplate.KeyValueTypes[lti].ValueType})
	}
	respLeadTemplate.KeyValueTypes = leadTemplateKeyValueTypes
	respLeadTemplate.Name = leadTemplate.Name
	respLeadTemplate.Id = leadTemplate.ID.Hex()
	return &respLeadTemplate
}

func createLeadTemplate(ctx context.Context, newLeadTemplate models.LeadTemplate) (*mongo.InsertOneResult, error) {
	helper := db.LeadTemplates(ctx)
	// TO_DO: Update the status dynamically after we have approval system and versioning in place
	newLeadTemplate.Status = models.ACTIVE
	return helper.InsertOne(ctx, newLeadTemplate)
}

func leadTemplateByStatus(ctx context.Context, status string) ([]models.LeadTemplate, error) {
	leadTemplates := []models.LeadTemplate{}
	helper := db.LeadTemplates(ctx)
	filter := bson.M{"status": status}
	curr, err := helper.Find(ctx, filter)
	if err != nil {
		return leadTemplates, err
	}
	for curr.Next(ctx) {
		var leadTemplate models.LeadTemplate
		err := curr.Decode(&leadTemplate)
		if err != nil {
			return leadTemplates, err
		}
		leadTemplates = append(leadTemplates, leadTemplate)
	}
	return leadTemplates, nil
}

func leadTeamplateByID(ctx context.Context, id string) (models.LeadTemplate, error) {
	var leadTemplate models.LeadTemplate
	helper := db.LeadTemplates(ctx)
	dbID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Error(err)
		return leadTemplate, err
	}
	filter := bson.M{"_id": dbID}
	err = helper.FindOne(ctx, filter).Decode(&leadTemplate)
	if err != nil {
		log.Error(err)
		return leadTemplate, fmt.Errorf("lead template with the id '%s' is not available", id)
	}
	return leadTemplate, nil
}

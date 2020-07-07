package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Status constants
const (
	ACTIVE   = "ACTIVE"
	DRAFT    = "DRAFT"
	INACTIVE = "INACTIVE"
)

// LeadTemplate type definition
type LeadTemplate struct {
	ID            primitive.ObjectID      `bson:"_id,omitempty"`
	CreatedBy     primitive.ObjectID      `bson:"createdBy"`
	KeyValueTypes []TemplateKeyValueTypes `bson:"keyValueTypes,omitempty"`
	Name          string                  `bson:"name"`
	Status        string                  `bson:"status"`
}

// TemplateKeyValueTypes Key and Value Type
type TemplateKeyValueTypes struct {
	Key       string `bson:"key,omitempty"`
	ValueType string `bson:"valueType,omitempty"`
}

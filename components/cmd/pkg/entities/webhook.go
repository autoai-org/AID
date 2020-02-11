// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package entities

import (
	"time"

	"github.com/autoai-org/aiflow/components/cmd/pkg/requests"
	"github.com/autoai-org/aiflow/components/cmd/pkg/storage"
)

// Webhook defines the basic structure for Webhook
type Webhook struct {
	ID        string    `db:"id"`
	RemoteURL string    `db:"remote_url"`
	Status    string    `db:"status"`
	Event     string    `db:"event"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// TableName defines the tablename for webhooks in database
func (w *Webhook) TableName() string {
	return "webhook"
}

// PK defines the primary key of Webhook
func (w *Webhook) PK() string {
	return "id"
}

// Save stores Webhook into database
func (w *Webhook) Save() error {
	db := storage.GetDefaultDB()
	db.Connect()
	return db.Insert(w)
}

// Trigger makes the http requests to the remoteURL
func (w *Webhook) Trigger(data map[string]string) {
	c := requests.NewClient()
	c.Post(w.RemoteURL, data)
}

// FetchWebhooks returns all webhooks
func FetchWebhooks() []Webhook {
	webhooksPointers := make([]*Webhook, 0)
	db := storage.GetDefaultDB()
	db.Fetch(&webhooksPointers)
	webhooks := make([]Webhook, len(webhooksPointers))
	for i := range webhooksPointers {
		webhooks[i] = *webhooksPointers[i]
	}
	return webhooks
}

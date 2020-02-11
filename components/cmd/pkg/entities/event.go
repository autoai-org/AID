// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package entities

import (
	"time"

	"github.com/autoai-org/aiflow/components/cmd/pkg/storage"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
)

// Event defines the basic structure of a spawned event
type Event struct {
	ID        string    `db:"id"`
	Title     string    `db:"title"`
	Data      string    `db:"data"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Status    string    `db:"status"`
	From      string    `db:"source"`
}

// TableName defines the tablename in database
func (e *Event) TableName() string {
	return "event"
}

// PK defines the primary key of Event
func (e *Event) PK() string {
	return "id"
}

// Save stores event into database
func (e *Event) Save() error {
	e.ID = utilities.GenerateUUIDv4()
	db := storage.GetDefaultDB()
	db.Connect()
	return db.Insert(e)
}

// IEvent is the generic Event interface
type IEvent interface {
	msg() string
	toEvent() Event
}

// Consume handles the event: trigger webhooks, local hooks, save to db, etc.
func Consume(ie IEvent) {
	// fetch all webhooks
	webhooks := FetchWebhooks()
	for i := range webhooks {
		webhooks[i].Trigger(map[string]string{"msg": ie.msg()})
	}
	event := ie.toEvent()
	event.Save()
}

// ImageBuiltEvent is the event that will be triggered after a image is built
type ImageBuiltEvent struct {
	ImageName string
}

func (ibe ImageBuiltEvent) msg() string {
	return "Your Image " + ibe.ImageName + " has been built successfully."
}

func (ibe ImageBuiltEvent) toEvent() Event {
	return Event{ID: utilities.GenerateUUIDv4(), Status: "Finished", Title: ibe.msg(), From: "Image Built", Data: "{ImageName:'" + ibe.ImageName + "'}"}
}

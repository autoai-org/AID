// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package entities

import (
	"time"

	"github.com/autoai-org/aid/components/cmd/pkg/storage"
	"github.com/autoai-org/aid/components/cmd/pkg/utilities"
)

// Node defines the struct of a node, which can be used to deploy model to other servers
type Node struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Address   string    `db:"address"`
	Token     string    `db:"token"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Status    string    `db:"status"`
}

// TableName defines the tablename of node in database
func (n *Node) TableName() string {
	return "node"
}

// PK defines the primary key of node
func (n *Node) PK() string {
	return "id"
}

// Save stores node into database
func (n *Node) Save() error {
	n.ID = utilities.GenerateUUIDv4()
	db := storage.GetDefaultDB()
	db.Connect()
	return db.Insert(n)
}

// FetchNodes returns all nodes in a single call
func FetchNodes() []Node {
	nodePointers := make([]*Node, 0)
	db := storage.GetDefaultDB()
	db.Fetch(&nodePointers)
	nodes := make([]Node, len(nodePointers))
	for i := range nodePointers {
		nodes[i] = *nodePointers[i]
	}
	return nodes
}

// GetNodeByName returns a single node by given node name
func GetNodeByName(nodeName string) Node {
	reqNode := Node{Name: nodeName}
	db := storage.GetDefaultDB()
	err := db.FetchOne(&reqNode)
	utilities.CheckError(err, "Cannot fetch node with the name "+nodeName)
	return reqNode
}

// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package entities

import "time"

// Node defines the struct of a node, which can be used to deploy model to other servers
type Node struct {
	ID        string    `db:"id"`
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

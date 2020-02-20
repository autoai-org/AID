// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package entity

// TrainRecord defines the structure of training record
type TrainRecord struct {
	RayID     string `json:"rayId" db:"RayId"`
	DataId    string `json:"dataset" db:"dataId"`
	Vendor    string `json:"vendor" db:"Vendor"`
	Package   string `json:"package" db:"Package"`
	Solver    string `json:"solver" db:"Solver"`
	Status    string `json:"status" db:"Status"`
	CreatedAt string `json:"createdAt" db:"CreatedAt"`
}

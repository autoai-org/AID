// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package entity

// FileObject defines the file object to be used
type FileObject struct {
	Name string `json:"name" db:"Filename"`
	Size int64  `json:"size" db:"Size"`
	Type string `json:"type" db:"Type"`
	UploadedAt string `json:"uploadedAt" db:"UploadedAt"`
	Filepath string `json:"filepath" db:"Filepath"`
}

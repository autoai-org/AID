// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package entity

// FileObject defines the file object to be used
type FileObject struct {
	Name string "json:`name`"
	Size int64  "json:`size`"
}

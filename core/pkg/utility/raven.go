// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package utility

import (
	raven "github.com/getsentry/raven-go"
)

// InitRaven initialize Sentry Settings
func InitRaven() {
	raven.SetDSN("https://fac4164f7c644a27bfb34b748a1c56b5:8db93294332f4df087560540d71ca7ae@sentry.io/1300718")
}

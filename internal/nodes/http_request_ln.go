// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved.
//go:build !plus
// +build !plus

package nodes

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
)

const (
	LNExpiresHeader = "X-Edge-Ln-Expires"
)

func (this *HTTPRequest) checkLnRequest() bool {
	return false
}

func (this *HTTPRequest) getLnOrigin() *serverconfigs.OriginConfig {
	return nil
}

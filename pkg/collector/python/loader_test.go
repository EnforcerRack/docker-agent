// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build python && test
// +build python,test

package python

import (
	"testing"
)

func TestLoadWheelCheck(t *testing.T) {
	testLoadWheelCheck(t)
}

func TestLoadCustomCheck(t *testing.T) {
	testLoadCustomCheck(t)
}

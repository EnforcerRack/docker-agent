// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package containercheck

import (
	"go.uber.org/fx"

	"github.com/DataDog/datadog-agent/comp/process/types"
	"github.com/DataDog/datadog-agent/pkg/process/checks"
)

var _ types.CheckComponent = (*check)(nil)

type check struct {
	containerCheck *checks.ContainerCheck
}

type result struct {
	fx.Out

	Check     types.ProvidesCheck
	Component Component
}

func newCheck() result {
	c := &check{
		containerCheck: checks.NewContainerCheck(),
	}
	return result{
		Check: types.ProvidesCheck{
			CheckComponent: c,
		},
		Component: c,
	}
}

func (c *check) Object() checks.Check {
	return c.containerCheck
}

// Copyright (c) 2021 Dell Inc., or its subsidiaries. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0

package prechecks

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

// OpenshiftVersionValidator will validate the openshift version of the cluster
type OpenshiftVersionValidator struct {
	MinimumVersion string
	MaximumVersion string
	ClusterData    []byte
	K8sClient      K8sClientVersionInterface
	Logger         echo.Logger
}

// Validate will validate the version of the openshift cluster is between the min/max supported versions
func (k OpenshiftVersionValidator) Validate() error {
	isOpenshift, err := k.K8sClient.IsOpenShift(k.ClusterData)
	if err != nil {
		return err
	}
	if !isOpenshift {
		k.Logger.Info("cluster is k8s, skipping openshift version validator")
		return nil
	}
	version, err := k.K8sClient.GetVersion(k.ClusterData)
	if err != nil {
		return err
	}
	minVersion, err := strconv.ParseFloat(k.MinimumVersion, 64)
	if err != nil {
		return err
	}
	maxVersion, err := strconv.ParseFloat(k.MaximumVersion, 64)
	if err != nil {
		return err
	}
	currentVersion, err := strconv.ParseFloat(version, 64)
	if err != nil {
		return err
	}
	if currentVersion < minVersion {
		return fmt.Errorf("version %s is less than minimum supported version of %s", version, k.MinimumVersion)
	}
	if currentVersion > maxVersion {
		return fmt.Errorf("version %s is greater than maximum supported version of %s", version, k.MaximumVersion)
	}
	return nil
}

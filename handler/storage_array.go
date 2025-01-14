// Copyright (c) 2021 Dell Inc., or its subsidiaries. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0

package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/dell/csm-deployment/model"
	"github.com/dell/csm-deployment/utils"
	"github.com/labstack/echo/v4"
)

// CreateStorageArray godoc
// @Summary Create a new storage array
// @Description Create a new storage array
// @ID create-storage-array
// @Tags storage-array
// @Accept  json
// @Produce  json
// @Param storageArray body storageArrayCreateRequest true "Storage Array info for creation"
// @Success 201 {object} storageArrayResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router /storage-arrays [post]
func (h *StorageArrayHandler) CreateStorageArray(c echo.Context) error {
	var storageArray model.StorageArray
	req := &storageArrayCreateRequest{}
	if err := req.bind(c, &storageArray); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewErrorResponse(http.StatusUnprocessableEntity, utils.ErrorSeverity, "", err))
	}

	arrayType, err := h.arrayStore.GetTypeByTypeName(strings.ToLower(req.StorageArrayType))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewErrorResponse(http.StatusUnprocessableEntity, utils.ErrorSeverity, "", err))
	}
	storageArray.StorageArrayTypeID = arrayType.ID

	if err := h.arrayStore.Create(&storageArray); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(http.StatusInternalServerError, utils.CriticalSeverity, "", err))
	}
	return c.JSON(http.StatusCreated, newStorageArrayResponse(&storageArray))
}

// UpdateStorageArray modifies a storage array
// @Summary Update a storage array
// @Description Update a storage array
// @ID update-storage-array
// @Tags storage-array
// @Accept  json
// @Produce  json
// @Param id path string true "Storage Array ID"
// @Param storageArray body storageArrayUpdateRequest true "Storage Array info for update"
// @Success 204 "No Content"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router /storage-arrays/{id} [patch]
func (h *StorageArrayHandler) UpdateStorageArray(c echo.Context) error {
	arrayID := c.Param("id")
	id, err := strconv.Atoi(arrayID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewErrorResponse(http.StatusUnprocessableEntity, utils.ErrorSeverity, "", err))
	}
	storageArray, err := h.arrayStore.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(http.StatusInternalServerError, utils.CriticalSeverity, "", err))
	}
	if storageArray == nil {
		return c.JSON(http.StatusNotFound, utils.NewErrorResponse(http.StatusNotFound, utils.ErrorSeverity, "", err))
	}

	var tmpStorageArray model.StorageArray
	req := &storageArrayUpdateRequest{}
	if err := req.bind(c, &tmpStorageArray); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewErrorResponse(http.StatusUnprocessableEntity, utils.ErrorSeverity, "", err))
	}

	if req.StorageArrayType != "" {
		arrayType, err := h.arrayStore.GetTypeByTypeName(strings.ToLower(req.StorageArrayType))
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, utils.NewErrorResponse(http.StatusUnprocessableEntity, utils.ErrorSeverity, "", err))
		}

		storageArray.StorageArrayType = *arrayType
		storageArray.StorageArrayTypeID = arrayType.ID
	}

	// update other properties
	if tmpStorageArray.UniqueID != "" {
		storageArray.UniqueID = tmpStorageArray.UniqueID
	}
	if tmpStorageArray.Password != nil {
		storageArray.Password = tmpStorageArray.Password
	}
	if tmpStorageArray.Username != "" {
		storageArray.Username = tmpStorageArray.Username
	}
	if tmpStorageArray.ManagementEndpoint != "" {
		storageArray.ManagementEndpoint = tmpStorageArray.ManagementEndpoint
	}

	if err := h.arrayStore.Update(storageArray); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(http.StatusInternalServerError, utils.CriticalSeverity, "", err))
	}
	return c.JSON(http.StatusNoContent, nil)
}

// ListStorageArrays godoc
// @Summary List all storage arrays
// @Description List all storage arrays
// @ID list-storage-arrays
// @Tags storage-array
// @Accept  json
// @Produce  json
// @Param unique_id query string false "Unique ID"
// @Param storage_type query string false "Storage Type"
// @Success 202 {object} []storageArrayResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router /storage-arrays [get]
func (h *StorageArrayHandler) ListStorageArrays(c echo.Context) error {
	uniqueID := c.QueryParam("unique_id")
	storageTypeName := c.QueryParam("storage_type")
	var arrays []model.StorageArray
	var err error
	if uniqueID != "" {
		arrays, err = h.arrayStore.GetAllByUniqueID(uniqueID)
	} else if storageTypeName != "" {
		arrays, err = h.arrayStore.GetAllByStorageType(storageTypeName)
	} else {
		arrays, err = h.arrayStore.GetAll()
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(http.StatusInternalServerError, utils.CriticalSeverity, "", err))
	}
	resp := make([]*storageArrayResponse, 0)
	for _, arr := range arrays {
		resp = append(resp, newStorageArrayResponse(&arr))
	}
	return c.JSON(http.StatusOK, resp)
}

// GetStorageArray godoc
// @Summary Get storage array
// @Description Get storage array
// @ID get-storage-array
// @Tags storage-array
// @Accept  json
// @Produce  json
// @Param id path string true "Storage Array ID"
// @Success 200 {object} storageArrayResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router /storage-arrays/{id} [get]
func (h *StorageArrayHandler) GetStorageArray(c echo.Context) error {
	arrayID := c.Param("id")
	id, err := strconv.Atoi(arrayID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewErrorResponse(http.StatusUnprocessableEntity, utils.ErrorSeverity, "", err))
	}
	storageArray, err := h.arrayStore.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(http.StatusInternalServerError, utils.CriticalSeverity, "", err))
	}
	if storageArray == nil {
		return c.JSON(http.StatusNotFound, utils.NewErrorResponse(http.StatusNotFound, utils.ErrorSeverity, "", err))
	}
	return c.JSON(http.StatusOK, newStorageArrayResponse(storageArray))
}

// DeleteStorageArray godoc
// @Summary Delete storage array
// @Description Delete storage array
// @ID delete-storage-array
// @Tags storage-array
// @Accept  json
// @Produce  json
// @Param id path string true "Storage Array ID"
// @Success 200 {string} string "Success"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Security ApiKeyAuth
// @Router /storage-arrays/{id} [delete]
func (h *StorageArrayHandler) DeleteStorageArray(c echo.Context) error {
	arrayID := c.Param("id")
	id, err := strconv.Atoi(arrayID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewErrorResponse(http.StatusUnprocessableEntity, utils.ErrorSeverity, "", err))
	}
	storageArray, err := h.arrayStore.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(http.StatusInternalServerError, utils.CriticalSeverity, "", err))
	}
	if storageArray == nil {
		return c.JSON(http.StatusNotFound, utils.NewErrorResponse(http.StatusNotFound, utils.ErrorSeverity, "", err))
	}
	if err := h.arrayStore.Delete(storageArray); err != nil {
		c.Logger().Errorf("error deleting storage array: %+v", err)
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(http.StatusInternalServerError, utils.CriticalSeverity, "", err))
	}
	return c.JSON(http.StatusOK, newStorageArrayResponse(storageArray))
}

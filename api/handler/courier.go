package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Courier godoc
// @ID create_courier
// @Router /courier [POST]
// @Summary Create Courier
// @Description Create Courier
// @Tags Courier
// @Accept json
// @Produce json
// @Param Courier body models.CreateCourier true "CreateCourierRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateCourier(c *gin.Context) {

	var createCourier models.CreateCourier

	err := c.ShouldBindJSON(&createCourier)
	if err != nil {
		h.handlerResponse(c, "create courier", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Courier().Create(context.Background(), &createCourier)
	if err != nil {
		h.handlerResponse(c, "storage.courier.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.Courier().GetByID(context.Background(), &models.CourierPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.courier.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create courier", http.StatusCreated, resp)
}

// Get By ID Courier godoc
// @ID get_by_id_Courier
// @Router /courier/{id} [GET]
// @Summary Get By ID Courier
// @Description Get By ID Courier
// @Tags Courier
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdCourier(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id courier", http.StatusBadRequest, "invalid courier id")
		return
	}

	resp, err := h.storages.Courier().GetByID(context.Background(), &models.CourierPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.courier.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id courier", http.StatusCreated, resp)
}

// Get List Courier godoc
// @ID get_list_courier
// @Router /courier [GET]
// @Summary Get List Courier
// @Description Get List Courier
// @Tags Courier
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListCourier(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list courier", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list Courier", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Courier().GetList(context.Background(), &models.GetListCourierRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.courier.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list courier response", http.StatusOK, resp)
}

// Get Update Courier godoc
// @ID update_courier
// @Router /courier/{id} [PUT]
// @Summary Update Courier
// @Description Update Courier
// @Tags Courier
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param courier body models.UpdateCourier true "UpdateCourierRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateCourier(c *gin.Context) {

	var updateCourier models.UpdateCourier

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id courier", http.StatusBadRequest, "invalid courier id")
		return
	}

	err := c.ShouldBindJSON(&updateCourier)
	if err != nil {
		h.handlerResponse(c, "update courier", http.StatusBadRequest, err.Error())
		return
	}

	updateCourier.Id = id

	rowsAffected, err := h.storages.Courier().Update(context.Background(), &updateCourier)
	if err != nil {
		h.handlerResponse(c, "storage.courier.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.courier.update", http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := h.storages.Courier().GetByID(context.Background(), &models.CourierPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.courier.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update courier", http.StatusAccepted, resp)
}

// Update Patch Courier godoc
// @ID updat_patch_courier
// @Router /courier/{id} [PATCH]
// @Summary Update Patch Courier
// @Description Update Patch Courier
// @Tags Courier
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param courier body models.PatchRequest true "UpdatePatchCourierRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdatePatchCourier(c *gin.Context) {

	var object models.PatchRequest

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id courier", http.StatusBadRequest, "invalid courier id")
		return
	}

	err := c.ShouldBindJSON(&object)
	if err != nil {
		h.handlerResponse(c, "update patch courier", http.StatusBadRequest, err.Error())
		return
	}

	object.ID = id

	rowsAffected, err := h.storages.Courier().Patch(context.Background(), &object)
	if err != nil {
		h.handlerResponse(c, "storage.courier.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.courier.patch", http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := h.storages.Courier().GetByID(context.Background(), &models.CourierPrimaryKey{Id: object.ID})
	if err != nil {
		h.handlerResponse(c, "storage.courier.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update patch courier", http.StatusAccepted, resp)
}

// Delete Courier godoc
// @ID delete_courier
// @Router /courier/{id} [DELETE]
// @Summary Delete Courier
// @Description Delete Courier
// @Tags Courier
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param courier body models.CourierPrimaryKey true "DeleteCourierRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteCourier(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id courier", http.StatusBadRequest, "invalid courier id")
		return
	}

	err := h.storages.Courier().Delete(context.Background(), &models.CourierPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.courier.update", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update courier", http.StatusAccepted, nil)
}

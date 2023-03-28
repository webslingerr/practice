package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Category godoc
// @ID create_category
// @Router /category [POST]
// @Summary Create Category
// @Description Create Category
// @Tags Category
// @Accept json
// @Produce json
// @Param Category body models.CreateCategory true "CreateCategoryRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateCategory(c *gin.Context) {

	var createCategory models.CreateCategory

	err := c.ShouldBindJSON(&createCategory)
	if err != nil {
		h.handlerResponse(c, "create category", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Category().Create(context.Background(), &createCategory)
	if err != nil {
		h.handlerResponse(c, "storage.Category.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.Category().GetByID(context.Background(), &models.CategoryPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Category.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Category", http.StatusCreated, resp)
}

// Get By ID Category godoc
// @ID get_by_id_category
// @Router /category/{id} [GET]
// @Summary Get By ID Category
// @Description Get By ID Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdCategory(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id Category", http.StatusBadRequest, "invalid Category id")
		return
	}

	resp, err := h.storages.Category().GetByID(context.Background(), &models.CategoryPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Category.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id Category", http.StatusCreated, resp)
}

// Get List Category godoc
// @ID get_list_Category
// @Router /category [GET]
// @Summary Get List Category
// @Description Get List Category
// @Tags Category
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListCategory(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list Category", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list Category", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Category().GetList(context.Background(), &models.GetListCategoryRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.Category.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list Category response", http.StatusOK, resp)
}

// Get Update Category godoc
// @ID update_category
// @Router /category/{id} [PUT]
// @Summary Update Category
// @Description Update Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param category body models.UpdateCategory true "UpdateCategoryRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateCategory(c *gin.Context) {

	var updateCategory models.UpdateCategory

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id Category", http.StatusBadRequest, "invalid Category id")
		return
	}

	err := c.ShouldBindJSON(&updateCategory)
	if err != nil {
		h.handlerResponse(c, "update Category", http.StatusBadRequest, err.Error())
		return
	}

	updateCategory.Id = id

	rowsAffected, err := h.storages.Category().Update(context.Background(), &updateCategory)
	if err != nil {
		h.handlerResponse(c, "storage.Category.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Category.update", http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := h.storages.Category().GetByID(context.Background(), &models.CategoryPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Category.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update Category", http.StatusAccepted, resp)
}

// Update Patch Category godoc
// @ID updat_patch_category
// @Router /category/{id} [PATCH]
// @Summary Update Patch Category
// @Description Update Patch Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Category body models.PatchRequest true "UpdatePatchCategoryRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdatePatchCategory(c *gin.Context) {

	var object models.PatchRequest

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id Category", http.StatusBadRequest, "invalid Category id")
		return
	}

	err := c.ShouldBindJSON(&object)
	if err != nil {
		h.handlerResponse(c, "update patch Category", http.StatusBadRequest, err.Error())
		return
	}

	object.ID = id

	rowsAffected, err := h.storages.Category().Patch(context.Background(), &object)
	if err != nil {
		h.handlerResponse(c, "storage.Category.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Category.patch", http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := h.storages.Category().GetByID(context.Background(), &models.CategoryPrimaryKey{Id: object.ID})
	if err != nil {
		h.handlerResponse(c, "storage.Category.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update patch Category", http.StatusAccepted, resp)
}

// Delete Category godoc
// @ID delete_category
// @Router /category/{id} [DELETE]
// @Summary Delete Category
// @Description Delete Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Category body models.CategoryPrimaryKey true "DeleteCategoryRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteCategory(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id Category", http.StatusBadRequest, "invalid Category id")
		return
	}

	err := h.storages.Category().Delete(context.Background(), &models.CategoryPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Category.update", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update Category", http.StatusAccepted, nil)
}

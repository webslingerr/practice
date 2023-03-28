package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Product godoc
// @ID create_product
// @Router /product [POST]
// @Summary Create Product
// @Description Create Product
// @Tags Product
// @Accept json
// @Produce json
// @Param Product body models.CreateProduct true "CreateProductRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateProduct(c *gin.Context) {

	var createProduct models.CreateProduct

	err := c.ShouldBindJSON(&createProduct)
	if err != nil {
		h.handlerResponse(c, "create Product", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Product().Create(context.Background(), &createProduct)
	if err != nil {
		h.handlerResponse(c, "storage.Product.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.Product().GetByID(context.Background(), &models.ProductPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Product.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Product", http.StatusCreated, resp)
}

// Get By ID Product godoc
// @ID get_by_id_product
// @Router /product/{id} [GET]
// @Summary Get By ID Product
// @Description Get By ID Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdProduct(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id Product", http.StatusBadRequest, "invalid Product id")
		return
	}

	resp, err := h.storages.Product().GetByID(context.Background(), &models.ProductPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Product.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id Product", http.StatusCreated, resp)
}

// Get List Product godoc
// @ID get_list_Product
// @Router /product [GET]
// @Summary Get List Product
// @Description Get List Product
// @Tags Product
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListProduct(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list Product", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list Product", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Product().GetList(context.Background(), &models.GetListProductRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.Product.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list Product response", http.StatusOK, resp)
}

// Get Update Product godoc
// @ID update_Product
// @Router /product/{id} [PUT]
// @Summary Update Product
// @Description Update Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Product body models.UpdateProduct true "UpdateProductRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateProduct(c *gin.Context) {

	var updateProduct models.UpdateProduct

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id Product", http.StatusBadRequest, "invalid Product id")
		return
	}

	err := c.ShouldBindJSON(&updateProduct)
	if err != nil {
		h.handlerResponse(c, "update Product", http.StatusBadRequest, err.Error())
		return
	}

	updateProduct.Id = id

	rowsAffected, err := h.storages.Product().Update(context.Background(), &updateProduct)
	if err != nil {
		h.handlerResponse(c, "storage.Product.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Product.update", http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := h.storages.Product().GetByID(context.Background(), &models.ProductPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Product.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update Product", http.StatusAccepted, resp)
}

// Update Patch Product godoc
// @ID updat_patch_Product
// @Router /product/{id} [PATCH]
// @Summary Update Patch Product
// @Description Update Patch Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Product body models.PatchRequest true "UpdatePatchProductRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdatePatchProduct(c *gin.Context) {

	var object models.PatchRequest

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id Product", http.StatusBadRequest, "invalid customer id")
		return
	}

	err := c.ShouldBindJSON(&object)
	if err != nil {
		h.handlerResponse(c, "update patch Product", http.StatusBadRequest, err.Error())
		return
	}

	object.ID = id

	rowsAffected, err := h.storages.Product().Patch(context.Background(), &object)
	if err != nil {
		h.handlerResponse(c, "storage.Product.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Product.patch", http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := h.storages.Product().GetByID(context.Background(), &models.ProductPrimaryKey{Id: object.ID})
	if err != nil {
		h.handlerResponse(c, "storage.Product.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update patch Product", http.StatusAccepted, resp)
}

// Delete Product godoc
// @ID delete_Product
// @Router /product/{id} [DELETE]
// @Summary Delete Product
// @Description Delete Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Product body models.ProductPrimaryKey true "DeleteProductRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteProduct(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id Product", http.StatusBadRequest, "invalid Product id")
		return
	}

	err := h.storages.Product().Delete(context.Background(), &models.ProductPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Product.update", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update Product", http.StatusAccepted, nil)
}

package product

import (
	"ecommerece/util"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID)

	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "Pls, give me a valid product id")
		return
	}

	err = h.svc.Delete(pId)

	if err != nil {

		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	util.SendData(w, http.StatusOK, "Successfully deleted product")
}

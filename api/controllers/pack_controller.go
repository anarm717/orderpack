package controllers

import (
	"GoOrderPackProject/api/config"
	"GoOrderPackProject/api/responses"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sort"
)

// GetPackSizes godoc
// @Summary Get pack sizes
// @Description Get pack sizes
// @Tags packs
// @Accept  json
// @Produce  json
// @Success 200 {array} int
// @Router /pack-sizes [get]
func (server *Server) GetPackSizes(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting pack sizes: ", config.ReadConfig().PackSizes)
	responses.JSON(w, http.StatusOK, config.ReadConfig().PackSizes)
}

// UpdatePackSizes godoc
// @Summary Update pack sizes
// @Description Update pack sizes
// @Param			tags body []int true "Update pack sizes"
// @Tags packs
// @Accept  json
// @Produce  json
// @Success 200 {array} int
// @Router /pack-sizes [post]
func (server *Server) UpdatePackSizes(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		log.Println("Error reading body for update pack sizes: ", err)
		return
	}
	var packSizes []int
	err = json.Unmarshal(body, &packSizes)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		log.Println("Error parsing body for update pack sizes: ", err)
		return
	}
	config.SetPackSizes(packSizes)
	responses.JSON(w, http.StatusOK, config.ReadConfig().PackSizes)
}

// CalculatePacks godoc
// @Summary Calculate pack sizes by quantity
// @Description Calculate pack sizes quantity
// @Param			tags body config.OrderRequest true "Calculate packs"
// @Tags packs
// @Accept  json
// @Produce  json
// @Success 200 {object} map[int]int
// @Router /calculate-packs [post]
func (server *Server) CalculatePacksByQuantity(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		log.Println("Error reading body for CalculatePack: ", err)
		return
	}
	orderRequest := config.OrderRequest{}
	err = json.Unmarshal(body, &orderRequest)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		log.Println("Error parsing body for CalculatePack: ", err)
		return
	}
	fulfillment := GetPacks(orderRequest.OrderQuantity, config.ReadConfig().PackSizes)
	responses.JSON(w, http.StatusOK, fulfillment)
}

func GetPacks(orderQuantity int, packSizes []int) map[int]int {
	log.Println("Calculating packs for:", orderQuantity)
	fulfillment, _ := fulfillOrder(orderQuantity, packSizes)
	for {
		optimized := false
		for packSize, numPacks := range fulfillment {
			fulfillment1, numPacks1 := fulfillOrder(packSize*numPacks, packSizes)
			if numPacks1 < numPacks {
				delete(fulfillment, packSize)
				for packSize2, numPacks2 := range fulfillment1 {
					fulfillment[packSize2] += numPacks2
				}
				optimized = true
			}
		}
		if !optimized {
			break
		}
	}
	log.Println("Result for ", orderQuantity, " is ", fulfillment)
	return fulfillment
}
func fulfillOrder(orderQuantity int, packSizes []int) (map[int]int, int) {
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))
	fulfillment := make(map[int]int)
	remainingQuantity := orderQuantity
	packCount := 0
	beforePackSize := 0
	for _, packSize := range packSizes {
		if remainingQuantity >= packSize {
			if (beforePackSize-remainingQuantity < remainingQuantity-packSize) && (beforePackSize != 0) {
				fulfillment[beforePackSize]++
				remainingQuantity = 0
			} else {
				numPacks := remainingQuantity / packSize
				fulfillment[packSize] = numPacks
				remainingQuantity -= numPacks * packSize
				packCount += numPacks
			}
		}
		beforePackSize = packSize
	}
	if remainingQuantity > 0 {
		minPackSize := packSizes[len(packSizes)-1]
		fulfillment[minPackSize]++
	}
	return fulfillment, packCount
}

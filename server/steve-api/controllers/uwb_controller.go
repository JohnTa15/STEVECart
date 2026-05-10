package controllers

import (
	"fmt"
	"steve-api/models"
	"time"
)

//updating UWBPosition
func updateUWBPosition(uwb_node_id string, x_coord float64, y_coord float64) (models.UWBData, error) {
	var uwbdata models.UWBData
	if uwb_node_id == "" && x_coord == 0.0 && y_coord == 0.0 {
		fmt.Println("Cart not found and x and y coordinates are empty..")
		return uwbdata, nil
	}
	uwbdata.UWB_NODEID = uwb_node_id
	uwbdata.X_Coordinate = x_coord
	uwbdata.Y_Coordinate = y_coord
	uwbdata.LastSeen_UWB = time.Now()
	uwbdata.Description = "Cart position updated"
	return uwbdata, nil
}

func shelvePosition(shelve_id string, x_coord float64, y_coord float64) (models.UWBData, error) {
	var uwbdata models.UWBData
	if shelve_id == "" && x_coord == 0.0 && y_coord == 0.0 {
		fmt.Println("Shelve not found and x and y coordinates are empty..")
		return uwbdata, nil
	}
	uwbdata.UWB_NODEID = shelve_id
	uwbdata.X_Coordinate = x_coord
	uwbdata.Y_Coordinate = y_coord
	uwbdata.Description = "Shelve position updated"
	return uwbdata, nil
}
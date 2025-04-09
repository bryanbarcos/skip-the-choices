package handlers

import (
	//"log"
	"fmt"

	"github.com/pocketbase/pocketbase/core"
)

func AddRestaurant(e *core.RequestEvent) error {

	app := e.App
	restaurantId := e.Request.PathValue("id")

	sqlQuery := fmt.Sprintf("INSERT INTO user_restaurants (user_id, rest_id) VALUES ('hm77ox88p0j5t4v', '%s')", restaurantId)
	_, err := app.DB().NewQuery(sqlQuery).Execute()

	if err != nil {
		fmt.Println("AddRestaurant - Error in query execution")
		return err
	}

	return nil
}

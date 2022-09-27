package users

import (
	"FinalProject/database"
	"FinalProject/helpers"
	"net/http"
	"strconv"
)

func DeleteUser(x http.ResponseWriter, y *http.Request) {
	id := y.URL.Query().Get("id")
	if id == "" {
		responseMsg := map[string]string{"Info": "Empty ID!"}
		helpers.ResponseJSON(x, http.StatusNotFound, responseMsg)
		return
	}
	userID, _ := strconv.Atoi(id)

	database.DB.Exec("delete from users where id = ?", userID)

	responseMsg := map[string]string{"Info": "Successfully deleted!"}
	helpers.ResponseJSON(x, http.StatusOK, responseMsg)

}
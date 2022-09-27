package photoscontroller

import (
	"FinalProject/helpers"
	"net/http"
)

func Home(x http.ResponseWriter, y *http.Request) {

	dataUser := []map[string]interface{}{
		{
			"ID": 1,
			"Title": "myPhoto 1",
			"Caption": "Good Morning!",
			"PhotoUrl": "https://pixabay.com/images/id-986304/",
			"UserID": "001",
		},
		{
			"ID": 2,
			"Title": "myPhoto 2",
			"Caption": "Good Afternoon!",
			"PhotoUrl": "https://pixabay.com/images/id-2277/",
			"UserID": "002",
		},
		{
			"ID": 3,
			"Title": "myPhoto 3",
			"Caption": "Good Night!",
			"PhotoUrl": "https://pixabay.com/images/id-494706/",
			"UserID": "003",
		},
		{
			"ID": 4,
			"Title": "myPhoto 4",
			"Caption": "Good Evening!",
			"PhotoUrl": "https://pixabay.com/images/id-1426859/",
			"UserID": "004",
		},
	}

	helpers.ResponseJSON(x, http.StatusOK, dataUser)

}
package app

import (
	"github.com/jinzhu/gorm"
	"github.com/opensourceai/go-api-service/pkg/e"
	"net/http"
)

func Recover(appG *Gin) {
	if err := recover(); err == gorm.ErrRecordNotFound.Error() {
		appG.Response(http.StatusNotFound, e.INVALID_PARAMS, nil)
	} else if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR, nil)
	}

}

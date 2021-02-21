package helpers

import (
	"github.com/gorilla/mux"
	"github.com/mrbardia72/dark-zarinpal/config"
	"github.com/mrbardia72/dark-zarinpal/zarinpal"
	"net/http"
	"strconv"
)



func CheckPaymentErr(w http.ResponseWriter) (*zarinpal.Zarinpal, error, bool) {
	zarinpal, err := zarinpal.NewZarinpal(config.MERCHAND_ID, config.SANDBOX)
	if err != nil {
		LogWriteHeader(w, "خطا در پرداخت.", http.StatusInternalServerError)
		return nil, nil, true
	}
	return zarinpal, err, false
}

func CheckPrice(w http.ResponseWriter, r *http.Request) (string, bool) {
	vars := mux.Vars(r)
	price, ok := vars["price"]
	if !ok {
		LogWriteHeader(w, "لطفا مبلغ را وارد کنید.", http.StatusBadRequest)
		return "", true
	}
	return price, false
}


func CheckPriceInt(w http.ResponseWriter, price string) (int, error, bool) {
	intPrice, err := strconv.Atoi(price)
	if err != nil {
		LogWriteHeader(w, "لطفا مبلغ را بصورت عدد وارد کنید", http.StatusBadRequest)
		return 0, nil, true
	}
	return intPrice, err, false
}
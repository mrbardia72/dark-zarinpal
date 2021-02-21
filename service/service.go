package service

import (
	"fmt"
	"github.com/mrbardia72/dark-zarinpal/config"
	"github.com/mrbardia72/dark-zarinpal/helpers"
	"net/http"
)

func CallBack(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "برگشت از درگاه")

	authority := r.URL.Query().Get("Authority")
	status := r.URL.Query().Get("Status")

	if authority == "" || status == "" || status != "OK" {
		helpers.LogWriteHeader(w,"خطایی در پرداخت رخ داده است.",http.StatusOK)
		return
	}

	price, done2 := helpers.CheckPrice(w, r)
	if done2 {
		return
	}

	intPrice, err, done := helpers.CheckPriceInt(w, price)
	if done {
		return
	}

	zarinpal, err, done3 := helpers.CheckPaymentErr(w)
	if done3 {
		return
	}

	verified, refId, statusCode, err := zarinpal.PaymentVerification(intPrice, authority)
	if err != nil {
		if statusCode == 101 {
			helpers.LogWriteHeader(w, "این پرداخت موفق بوده و قبلا این عملیات انجام شده است.", http.StatusOK)
			return
		}

		helpers.LogWriteHeader(w, "خطا در پرداخت.", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "پرداخت موفقیت آمیز بود . شماره پیگیری : ", refId)
	fmt.Println(w, "Payment Verified : ", verified, " ,  refId: ", refId, " statusCode: ", statusCode)
}


func Bank(w http.ResponseWriter, r *http.Request) {

	price, done2 := helpers.CheckPrice(w, r)
	if done2 {
		return
	}

	zarinpal, err, done3 := helpers.CheckPaymentErr(w)
	if done3 {
		return
	}

	intPrice, err, done := helpers.CheckPriceInt(w, price)
	if done {
		return
	}

	paymentUrl, authority, statusCode, err := zarinpal.NewPaymentRequest(
		intPrice,
		"http://localhost"+config.SERVER_PORT+"/CallBack"+price,
		"پرداخت دارک کد",
		"darkcode@gmail.com",
		"09360750299",
		)
	if err != nil {
		if statusCode == -3 {
			helpers.LogWriteHeader(w, "مبلغ قابل پرداخت نیست.", http.StatusBadRequest)
			return
		}
		helpers.LogWriteHeader(w, "خطایی در پرداخت رخ داده است.", http.StatusBadRequest)
		return
	}
	//Create Record in DB
	fmt.Println("PaymentURL: ", paymentUrl, " statusCode : ", statusCode, " Authority: ", authority)
	fmt.Println("price",intPrice,"mobile",email)
	http.Redirect(w, r, paymentUrl, 302)
}




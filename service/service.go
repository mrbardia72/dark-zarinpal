package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mrbardia72/dark-zarinpal/config"
	"github.com/mrbardia72/dark-zarinpal/helpers"
	"github.com/mrbardia72/dark-zarinpal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

var logpayCollection = config.DbConfig().Database("zarinpal").Collection("logpay")
var paymentCollection = config.DbConfig().Database("zarinpal").Collection("payment")

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
	date_now := time.Now().Format("02-01-2006")
	time_now := time.Now().Format("15:04:05")

	payment := models.Payment{
		Status: statusCode ,
		Verified:verified,
		Refid:refId,
		Date:date_now,
		Time:time_now,
	}

	//var jsonPayment []byte
	//jsonPayment, _ = json.Marshal(payment)
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println(string(jsonPayment))

	ctx := context.Background()
	insertResult, err := paymentCollection.InsertOne(ctx, &payment)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult)
	fmt.Fprintln(w, "پرداخت یا موقفیت انجام شد : ", " ,  کدپیگیری: ", refId)

}

func GetAllLogPaymet(w http.ResponseWriter, r *http.Request)  {

	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M

	cur, err := logpayCollection.Find(ctx, bson.D{{}})
	if err != nil {
		fmt.Println(err)
	}

	for cur.Next(ctx) {

		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}
	cur.Close(context.Background())
	fmt.Println("get all log payment information")
	json.NewEncoder(w).Encode(results)

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

	paymentUrl, authority,emailUser,descriptionUser,mobileUser, statusCode, err := zarinpal.NewPaymentRequest(
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

	date_now := time.Now().Format("02-01-2006")
	time_now := time.Now().Format("15:04:05")

	logpay := models.Logpay{
		Status: statusCode ,
		Authority:authority,
		Amount:intPrice,
		Email:emailUser,
		Description:descriptionUser,
		Mobile:mobileUser,
		Date:date_now,
		Time:time_now,
	}

	//var jsonLogPay []byte
	//jsonLogPay, _ = json.Marshal(logpay)
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println(string(jsonLogPay))

	//mongo
	ctx := context.Background()
	insertResult, err := logpayCollection.InsertOne(ctx, &logpay)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult)
	http.Redirect(w, r, paymentUrl, 302)
}

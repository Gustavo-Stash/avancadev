package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Result struct {
	Status string
}

var coupons Coupons

type Coupon struct {
	Code string
	Status string
}

type Coupons struct {
	Coupon []Coupon
}

func (c Coupons) Check(code string) string {
	for _, item := range c.Coupon {
		if code == item.Code {
			return item.Status
		}
	}
	return "invalid"
}

func main() {
	coupons.Coupon = append(coupons.Coupon, Coupon{Code: "abc", Status: "unused"})
	coupons.Coupon = append(coupons.Coupon, Coupon{Code: "123", Status: "used"})
	coupons.Coupon = append(coupons.Coupon, Coupon{Code: "zzz", Status: "unused"})

	http.HandleFunc("/", home)
	http.ListenAndServe(":9093", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	coupon := r.PostFormValue("coupon")
	status := coupons.Check(coupon)

	result := Result{Status: status}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error converting json")
	}

	fmt.Fprintf(w, string(jsonResult))

}
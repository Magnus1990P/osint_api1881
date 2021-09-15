package main

import (
	"fmt"
	"encoding/json"
)

type maltegoPhoneNumber struct {
	DisplayName	      string //Phone Number
  EntityName	      string //Const value "maltego.PhoneNumber"
  ShortDescription	string //Const value "A telephone number"
  EntityCategory    string //Const value "Personal"
  BaseEntity        string //Const value "maltego.Unknown"

  Phonenumber       string	Phone Number
  phonenumber.countrycode+	string	Country Code
  phonenumber.citycode+	string	City Code
  phonenumber.areacode+	string	Area Code
  phonenumber.lastnumbers+	string	Last Digits
}


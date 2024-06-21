package utils

import (
	// "aot.limousine-service/utils"

	"time"

	"github.com/FourWD/aot.middleware/orm"
	"github.com/FourWD/middleware/common"
	"github.com/google/uuid"
)

func SaveCustomer(payload PayloadCustomer) (bool, string) {
	Customer := new(orm.Customer)
	Customer.ID = uuid.NewString()
	Customer.Code = ""
	Customer.Firstname = payload.Firstname
	Customer.Lastname = payload.Lastname
	Customer.CompanyName = payload.CompanyName
	Customer.PrefixID = payload.PrefixID

	if payload.BirthDay.IsZero() {
		Customer.Birthday = time.Now()
	} else {
		Customer.Birthday = payload.BirthDay
	}

	Customer.IDCardNo = payload.IDCardNo
	Customer.PassportNo = payload.PassportNo
	Customer.CompanyName = payload.CompanyName
	Customer.TaxNo = payload.TaxNo
	Customer.PhoneNo1 = payload.Telephone
	if payload.IsTax {
		Customer.IsTax = true
		Customer.CompanyName = payload.CompanyName
		Customer.Address = payload.AddressCompany
		Customer.Postcode = payload.PostCodeCompany
		Customer.TaxNo = payload.TaxNo
		Customer.PhoneNo1 = payload.PhoneCompany
	}

	common.Database.Create(&Customer)
	GenCustomerCode(Customer.ID)

	return true, Customer.ID
}

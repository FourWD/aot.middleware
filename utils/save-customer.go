package utils

import (
	// "aot.limousine-service/utils"

	"github.com/FourWD/aot.middleware/orm"
	"github.com/FourWD/middleware/common"
	"github.com/google/uuid"
)

func SaveCustomer(payload PayloadCustomer) (bool, string) {
	Customer := new(orm.Customer)
	CustomerAddress := new(orm.CustomerAddress)

	Customer.ID = uuid.NewString()
	Customer.Code = ""
	Customer.Firstname = payload.Firstname
	Customer.Lastname = payload.Lastname
	CustomerAddress.CompanyName = payload.CompanyName
	Customer.PrefixID = payload.PrefixID

	if payload.BirthDay.IsZero() {
		Customer.Birthday = common.NilDate()
	}

	Customer.IDCardNo = payload.IDCardNo
	Customer.PassportNo = payload.PassportNo
	CustomerAddress.CompanyName = payload.CompanyName
	CustomerAddress.TaxNo = payload.TaxNo
	Customer.PhoneNo1 = payload.Telephone
	if payload.IsTax {
		CustomerAddress.CompanyName = payload.CompanyName
		CustomerAddress.Address = payload.AddressCompany
		CustomerAddress.Postcode = payload.PostCodeCompany
		CustomerAddress.TaxNo = payload.TaxNo
	}

	common.Database.Create(&Customer)
	UpdateCustomerCode(Customer.ID)

	return true, Customer.ID
}

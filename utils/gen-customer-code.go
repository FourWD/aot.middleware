package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/FourWD/aot.middleware/orm"
	"github.com/FourWD/middleware/common"
)

func UpdateCustomerCode(id string) string {

	var customer orm.Customer

	sql := fmt.Sprintf(`SELECT running_no,created_at FROM customers WHERE id = '%s'`, id)
	common.Database.Raw(sql).Scan(&customer)
	// common.PrintError("Error sql", err.Error.Error())

	type Count struct {
		Count int64 `json:"count"`
	}

	var count Count
	year := time.Now().Year()
	month := time.Now().Month()

	query := fmt.Sprintf(`SELECT COUNT(*) FROM customers WHERE YEAR(created_at) = %d AND MONTH(created_at) = %d
	AND running_no <= '%d'`, year, month, customer.RunningNo)
	common.Database.Raw(query).Scan(&count.Count)
	common.Print(query, fmt.Sprintf("%d", count.Count))

	fyear := strconv.Itoa(year)
	fmonth := fmt.Sprintf("%02d", month)
	code := fmt.Sprintf("%s%s%05d", fyear[2:4], fmonth, count.Count)

	update := fmt.Sprintf(`UPDATE customers set code = '%s' where id = '%s'`, code, id)
	checkupdate := common.Database.Exec(update)
	if checkupdate.RowsAffected == 0 {
		common.PrintError("Error checkupdate", id)
	}
	common.Print("check", code)

	return ""
}

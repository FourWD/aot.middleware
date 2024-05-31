package utils

import (
	"fmt"

	"github.com/FourWD/aot.middleware/orm"
	"github.com/FourWD/middleware/common"
)

func getCounter(id string) orm.Counter {

	counter := new(orm.Counter)
	sql := fmt.Sprintf(`select id, name as counter_name from counters  where  id = '%s'`, id)
	common.Database.Raw(sql).Scan(&counter)
	common.Print("CountSubModel", sql)

	return *counter

}

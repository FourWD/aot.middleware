package utils

import (
	"fmt"

	"github.com/FourWD/middleware/common"
)

// func UpdateSlipNo(id string, vehciesubmodelid string, counterid string) string {

// 	type Payload struct {
// 		RunningNo             int    `json:"running_no"`
// 		SlipVehicleSubModelID string `json:"slip_vehicle_sub_model_id"`
// 		IYear                 int    `json:"iyear"`
// 		IMonth                int    `json:"imonth"`
// 		VehicleSubModelName   string `json:"vehicle_sub_model_name"`
// 		VehicleSubModelID     string `json:"vehicle_sub_model_id"`
// 		Name                  string `json:"name"`
// 		CountSubModel         int64  `json:"count_sub_model"`
// 	}
// 	//B2-ZZ-2023-08-00001 logicslipno
// 	slip := new(Payload)

// 	sqls := fmt.Sprintf(`SELECT running_no FROM slips WHERE id = '%s'`, id)
// 	common.Database.Raw(sqls).Scan(&slip)
// 	common.PrintError("a", sqls)

// 	type SubModel struct {
// 		Name                  string `json:"name"`
// 		SubModelGroup         string `json:"sub_model_group"`
// 		IsShow                string `json:"is_show"`
// 		CounterID             string `json:"counter_id"`
// 		CounterName           string `json:"counter_name"`
// 		SlipVehicleSubModelID string `json:"slip_vehicle_sub_model_id"`
// 	}
// 	submodel := new(SubModel)
// 	querysubmodel := fmt.Sprintf(`select id,sub_model_group from vehicle_sub_models  where  id = '%s'`, vehciesubmodelid)
// 	common.Database.Raw(querysubmodel).Scan(&submodel)
// 	common.Print("CountSubModel", querysubmodel)
// 	// sql := fmt.Sprintf(`
// 	// SELECT  YEAR(s.created_at) as iyear,  MONTH(s.created_at) as imonth,s.running_no,s.slip_vehicle_sub_model_id, s.counter_id , vsm.name_en as vehicle_sub_model_name, c.name as counter_name
// 	// 	FROM slips s
// 	// 	LEFT JOIN vehicle_sub_models vsm ON s.slip_vehicle_sub_model_id = vsm.id
// 	// 	LEFT JOIN counters c on s.counter_id = c.id
// 	// 	WHERE s.id = '%s'`, id)

// 	sql := fmt.Sprintf(`select count(*) as count_sub_model from vehicle_sub_models where sub_model_group = ( select sub_model_group from vehicle_sub_models where id = '%s')`, vehciesubmodelid)
// 	common.Database.Raw(sql).Scan(&slip)
// 	common.Print("CountSubModel ", fmt.Sprintf("%d", slip.CountSubModel))

// 	type Count struct {
// 		Count        int64 `json:"icount" gorm:"column:icount"`
// 		CountNo      int64 `json:"count1" gorm:"column:count1"`
// 		ReceiptCount int64 `json:"rr_count" gorm:"column:rr_count"`
// 	}

// 	var count Count
// 	// query := fmt.Sprintf(`SELECT COUNT(*) as icount  FROM slips WHERE running_no <= '%d'
// 	// AND YEAR(created_at) = (SELECT YEAR(created_at) FROM slips WHERE id = '%s')
// 	// AND MONTH(created_at) = (SELECT MONTH(created_at) FROM slips WHERE id = '%s')
// 	// AND counter_id = (SELECT counter_id FROM slips WHERE id = '%s')
// 	// AND slip_vehicle_sub_model_id = (SELECT slip_vehicle_sub_model_id FROM slips WHERE id = '%s')`, slip.RunningNo, id, id, id, id)
// 	// common.Database.Raw(query).Scan(&count)

// 	query := fmt.Sprintf(`SELECT COUNT(*) as icount , count(*) +1 as count1 FROM slips WHERE running_no <= '%d'
// 	AND YEAR(created_at) = (SELECT YEAR(created_at) FROM slips WHERE id = '%s')
// 	AND MONTH(created_at) = (SELECT MONTH(created_at) FROM slips WHERE id = '%s')
// 	AND counter_id = (SELECT counter_id FROM slips WHERE id = '%s')
// 	AND slip_vehicle_sub_model_id in (select id from vehicle_sub_models where sub_model_group  =  '%s' )`, slip.RunningNo, id, id, id, submodel.SubModelGroup)
// 	common.Database.Raw(query).Scan(&count)
// 	common.Print(query, fmt.Sprintf("%d", count.Count))

// 	queryreceipt := fmt.Sprintf(`SELECT COUNT(*) as rr_count  FROM slips WHERE running_no <= '%d'
// 	AND YEAR(created_at) = (SELECT YEAR(created_at) FROM slips WHERE id = '%s')
// 	AND MONTH(created_at) = (SELECT MONTH(created_at) FROM slips WHERE id = '%s')
// 	AND counter_id = (SELECT counter_id FROM slips WHERE id = '%s')`, slip.RunningNo, id, id, id)
// 	common.Database.Raw(queryreceipt).Scan(&count)
// 	common.Print(queryreceipt, fmt.Sprintf("%d", count.ReceiptCount))

// 	countrunning := count.CountNo //count+1
// 	common.Print("count +1", fmt.Sprintf("%d", countrunning))
// 	resultmod := count.Count%slip.CountSubModel + 1
// 	date := time.Now().Format("200601")

// 	// slipno := generateSlipNumber(submodel, countrunning, date, int(resultmod))
// 	// receiptno := generateReceiptNumber(submodel, date, receipt)

// 	// fmonth := fmt.Sprintf("%02d", slip.IMonth)
// 	// code := fmt.Sprintf(`%d%s-%05d`, slip.IYear, fmonth, count.CountNo)
// 	// common.Print("code qqqa", code)
// 	// date := time.Now().Format("200601")
// 	// X2 - A3 - 202308 - 00060

// 	// querycounter := fmt.Sprintf(`select id, name as counter_name from counters  where  id = '%s'`, counterid)
// 	// common.Database.Raw(querycounter).Scan(&submodel)
// 	// common.Print("CountSubModel", querycounter)
// 	counter := getCounter(id)
// 	submodel.CounterID = counter.ID
// 	submodel.CounterName = counter.Name

// 	sqlwhere := fmt.Sprintf(`with cte as (select row_number() over (order by row_order ) as row_num,name , id as slip_vehicle_sub_model_id from vehicle_sub_models where sub_model_group = '%s')
// 				select * from cte where row_num = '%d' `, submodel.SubModelGroup, resultmod)
// 	common.Database.Raw(sqlwhere).Scan(&submodel)
// 	common.Print(sqlwhere, submodel.Name) // get sub_model_group

// 	slipno := fmt.Sprintf(`%s-%s-%s-%06d`, submodel.Name, submodel.CounterName, date, countrunning)

// 	receiptno := fmt.Sprintf(`%s-%s-%s%06d`, submodel.CounterName, date, submodel.Name, count.ReceiptCount)

// 	common.Print("code check slipno", slipno)

// 	updateslip := fmt.Sprintf(`UPDATE slips set slip_no = '%s' , slip_vehicle_sub_model_id = '%s', receipt_no = '%s' where id = '%s'`, slipno, submodel.SlipVehicleSubModelID, receiptno, id)
// 	checkupdate := common.Database.Exec(updateslip)
// 	if checkupdate.RowsAffected == 0 {
// 		common.PrintError("Error checkupdate", id)
// 	}

// 	return ""
// }

func UpdateSlipNo(id string) string {
	print("update-slip")
	slipno := GenSlipRetail(id)

	updateslip := fmt.Sprintf(`UPDATE slips set slip_no = '%s' where id = '%s'`, slipno, id)

	println(updateslip)
	checkupdate := common.Database.Debug().Exec(updateslip)
	if checkupdate.RowsAffected == 0 {
		common.PrintError("Error checkupdate", id)
	}
	return ""
}

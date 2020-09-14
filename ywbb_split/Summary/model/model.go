package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
	_ "ywbb_split/Summary/logconfig"
)

// 定义一个全局对象db

var db *sql.DB

// 定义一个初始化数据库的函数
func InitDB() (err error) {
	// DSN:Data Source Name
	dsn := "输入数据库账号密码"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// 插入数据
func InsertRowDemo(sqlStr string, data []string) {
	fmt.Println(data)
	//  AreaName（地区）, month(月), week(周), area(南北两区)
	ret, err := db.Exec(sqlStr, data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7], data[8], data[9],
		data[10], data[11], data[12], data[13], data[14], data[15], "北京", 7, 5, 2)
	//ret, err := db.Exec(sqlStr,  "as")
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

type RowData struct {
	City, Region, Num, Week, C, D, E, F, G, I, J, L, M, O, P, R, S, W, X, Y, Z, K, N, Q, T, U, AA, AB string
}
type S struct {
	sum int
}

// 查询业绩管理
func QueryMultiRowDemo(month, area string) []RowData {
	err := InitDB()
	if err != nil {
		log.Println(err)
		panic("业绩管理初始化数据库连接出错")
	}
	// 16个sum

	var sqlStr string
	if strings.Contains("南区北区", area) {

		//sqlStr = fmt.Sprintf("SELECT  week,SUM(CASE ISNULL(new_signed_num_target) or new_signed_num_target='' WHEN 1 THEN 0 ELSE new_signed_num_target END), SUM(CASE ISNULL(new_signed_num_actual) or new_signed_num_actual='' WHEN 1 THEN 0 ELSE new_signed_num_actual END), SUM(CASE ISNULL(new_sbdh_num_target) or new_sbdh_num_target='' WHEN 1 THEN 0 ELSE new_sbdh_num_target END), SUM(CASE ISNULL(new_sbdh_num_actual) or new_sbdh_num_actual='' WHEN 1 THEN 0 ELSE new_sbdh_num_actual END), SUM(CASE ISNULL(signing_fee_target) or signing_fee_target='' WHEN 1 THEN 0 ELSE signing_fee_target END),SUM(CASE ISNULL(signing_fee_actual) or signing_fee_actual='' WHEN 1 THEN 0 ELSE signing_fee_actual END), SUM(CASE ISNULL(activated_num_target) or activated_num_target='' WHEN 1 THEN 0 ELSE activated_num_target END), SUM(CASE ISNULL(activated_num_actual) or activated_num_actual='' WHEN 1 THEN 0 ELSE activated_num_actual END), SUM(CASE ISNULL(sbdh_activated_num_target) or sbdh_activated_num_target='' WHEN 1 THEN 0 ELSE sbdh_activated_num_target END), SUM(CASE ISNULL(sbdh_activated_num_actual) or sbdh_activated_num_actual='' WHEN 1 THEN 0 ELSE sbdh_activated_num_actual END), SUM(CASE ISNULL(activated_fee_target) or activated_fee_target='' WHEN 1 THEN 0 ELSE activated_fee_target END), SUM(CASE ISNULL(activated_fee_actual) or activated_fee_actual='' WHEN 1 THEN 0 ELSE activated_fee_actual END),SUM(CASE ISNULL(sbdh_old_added_num) or sbdh_old_added_num='' WHEN 1 THEN 0 ELSE sbdh_old_added_num END), SUM(CASE ISNULL(sbdh_reduce_num) or sbdh_reduce_num='' WHEN 1 THEN 0 ELSE sbdh_reduce_num END), SUM(CASE ISNULL(loss_customer_num) or loss_customer_num='' WHEN 1 THEN 0 ELSE loss_customer_num END), SUM(CASE ISNULL(loss_num) or loss_num='' WHEN 1 THEN 0 ELSE loss_num END) FROM `kpitable` where month=\"%s\"  and region=\"%s\"  GROUP BY %s", month, area, "week")
		sqlStr = fmt.Sprintf("SELECT  week,SUM(CASE ISNULL(new_signed_num_target) or new_signed_num_target='' WHEN 1 THEN 0 ELSE new_signed_num_target END), SUM(CASE ISNULL(new_signed_num_actual) or new_signed_num_actual='' WHEN 1 THEN 0 ELSE new_signed_num_actual END), SUM(CASE ISNULL(new_sbdh_num_target) or new_sbdh_num_target='' WHEN 1 THEN 0 ELSE new_sbdh_num_target END), SUM(CASE ISNULL(new_sbdh_num_actual) or new_sbdh_num_actual='' WHEN 1 THEN 0 ELSE new_sbdh_num_actual END), SUM(CASE ISNULL(signing_fee_target) or signing_fee_target='' WHEN 1 THEN 0 ELSE signing_fee_target END),SUM(CASE ISNULL(signing_fee_actual) or signing_fee_actual='' WHEN 1 THEN 0 ELSE signing_fee_actual END), SUM(CASE ISNULL(signing_fee_profit) or signing_fee_profit='' WHEN 1 THEN 0 ELSE signing_fee_profit END),SUM(CASE ISNULL(activated_num_target) or activated_num_target='' WHEN 1 THEN 0 ELSE activated_num_target END), SUM(CASE ISNULL(activated_num_actual) or activated_num_actual='' WHEN 1 THEN 0 ELSE activated_num_actual END), SUM(CASE ISNULL(sbdh_activated_num_target) or sbdh_activated_num_target='' WHEN 1 THEN 0 ELSE sbdh_activated_num_target END), SUM(CASE ISNULL(sbdh_activated_num_actual) or sbdh_activated_num_actual='' WHEN 1 THEN 0 ELSE sbdh_activated_num_actual END), SUM(CASE ISNULL(activated_fee_target) or activated_fee_target='' WHEN 1 THEN 0 ELSE activated_fee_target END), SUM(CASE ISNULL(activated_fee_actual) or activated_fee_actual='' WHEN 1 THEN 0 ELSE activated_fee_actual END), SUM(CASE ISNULL(activated_fee_profit) or activated_fee_profit='' WHEN 1 THEN 0 ELSE activated_fee_profit END), SUM(CASE ISNULL(sbdh_old_added_num) or sbdh_old_added_num='' WHEN 1 THEN 0 ELSE sbdh_old_added_num END), SUM(CASE ISNULL(sbdh_reduce_num) or sbdh_reduce_num='' WHEN 1 THEN 0 ELSE sbdh_reduce_num END), SUM(CASE ISNULL(loss_customer_num) or loss_customer_num='' WHEN 1 THEN 0 ELSE loss_customer_num END), SUM(CASE ISNULL(loss_num) or loss_num='' WHEN 1 THEN 0 ELSE loss_num END) FROM `kpitable` where month=\"%s\"  and region=\"%s\"  GROUP BY %s", month, area, "week")

		return ExecuteSql(sqlStr)

	} else if area == "南区+北区" {

		//sqlStr = fmt.Sprintf("SELECT  week,SUM(CASE ISNULL(new_signed_num_target) or new_signed_num_target='' WHEN 1 THEN 0 ELSE new_signed_num_target END), SUM(CASE ISNULL(new_signed_num_actual) or new_signed_num_actual='' WHEN 1 THEN 0 ELSE new_signed_num_actual END), SUM(CASE ISNULL(new_sbdh_num_target) or new_sbdh_num_target='' WHEN 1 THEN 0 ELSE new_sbdh_num_target END), SUM(CASE ISNULL(new_sbdh_num_actual) or new_sbdh_num_actual='' WHEN 1 THEN 0 ELSE new_sbdh_num_actual END), SUM(CASE ISNULL(signing_fee_target) or signing_fee_target='' WHEN 1 THEN 0 ELSE signing_fee_target END),SUM(CASE ISNULL(signing_fee_actual) or signing_fee_actual='' WHEN 1 THEN 0 ELSE signing_fee_actual END), SUM(CASE ISNULL(activated_num_target) or activated_num_target='' WHEN 1 THEN 0 ELSE activated_num_target END), SUM(CASE ISNULL(activated_num_actual) or activated_num_actual='' WHEN 1 THEN 0 ELSE activated_num_actual END), SUM(CASE ISNULL(sbdh_activated_num_target) or sbdh_activated_num_target='' WHEN 1 THEN 0 ELSE sbdh_activated_num_target END), SUM(CASE ISNULL(sbdh_activated_num_actual) or sbdh_activated_num_actual='' WHEN 1 THEN 0 ELSE sbdh_activated_num_actual END), SUM(CASE ISNULL(activated_fee_target) or activated_fee_target='' WHEN 1 THEN 0 ELSE activated_fee_target END), SUM(CASE ISNULL(activated_fee_actual) or activated_fee_actual='' WHEN 1 THEN 0 ELSE activated_fee_actual END),SUM(CASE ISNULL(sbdh_old_added_num) or sbdh_old_added_num='' WHEN 1 THEN 0 ELSE sbdh_old_added_num END), SUM(CASE ISNULL(sbdh_reduce_num) or sbdh_reduce_num='' WHEN 1 THEN 0 ELSE sbdh_reduce_num END), SUM(CASE ISNULL(loss_customer_num) or loss_customer_num='' WHEN 1 THEN 0 ELSE loss_customer_num END), SUM(CASE ISNULL(loss_num) or loss_num='' WHEN 1 THEN 0 ELSE loss_num END) FROM `kpitable` where month=\"%s\"   GROUP BY %s", month, "week")
		sqlStr = fmt.Sprintf("SELECT  week,SUM(CASE ISNULL(new_signed_num_target) or new_signed_num_target='' WHEN 1 THEN 0 ELSE new_signed_num_target END), SUM(CASE ISNULL(new_signed_num_actual) or new_signed_num_actual='' WHEN 1 THEN 0 ELSE new_signed_num_actual END), SUM(CASE ISNULL(new_sbdh_num_target) or new_sbdh_num_target='' WHEN 1 THEN 0 ELSE new_sbdh_num_target END), SUM(CASE ISNULL(new_sbdh_num_actual) or new_sbdh_num_actual='' WHEN 1 THEN 0 ELSE new_sbdh_num_actual END), SUM(CASE ISNULL(signing_fee_target) or signing_fee_target='' WHEN 1 THEN 0 ELSE signing_fee_target END),SUM(CASE ISNULL(signing_fee_actual) or signing_fee_actual='' WHEN 1 THEN 0 ELSE signing_fee_actual END), SUM(CASE ISNULL(signing_fee_profit) or signing_fee_profit='' WHEN 1 THEN 0 ELSE signing_fee_profit END),SUM(CASE ISNULL(activated_num_target) or activated_num_target='' WHEN 1 THEN 0 ELSE activated_num_target END), SUM(CASE ISNULL(activated_num_actual) or activated_num_actual='' WHEN 1 THEN 0 ELSE activated_num_actual END), SUM(CASE ISNULL(sbdh_activated_num_target) or sbdh_activated_num_target='' WHEN 1 THEN 0 ELSE sbdh_activated_num_target END), SUM(CASE ISNULL(sbdh_activated_num_actual) or sbdh_activated_num_actual='' WHEN 1 THEN 0 ELSE sbdh_activated_num_actual END), SUM(CASE ISNULL(activated_fee_target) or activated_fee_target='' WHEN 1 THEN 0 ELSE activated_fee_target END), SUM(CASE ISNULL(activated_fee_actual) or activated_fee_actual='' WHEN 1 THEN 0 ELSE activated_fee_actual END), SUM(CASE ISNULL(activated_fee_profit) or activated_fee_profit='' WHEN 1 THEN 0 ELSE activated_fee_profit END), SUM(CASE ISNULL(sbdh_old_added_num) or sbdh_old_added_num='' WHEN 1 THEN 0 ELSE sbdh_old_added_num END), SUM(CASE ISNULL(sbdh_reduce_num) or sbdh_reduce_num='' WHEN 1 THEN 0 ELSE sbdh_reduce_num END), SUM(CASE ISNULL(loss_customer_num) or loss_customer_num='' WHEN 1 THEN 0 ELSE loss_customer_num END), SUM(CASE ISNULL(loss_num) or loss_num='' WHEN 1 THEN 0 ELSE loss_num END)  FROM `kpitable` where month=\"%s\"   GROUP BY %s", month, "week")

		return ExecuteSql(sqlStr)
	} else {
		// 查询城市 每月每周 的数据
		var RowObjList []RowData
		for _, week := range []string{"1周", "2周", "3周", "4周", "5周"} {

			//sqlStr = fmt.Sprintf("SELECT  week, CASE ISNULL(new_signed_num_target) or new_signed_num_target='' WHEN 1 THEN 0 ELSE new_signed_num_target END, CASE ISNULL(new_signed_num_actual) or new_signed_num_actual='' WHEN 1 THEN 0 ELSE new_signed_num_actual END, CASE ISNULL(new_sbdh_num_target) or new_sbdh_num_target='' WHEN 1 THEN 0 ELSE new_sbdh_num_target END, CASE ISNULL(new_sbdh_num_actual) or new_sbdh_num_actual='' WHEN 1 THEN 0 ELSE new_sbdh_num_actual END, CASE ISNULL(signing_fee_target) or signing_fee_target='' WHEN 1 THEN 0 ELSE signing_fee_target END,CASE ISNULL(signing_fee_actual) or signing_fee_actual='' WHEN 1 THEN 0 ELSE signing_fee_actual END, CASE ISNULL(activated_num_target) or activated_num_target='' WHEN 1 THEN 0 ELSE activated_num_target END, CASE ISNULL(activated_num_actual) or activated_num_actual='' WHEN 1 THEN 0 ELSE activated_num_actual END, CASE ISNULL(sbdh_activated_num_target) or sbdh_activated_num_target='' WHEN 1 THEN 0 ELSE sbdh_activated_num_target END,CASE ISNULL(sbdh_activated_num_actual) or sbdh_activated_num_actual='' WHEN 1 THEN 0 ELSE sbdh_activated_num_actual END, CASE ISNULL(activated_fee_target) or activated_fee_target='' WHEN 1 THEN 0 ELSE activated_fee_target END, CASE ISNULL(activated_fee_actual) or activated_fee_actual='' WHEN 1 THEN 0 ELSE activated_fee_actual END,CASE ISNULL(sbdh_old_added_num) or sbdh_old_added_num='' WHEN 1 THEN 0 ELSE sbdh_old_added_num END, CASE ISNULL(sbdh_reduce_num) or sbdh_reduce_num='' WHEN 1 THEN 0 ELSE sbdh_reduce_num END, CASE ISNULL(loss_customer_num) or loss_customer_num='' WHEN 1 THEN 0 ELSE loss_customer_num END, CASE ISNULL(loss_num) or loss_num='' WHEN 1 THEN 0 ELSE loss_num END FROM `kpitable` where month='%s'  and branch='%s' and week='%s'", month, area, week)
			sqlStr = fmt.Sprintf("SELECT  week,SUM(CASE ISNULL(new_signed_num_target) or new_signed_num_target='' WHEN 1 THEN 0 ELSE new_signed_num_target END), SUM(CASE ISNULL(new_signed_num_actual) or new_signed_num_actual='' WHEN 1 THEN 0 ELSE new_signed_num_actual END), SUM(CASE ISNULL(new_sbdh_num_target) or new_sbdh_num_target='' WHEN 1 THEN 0 ELSE new_sbdh_num_target END), SUM(CASE ISNULL(new_sbdh_num_actual) or new_sbdh_num_actual='' WHEN 1 THEN 0 ELSE new_sbdh_num_actual END), SUM(CASE ISNULL(signing_fee_target) or signing_fee_target='' WHEN 1 THEN 0 ELSE signing_fee_target END),SUM(CASE ISNULL(signing_fee_actual) or signing_fee_actual='' WHEN 1 THEN 0 ELSE signing_fee_actual END), SUM(CASE ISNULL(signing_fee_profit) or signing_fee_profit='' WHEN 1 THEN 0 ELSE signing_fee_profit END),SUM(CASE ISNULL(activated_num_target) or activated_num_target='' WHEN 1 THEN 0 ELSE activated_num_target END), SUM(CASE ISNULL(activated_num_actual) or activated_num_actual='' WHEN 1 THEN 0 ELSE activated_num_actual END), SUM(CASE ISNULL(sbdh_activated_num_target) or sbdh_activated_num_target='' WHEN 1 THEN 0 ELSE sbdh_activated_num_target END), SUM(CASE ISNULL(sbdh_activated_num_actual) or sbdh_activated_num_actual='' WHEN 1 THEN 0 ELSE sbdh_activated_num_actual END), SUM(CASE ISNULL(activated_fee_target) or activated_fee_target='' WHEN 1 THEN 0 ELSE activated_fee_target END), SUM(CASE ISNULL(activated_fee_actual) or activated_fee_actual='' WHEN 1 THEN 0 ELSE activated_fee_actual END), SUM(CASE ISNULL(activated_fee_profit) or activated_fee_profit='' WHEN 1 THEN 0 ELSE activated_fee_profit END), SUM(CASE ISNULL(sbdh_old_added_num) or sbdh_old_added_num='' WHEN 1 THEN 0 ELSE sbdh_old_added_num END), SUM(CASE ISNULL(sbdh_reduce_num) or sbdh_reduce_num='' WHEN 1 THEN 0 ELSE sbdh_reduce_num END), SUM(CASE ISNULL(loss_customer_num) or loss_customer_num='' WHEN 1 THEN 0 ELSE loss_customer_num END), SUM(CASE ISNULL(loss_num) or loss_num='' WHEN 1 THEN 0 ELSE loss_num END)  FROM `kpitable` where month='%s'  and branch='%s' and week='%s'", month, area, week)
			RowObjList = append(RowObjList, ExecuteSql(sqlStr)...)

		}
		return RowObjList
	}

}

func ExecuteSql(sqlStr string) []RowData {

	// and region=\"%s\"
	rows, err := db.Query(sqlStr)
	if err != nil {

		log.Printf("query failed, err:%v\n", err)
		//return nil
		panic("业绩管理查询出错")
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()
	var RowObjList []RowData
	//fmt.Println(rows)

	// 循环读取结果集中的数据
	for rows.Next() {
		var rowData RowData
		//c, d, f, g, i, j, l, m, o, p, r, s, w, x, y, z
		err := rows.Scan(&rowData.Week, &rowData.C, &rowData.D, &rowData.F, &rowData.G, &rowData.I, &rowData.J, &rowData.K,
			&rowData.M, &rowData.N, &rowData.P, &rowData.Q, &rowData.S, &rowData.T, &rowData.U, &rowData.Y, &rowData.Z, &rowData.AA, &rowData.AB)

		//err := rows.Scan(&rowData.Week, &rowData.C)
		if err != nil {
			log.Printf("scan failed, err:%v\n", err)
			//return nil
			//fmt.Println("打印错误", err)
			panic("业绩管理查询出错")
		}
		//fmt.Println(month, rowData.Week, rowData.C)
		RowObjList = append(RowObjList, rowData)
		//fmt.Println(rowData.c, rowData.d, rowData.f, rowData.g, rowData.i, rowData.j,
		//	rowData.l, rowData.m, rowData.o, rowData.p, rowData.r, rowData.s, rowData.w, rowData.x, rowData.y, rowData.z)
	}
	return RowObjList

}

// 查询客户商机池管理
func QueryCustom(month, area string) []RowData {
	err := InitDB()
	if err != nil {
		log.Println(err)
		panic("初始化客户商机数据库连接出错")
	}

	var sqlStr string
	if strings.Contains("南区北区", area) {

		sqlStr = fmt.Sprintf("SELECT  week,SUM(CASE ISNULL(A_target_month) or A_target_month='' WHEN 1 THEN 0 ELSE A_target_month END), "+
			"SUM(CASE ISNULL(A_actual_week) or A_actual_week='' WHEN 1 THEN 0 ELSE A_actual_week END), SUM( CASE ISNULL(B_target_month) or B_target_month='' WHEN 1 THEN 0 ELSE B_target_month END), "+
			"SUM(CASE ISNULL(B_actual_week) or B_actual_week='' WHEN 1 THEN 0 ELSE B_actual_week END), SUM(CASE ISNULL(C_target_month) or C_target_month='' WHEN 1 THEN 0 ELSE C_target_month END),"+
			"SUM(CASE ISNULL(C_actual_week) or C_actual_week='' WHEN 1 THEN 0 ELSE C_actual_week END), SUM(CASE ISNULL(oppo_num) or oppo_num='' WHEN 1 THEN 0 ELSE oppo_num END), SUM(CASE ISNULL(staff_num) or staff_num='' WHEN 1 THEN 0 ELSE staff_num END), "+
			"SUM(CASE ISNULL(saler_num) or saler_num='' WHEN 1 THEN 0 ELSE saler_num END), SUM(CASE ISNULL(manager_num) or manager_num='' WHEN 1 THEN 0 ELSE manager_num END) "+
			"FROM customer_opp_management where month=\"%s\" and region=\"%s\" GROUP BY %s", month, area, "week")

		return executeSqlCustom(sqlStr)

	} else if area == "南区+北区" {

		sqlStr = fmt.Sprintf("SELECT  week,SUM(CASE ISNULL(A_target_month) or A_target_month='' WHEN 1 THEN 0 ELSE A_target_month END), "+
			"SUM(CASE ISNULL(A_actual_week) or A_actual_week='' WHEN 1 THEN 0 ELSE A_actual_week END), SUM( CASE ISNULL(B_target_month) or B_target_month='' WHEN 1 THEN 0 ELSE B_target_month END), "+
			"SUM(CASE ISNULL(B_actual_week) or B_actual_week='' WHEN 1 THEN 0 ELSE B_actual_week END), SUM(CASE ISNULL(C_target_month) or C_target_month='' WHEN 1 THEN 0 ELSE C_target_month END),"+
			"SUM(CASE ISNULL(C_actual_week) or C_actual_week='' WHEN 1 THEN 0 ELSE C_actual_week END), SUM(CASE ISNULL(oppo_num) or oppo_num='' WHEN 1 THEN 0 ELSE oppo_num END), SUM(CASE ISNULL(staff_num) or staff_num='' WHEN 1 THEN 0 ELSE staff_num END), "+
			"SUM(CASE ISNULL(saler_num) or saler_num='' WHEN 1 THEN 0 ELSE saler_num END), SUM(CASE ISNULL(manager_num) or manager_num='' WHEN 1 THEN 0 ELSE manager_num END) "+
			"FROM customer_opp_management where month=\"%s\"  GROUP BY %s", month, "week")
		return executeSqlCustom(sqlStr)

	} else {
		// 查询 每月 每周
		var RowObjList []RowData
		for _, week := range []string{"1周", "2周", "3周", "4周", "5周"} {

			sqlStr = fmt.Sprintf("SELECT  week,CASE ISNULL(A_target_month) or A_target_month='' WHEN 1 THEN 0 ELSE A_target_month END, CASE ISNULL(A_actual_week) or A_actual_week='' WHEN 1 THEN 0 ELSE A_actual_week END, CASE ISNULL(B_target_month) or B_target_month='' WHEN 1 THEN 0 ELSE B_target_month END, CASE ISNULL(B_actual_week) or B_actual_week='' WHEN 1 THEN 0 ELSE B_actual_week END, CASE ISNULL(C_target_month) or C_target_month='' WHEN 1 THEN 0 ELSE C_target_month END,CASE ISNULL(C_actual_week) or C_actual_week='' WHEN 1 THEN 0 ELSE C_actual_week END, CASE ISNULL(oppo_num) or oppo_num='' WHEN 1 THEN 0 ELSE oppo_num END, CASE ISNULL(staff_num) or staff_num='' WHEN 1 THEN 0 ELSE staff_num END, CASE ISNULL(saler_num) or saler_num='' WHEN 1 THEN 0 ELSE saler_num END, CASE ISNULL(manager_num) or manager_num='' WHEN 1 THEN 0 ELSE manager_num END FROM customer_opp_management where month='%s' and branch='%s' and week='%s'", month, area, week)
			RowObjList = append(RowObjList, executeSqlCustom(sqlStr)...)
		}

		return RowObjList

	}

}

func executeSqlCustom(sqlStr string) []RowData {

	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Printf("query failed, err:%v\n", err)
		panic("客户商机查询出错")
		//return nil
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()
	var RowObjList []RowData
	//fmt.Println(rows)

	// 循环读取结果集中的数据
	for rows.Next() {
		var rowData RowData
		//c, d, f, g, i, j, l, m, o, p, r, s, w, x, y, z
		err := rows.Scan(&rowData.Week, &rowData.C, &rowData.D, &rowData.F, &rowData.G, &rowData.I, &rowData.J,
			&rowData.L, &rowData.M, &rowData.O, &rowData.P)

		if err != nil {
			log.Println(err)
			panic("客户商机查询出错")
			//fmt.Printf("scan failed, err:%v\n", err)
			//return nil
		}

		RowObjList = append(RowObjList, rowData)

	}
	return RowObjList

}

// 新签客户成交
func QueryNewCustom(month, area string) []RowData {
	err := InitDB()
	if err != nil {
		log.Println(err)
		panic("新客户成交初始化数据库连接出错")
	}

	var sqlStr string
	if strings.Contains("南区北区", area) {

		sqlStr = fmt.Sprintf(" SELECT  week,SUM(CASE ISNULL(tel_dev) or tel_dev='' WHEN 1 THEN 0 ELSE tel_dev END), SUM(CASE ISNULL(transfer_introduce) or transfer_introduce='' WHEN 1 THEN 0 ELSE transfer_introduce END), SUM(CASE ISNULL(market_activities) or market_activities='' WHEN 1 THEN 0 ELSE market_activities END), SUM(CASE ISNULL(400_customer_service) or 400_customer_service='' WHEN 1 THEN 0 ELSE 400_customer_service END), SUM(CASE ISNULL(community) or community='' WHEN 1 THEN 0 ELSE community END) FROM `new_signed_deal` where month=\"%s\"  and region=\"%s\"  GROUP BY %s", month, area, "week")
		return executeSqlNewCustom(sqlStr)

	} else if area == "南区+北区" {

		sqlStr = fmt.Sprintf(" SELECT  week,SUM(CASE ISNULL(tel_dev) or tel_dev='' WHEN 1 THEN 0 ELSE tel_dev END), SUM(CASE ISNULL(transfer_introduce) or transfer_introduce='' WHEN 1 THEN 0 ELSE transfer_introduce END), SUM(CASE ISNULL(market_activities) or market_activities='' WHEN 1 THEN 0 ELSE market_activities END), SUM(CASE ISNULL(400_customer_service) or 400_customer_service='' WHEN 1 THEN 0 ELSE 400_customer_service END), SUM(CASE ISNULL(community) or community='' WHEN 1 THEN 0 ELSE community END) FROM `new_signed_deal` where month=\"%s\"   GROUP BY %s", month, "week")
		return executeSqlNewCustom(sqlStr)
	} else {

		var RowObjList []RowData
		for _, week := range []string{"1周", "2周", "3周", "4周", "5周"} {

			sqlStr = fmt.Sprintf("SELECT  week,CASE ISNULL(tel_dev) or tel_dev='' WHEN 1 THEN 0 ELSE tel_dev END, CASE ISNULL(transfer_introduce) or transfer_introduce='' WHEN 1 THEN 0 ELSE transfer_introduce END, CASE ISNULL(market_activities) or market_activities='' WHEN 1 THEN 0 ELSE market_activities END, CASE ISNULL(400_customer_service) or 400_customer_service='' WHEN 1 THEN 0 ELSE 400_customer_service END, CASE ISNULL(community) or community='' WHEN 1 THEN 0 ELSE community END FROM `new_signed_deal` where month='%s' and branch='%s' and week='%s'", month, area, week)
			RowObjList = append(RowObjList, executeSqlNewCustom(sqlStr)...)
		}
		return RowObjList

	}

}

func executeSqlNewCustom(sqlStr string) []RowData {

	// and region=\"%s\"
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Printf("query failed, err:%v\n", err)
		panic("新客户成交查询出错")
		//return nil
	}
	defer rows.Close()
	var RowObjList []RowData
	//fmt.Println(rows)

	// 循环读取结果集中的数据
	for rows.Next() {
		var rowData RowData
		//c, d, f, g, i, j, l, m, o, p, r, s, w, x, y, z
		err := rows.Scan(&rowData.Week, &rowData.C, &rowData.D, &rowData.E, &rowData.F, &rowData.G)

		//err := rows.Scan(&rowData.Week, &rowData.C)
		if err != nil {
			log.Printf("scan failed, err:%v\n", err)
			panic("新客户成交查询出错")
			//return nil
		}

		RowObjList = append(RowObjList, rowData)

	}
	return RowObjList

}

func QueryTest(month, area string) []RowData {
	err := InitDB()
	if err != nil {
		log.Println(err)
		panic("初始化客户商机数据库连接出错")
	}

	var sqlStr string
	//if strings.Contains("南区北区", area) {
	//
	//	sqlStr = fmt.Sprintf("SELECT  t.week,SUM(`a_target`), SUM(`A_actual_week`), SUM(`b_target`), SUM(`B_actual_week`), " +
	//		"SUM(`c_target`),SUM(`C_actual_week`), SUM(`oppo_num`), SUM(`staff_num`), SUM(`saler_num`), SUM(`manager_num`) " +
	//		" FROM (SELECT *, IFNULL(A_target_month,0) a_target, IFNULL(B_target_month,0) b_target, IFNULL(C_target_month,0) " +
	//		"c_target FROM customer_opp_management) t where month=\"%s\" and region=\"%s\"  GROUP BY  %s", month, area, "week")
	//
	//} else {
	//
	//	sqlStr = fmt.Sprintf("SELECT  t.week,SUM(`a_target`), SUM(`A_actual_week`), SUM(`b_target`), SUM(`B_actual_week`), " +
	//		"SUM(`c_target`),SUM(`C_actual_week`), SUM(`oppo_num`), SUM(`staff_num`), SUM(`saler_num`), SUM(`manager_num`) " +
	//		" FROM (SELECT *, IFNULL(A_target_month,0) a_target, IFNULL(B_target_month,0) b_target, IFNULL(C_target_month,0) " +
	//		"c_target FROM customer_opp_management) t where month=\"%s\"   GROUP BY  %s", month, "week")
	//}

	sqlStr = fmt.Sprintf("SELECT  week, sum(CASE ISNULL(A_target_month) or A_target_month=\"\" WHEN 1 THEN 0 ELSE A_target_month END ) FROM customer_opp_management where month=\"%s\"   GROUP BY  %s", month, "week")
	//sqlStr = fmt.Sprintf("SELECT  week, sum(A_target_month)  FROM customer_opp_management where month=\"%s\"   GROUP BY  %s", month, "week")
	fmt.Println(sqlStr)
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Printf("query failed, err:%v\n", err)
		fmt.Println(err)
		//panic("客户商机查询出错")
		//return nil
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()
	var RowObjList []RowData
	//fmt.Println(rows)

	// 循环读取结果集中的数据
	for rows.Next() {
		var rowData RowData
		//c, d, f, g, i, j, l, m, o, p, r, s, w, x, y, z
		err := rows.Scan(&rowData.Week, &rowData.C)

		if err != nil {
			log.Println(err)
			fmt.Println(err)
			panic("客户商机查询出错")
			//fmt.Printf("scan failed, err:%v\n", err)
			//return nil
		}
		fmt.Println(rowData.C)
		RowObjList = append(RowObjList, rowData)

	}
	return RowObjList
}

func QueryComplete(area string) []RowData {
	err := InitDB()
	if err != nil {
		log.Println(err)
		panic("新客户成交初始化数据库连接出错")
	}

	var sqlStr string

	if strings.Contains(area, "北区") || strings.Contains(area, "南区") {
		sqlStr = fmt.Sprintf(" SELECT  region,SUM(CASE ISNULL(num) or num='' WHEN 1 THEN 0 ELSE num END)  FROM `completion_rate_management`   GROUP BY region")

	} else {

		sqlStr = fmt.Sprintf(" SELECT  branch,CASE ISNULL(num) or num='' WHEN 1 THEN 0 ELSE num END  FROM `completion_rate_management`   where branch='%s'", area)

	}

	//fmt.Println(sqlStr)
	// and region=\"%s\"
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Printf("query failed, err:%v\n", err)
		fmt.Println(err)
		panic("完成率查询出错")
		//return nil
	}
	defer rows.Close()
	var RowObjList []RowData
	//fmt.Println(rows)

	// 循环读取结果集中的数据
	for rows.Next() {
		var rowData RowData
		//c, d, f, g, i, j, l, m, o, p, r, s, w, x, y, z
		err := rows.Scan(&rowData.Region, &rowData.Num)

		//err := rows.Scan(&rowData.Week, &rowData.C)
		if err != nil {
			log.Printf("scan failed, err:%v\n", err)
			panic("新客户成交查询出错")
			//return nil
		}
		//fmt.Println(rowData.Region, rowData.Num)
		RowObjList = append(RowObjList, rowData)

	}
	return RowObjList
}

type AreaName struct {
	Name string
	Date string
}

func QueryCity() []AreaName {
	err := InitDB()
	if err != nil {
		log.Println(err)
		panic("查询城市初始化数据库连接出错")
	}

	var sqlStr string
	sqlStr = "SELECT branch, DATE_FORMAT(max(update_time), '%Y-%m-%d') FROM `kpitable` GROUP BY branch"
	//if area == "南区+北区" {
	//	sqlStr = fmt.Sprintf("select branch from kpitable where month = '%s' and week = '%s'  ", month, week)
	//} else {
	//
	//	sqlStr = fmt.Sprintf("select branch from kpitable where month = '%s' and week = '%s'  and region = '%s'", month, week, area)
	//}

	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Printf("query failed, err:%v\n", err)
		fmt.Println(err)
		panic("地区查询出错")
		//return nil
	}
	defer rows.Close()
	var AreaNameObjList []AreaName
	//fmt.Println(rows)

	// 循环读取结果集中的数据
	for rows.Next() {
		var areaname AreaName
		//c, d, f, g, i, j, l, m, o, p, r, s, w, x, y, z
		err := rows.Scan(&areaname.Name, &areaname.Date)

		//err := rows.Scan(&rowData.Week, &rowData.C)
		if err != nil {
			log.Printf("scan failed, err:%v\n", err)
			panic("新客户成交查询出错")
			//return nil
		}
		//fmt.Println(rowData.Region, rowData.Num)
		AreaNameObjList = append(AreaNameObjList, areaname)
	}

	return AreaNameObjList
}

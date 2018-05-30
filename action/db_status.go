package action

import(
	"ormtest/base"
	"database/sql"
	"strings"
	"reflect"
	"fmt"
	"strconv"
	"github.com/garyburd/redigo/redis"
	"time"
)

//监控基础的点
type MySQLInfo struct {
	Bytes_received          int64
	Bytes_sent              int64
	Com_commit              int64
	Com_delete              int64
	Com_select              int64
	Com_rollback            int64
	Com_update              int64
	Com_insert              int64
	Com_execute_sql         int64
	Created_tmp_disk_tables int64
	Slow_queries            int64
	Threads_connected       int64
	Threads_running         int64
	Questions				int64
	Created_tmp_tables		int64
	Com_lock_tables			int64
}

//监控差值
type Phase struct {
	Com_commit              int64
	Com_delete              int64
	Com_select              int64
	Com_rollback            int64
	Com_update              int64
	Com_insert              int64
	Com_execute_sql         int64
	Bytes_received          int64
	Questions				int64
	Created_tmp_tables		int64
	Com_lock_tables			int64
}


//将数据存入redis
func pre_phase(cdbname string,obj string,val int64) int64{
	phase := &Phase{}
	xxx := new(base.Redis)
	xxx.Host="127.0.0.1"
	xxx.Port="6379"
	Rediscon := xxx.Connect(cdbname)
	pha := reflect.ValueOf(phase)
	field := pha.Elem().FieldByName(strings.Title(strings.ToLower(obj)))
	if field.IsValid(){
		cdbname = cdbname+"&"+ obj
		is_key_exit := base.Redis_Exists(Rediscon,cdbname)
		if is_key_exit{
			old_val,err := redis.Int64(Rediscon.Do("GET",cdbname))
			if err != nil {
				base.Logger.Warning(err)
			}
			err = base.Redis_set(Rediscon,cdbname,val)
			val = val - old_val
			if err != nil{
				return 0
			}
			return val/60
		}
		err := base.Redis_set(Rediscon,cdbname,val)
		if err != nil {
			base.Logger.Warning(err)
			return 0
		}
		return val
	}
	return 0
}

func Exec(db *sql.DB,cname string,channel chan map[string]int64,bool bool){
	mysqlinfo := &MySQLInfo{}
	ref := reflect.ValueOf(mysqlinfo)
	defer db.Close()
	temp := make(map[string]int64)
	var name string
	var value string
	cur_status, err := db.Query("show global status;")
	if err != nil {
		base.Logger.Warning(err,cname)
		temp = nil
	}else {
		//如果select执行成功，需要执行关闭的操作
		defer cur_status.Close()
	}
	//这里是sql执行没问题后，取相应的值
	if err == nil {
		for cur_status.Next() {
			cur_status.Scan(&name, &value)
			//根据struct取相应的值
			field := ref.Elem().FieldByName(strings.Title(strings.ToLower(name)))
			if field.IsValid() {
				switch field.Kind() {
				case reflect.Int64:
					name1 := cname + "&" + name
					val, _ := strconv.ParseInt(value, 10, 64)
					val = pre_phase(cname, name, val)
					temp[name1] = val
					field.SetInt(val)
				case reflect.String:
					field.SetString(value)
				}
			}
		}
	}
	channel <- temp
	if bool{
		//到最后一个cdb，停留一秒时间，并且关闭channel
		time.Sleep(1 * time.Second)
		close(channel)
	}
}


func Status(){
	channel := make(chan map[string]int64)
	i := 1
	flag := false
	length := len(base.Cdb)
	for cname,ip := range base.Cdb{
		if i == length{
			flag = true
		}else{
			i ++
		}
		db,err := base.JDBC(ip)
		if err != nil {
			base.Logger.Warning(err)
		}
		go Exec(db,cname,channel,flag)
	}
	for {
		temp,open := <- channel
		//如果没有值了，停止循环
		if !open{
			break
		}
		for cdbinfo,value := range temp{
			cdb := strings.Split(cdbinfo,"&")
			cdbname,obj :=cdb[0],cdb[1]
			fmt.Println(cdbname,obj,value)
		}
	}
}

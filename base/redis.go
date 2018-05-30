package base
import (
	"github.com/garyburd/redigo/redis"
)

type Redis struct {
	Host string
	Port string
	Passwd string
}

func (r *Redis) Connect(cname string) (redis.Conn){
	c,err := redis.Dial("tcp",r.Host+":"+r.Port)
	if err != nil {
		Logger.Error(err,cname)
		return nil
	}
	return c
}

func Redis_set(conn redis.Conn,cdbname string,val int64) error{
	_, err := conn.Do("SET",cdbname,val)
	return err
}

func Redis_Exists(conn redis.Conn,cdbname string) (bool){
	is_key_exists,err := redis.Bool(conn.Do("EXISTS",cdbname))
	if err != nil {
		Logger.Warning(err)
	}
	return is_key_exists

}





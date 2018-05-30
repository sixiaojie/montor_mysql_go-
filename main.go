package main

import(
	"ormtest/action"
	"os"
	"fmt"
	"ormtest/zabbix"
)

func main() {
	//action.Status()
	args := os.Args
	if len(args) != 2{
		fmt.Printf("%s 格式为:%s find/status\n",args[0],args[0])
		os.Exit(1)
	}else{
		if args[1] == "status"{
			action.Status()
		}else if args[1] == "find"{
			zabbix.Discovery_db()
		}
	}
}

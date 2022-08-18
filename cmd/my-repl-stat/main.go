package main

import (
	"flag"
	"fmt"
	"os/exec"
	"runtime"
)

// variables declaration
var dbEndPoint string


func execute() {
	//fmt.Printf("dbEndPoint = %v\n", dbEndPoint)
	bashCommand := fmt.Sprintf(`mysql -h %v -e "show slave status\G"| egrep "Slave_IO_Running:|Slave_SQL_Running:|Seconds_Behind_Master:|Master_Host:" | grep -v Warning`, dbEndPoint)
	//fmt.Printf("Execing command:\n%v\n\n", bashCommand)
	out, err := exec.Command("bash", "-c", bashCommand).Output()
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println("Command Successfully Executed")
	fmt.Println(string(out))
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
		return
	}

	// flags declaration using flag package
	flag.StringVar(&dbEndPoint, "h", "", "Specify your AWS EndPoint.")
	flag.Parse() // after declaring flags we need to call it

	if dbEndPoint != "" {
		execute()
		return
	}
	flag.PrintDefaults()
}


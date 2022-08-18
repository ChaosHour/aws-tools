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
	bashCommand := fmt.Sprintf(`mysql -h %v -t -e "SELECT SUBSTRING_INDEX(host, ':', 1) AS host_short, GROUP_CONCAT(DISTINCT USER) AS users, COUNT(*) FROM information_schema.processlist GROUP BY host_short UNION ALL SELECT 'total', '', count(*) as TOTAL FROM information_schema.processlist ORDER BY 3, 2"`, dbEndPoint)
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


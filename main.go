package main

import (
	"flag"
	"fmt"
	"os/exec"
	"runtime"
)

// variables declaration
var dbInstanceIdentifier string
var dbRegion string
var logType string

func cluster() {
	fmt.Printf("dbRegion = %v\n", dbRegion)
	bashCommand := fmt.Sprintf(`aws rds describe-db-instances --region %v | jq '.DBInstances[].DBInstanceIdentifier'`, dbRegion)
	fmt.Printf("Execing command:\n%v\n\n", bashCommand)
	out, err := exec.Command("bash", "-c", bashCommand).Output()
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println("Command Successfully Executed")
	fmt.Println(string(out))
}

func execute() {
	fmt.Printf("dbInstanceIdentifier = %v\n", dbInstanceIdentifier)
	fmt.Printf("logType = %v\n", logType)
	bashCommand := fmt.Sprintf(`aws rds describe-db-log-files --db-instance-identifier %v | jq '.DescribeDBLogFiles[] | select(.LogFileName|test("%v.")) | .LogFileName'`, dbInstanceIdentifier, logType)
	fmt.Printf("Execing command:\n%v\n\n", bashCommand)
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
	flag.StringVar(&dbRegion, "r", "us-west-1", "Select a aws region to view your clusters.")
	flag.StringVar(&logType, "l", "error", "Log type to use Example: slow, error.")
	flag.StringVar(&dbInstanceIdentifier, "d", "", "Specify your db-instance-identifier to view it's logs.")
	flag.Parse() // after declaring flags we need to call it

	if dbInstanceIdentifier != "" {
		execute()
		return
	}

	if dbRegion != "" {
		cluster()
		return
	}

	flag.PrintDefaults()
}


package main

import (
	"flag"
	"fmt"
	"os/exec"
	"runtime"
)

// variables declaration
var dbEndPoint string
var dbUser string

func execute() {
	//fmt.Printf("dbEndPoint = %v\n", dbEndPoint)
	bashCommand := fmt.Sprintf(`psql -U %v -h %v -x -c "select usename, state, count(*) from pg_stat_activity  where pid <> pg_backend_pid() group by 1,2 order by 1" -d postgres`, dbUser, dbEndPoint)
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
	flag.StringVar(&dbUser, "u", "", "Specify your AWS User.")
	flag.Parse() // after declaring flags we need to call it

	// Make sure that the flags are set
	if dbEndPoint == "" || dbUser == "" {
		flag.PrintDefaults()
		return
	}
	execute()
}

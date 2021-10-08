package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
)

var (
	dbIdent  = flag.String("d", "", "db-cluster-identifier")
	dbRegion = flag.String("r", "us-west-1", "Select a AWS Region to view you Clusters")
)

func main() {
	flag.Parse()

	args := []string{
		"aws",
		"rds",
		"--region", fmt.Sprintf("%s", *dbRegion),
		"describe-db-cluster-endpoints",
		"--db-cluster-identifier", fmt.Sprintf("%s", *dbIdent),
	}

	cmd := exec.Command(args[0], args[1:]...)

	fmt.Printf("Executing command:\n%v\n\n", args)

	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Something went wrong, please review and try again.: %v", err)
	}
	fmt.Printf("%s\n", b)
}

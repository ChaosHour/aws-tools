// aws cli wrapper for go to interact with aws services
//

package main

// import packages
import (
	"flag"
	"fmt"
	"log"
	"os/exec"
)

// define vars
var (
	dbIdent  = flag.String("d", "", "db-cluster-identifier")
	dbRegion = flag.String("r", "us-west-1", "Select a AWS Region to view you Clusters")
)

// define functions
func main() {
	// parse flags
	flag.Parse()

	// aws rds --region $dbRegion describe-db-clusters --db-cluster-id $dbIdent --query "*[].['Cluster:',DBClusterIdentifier,DBClusterMembers[*].['Instance:',DBInstanceIdentifier,IsClusterWriter]]" --output text

	args := []string{
		"aws",
		"rds",
		"--region", fmt.Sprintf("%s", *dbRegion),
		"describe-db-clusters",
		"--db-cluster-id", fmt.Sprintf("%s", *dbIdent),
		"--query", "*[].['Cluster:',DBClusterIdentifier,DBClusterMembers[*].['Instance:',DBInstanceIdentifier,IsClusterWriter]]",
		"--output", "text",
	}

	cmd := exec.Command(args[0], args[1:]...)

	fmt.Printf("Executing command:\n%v\n\n", args)

	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Something went wrong, please review and try again.: %v", err)
	}
	fmt.Printf("%s\n", b)
}

// aws dns route53 resolve-ip

package main

// Imports
import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

// Variables
var (
	domain string
	zoneID string
)

// Functions
func main() {
	// Get the domain name
	domain = os.Args[1]

	// Get the zone ID
	zoneID = getZoneID()

	// Get the IP address
	ip := getIP()

	// Update the record
	updateRecord(ip)
}

// Get the zone ID
func getZoneID() string {
	// Create a new session
	sess := session.Must(session.NewSession())

	// Create a new Route53 client
	svc := route53.New(sess)

	// Get the hosted zone ID
	params := &route53.ListHostedZonesInput{}
	result, err := svc.ListHostedZones(params)
	if err != nil {
		panic(err)
	}

	// Loop through the hosted zones
	for _, zone := range result.HostedZones {
		// Check if the domain is in the hosted zone
		if strings.Contains(*zone.Name, domain) {
			// Return the zone ID
			return *zone.Id
		}
	}

	// Return an error
	fmt.Println("Error: Domain not found")
	os.Exit(1)

	// Return an empty string
	return ""
}

// Get the IP address
func getIP() string {
	// Create a new session
	sess := session.Must(session.NewSession())

	// Create a new Route53 client
	svc := route53.New(sess)

	// Get the record
	params := &route53.ListResourceRecordSetsInput{
		HostedZoneId: aws.String(zoneID),
	}
	result, err := svc.ListResourceRecordSets(params)
	if err != nil {
		panic(err)
	}

	// Loop through the records
	for _, record := range result.ResourceRecordSets {
		// Check if the record is an A record
		if *record.Type == "A" {
			// Return the IP address
			return *record.ResourceRecords[0].Value
		}
	}

	// Return an error
	fmt.Println("Error: No A record found")
	os.Exit(1)

	// Return an empty string
	return ""
}

// Update the record
func updateRecord(ip string) {
	// Create a new session
	sess := session.Must(session.NewSession())

	// Create a new Route53 client
	svc := route53.New(sess)

	// Get the record
	params := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("UPSERT"),
					ResourceRecordSet: &route53.ResourceRecordSet{
						Name: aws.String(domain),
						Type: aws.String("A"),
						ResourceRecords: []*route53.ResourceRecord{
							{
								Value: aws.String(ip),
							},
						},
						TTL: aws.Int64(300),
					},
				},
			},
		},
		HostedZoneId: aws.String(zoneID),
	}
	result, err := svc.ChangeResourceRecordSets(params)
	if err != nil {
		panic(err)
	}

	// Print the change ID
	fmt.Println(*result.ChangeInfo.Id)
}

// EOF

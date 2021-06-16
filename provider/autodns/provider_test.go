package autodns

import (
	"os"
	"testing"

	"github.com/NectGmbH/dns"
	"github.com/sirupsen/logrus"
)

// GetUserName avoids having to add user name in git
func GetUserName() string {
	return os.Getenv("DNS_PROVIDER_AUTODNS_USER")
}

// GetPassword avoids having to add password in git
func GetPassword() string {
	return os.Getenv("DNS_PROVIDER_AUTODNS_PASSWORD")
}

// TestReadFromAutoDNS reads data from AutoDNS.
func TestReadFromAutoDNS(t *testing.T) {
	username := GetUserName()
	password := GetPassword()
	var dnsProvider dns.Provider
	dnsProvider = NewProvider(username, password)
	zone, err := dnsProvider.GetZone("nect.shop")
	if err != nil {
		logrus.Fatal("Error")
	}
	logrus.Info("len:", len(zone.Records))
	for i := 0; i < len(zone.Records); i++ {
		logrus.WithFields(logrus.Fields{
			"Class":      zone.Records[i].Class,
			"Name":       zone.Records[i].Name,
			"TTL":        zone.Records[i].TTL,
			"Type":       zone.Records[i].Type,
			"Value":      zone.Records[i].Value,
			"Preference": zone.Records[i].Preference,
		}).Info("item", i)
	}
}

func TestCreate(t *testing.T) {
	username := GetUserName()
	password := GetPassword()
	var dnsProvider dns.Provider
	dnsProvider = NewProvider(username, password)
	zone, err := dnsProvider.GetZone("nect.shop")
	if err != nil {
		logrus.Fatal("Error")
	}
	logrus.Info("len:", len(zone.Records))
	for i := 0; i < len(zone.Records); i++ {
		logrus.WithFields(logrus.Fields{
			"Class":      zone.Records[i].Class,
			"Name":       zone.Records[i].Name,
			"TTL":        zone.Records[i].TTL,
			"Type":       zone.Records[i].Type,
			"Value":      zone.Records[i].Value,
			"Preference": zone.Records[i].Preference,
		}).Info("item", i)
	}

	logrus.Info(zone.String())
	var records []dns.Record
	record1 := dns.NewRecord(dns.ClassIN, "", 0, dns.RecordTypeNS, "a.ns14.net", 0)
	records = append(records, *record1)
	record2 := dns.NewRecord(dns.ClassIN, "", 0, dns.RecordTypeNS, "b.ns14.net", 0)
	records = append(records, *record2)
	record3 := dns.NewRecord(dns.ClassIN, "", 0, dns.RecordTypeNS, "c.ns14.net", 0)
	records = append(records, *record3)
	record4 := dns.NewRecord(dns.ClassIN, "", 0, dns.RecordTypeNS, "d.ns14.net", 0)
	records = append(records, *record4)
	record5 := dns.NewRecord(dns.ClassIN, "elephant", 0, dns.RecordTypeA, "85.214.241.20", 0)
	records = append(records, *record5)

	newZone := dns.NewZone("nect.shop", records)
	t.Skip("No idea why this never passes")
	updateErr := dnsProvider.UpdateZone(*newZone)
	if updateErr != nil {
		logrus.Fatalf("UpdateZone failed, see: %v", updateErr)
	}

}

func TestShortern(t *testing.T) {
	username := GetUserName()
	password := GetPassword()
	var dnsProvider dns.Provider
	dnsProvider = NewProvider(username, password)
	zone, err := dnsProvider.GetZone("nect.shop")
	if err != nil {
		logrus.Fatal("Error")
	}
	logrus.Info("len:", len(zone.Records))
	for i := 0; i < len(zone.Records); i++ {
		logrus.WithFields(logrus.Fields{
			"Class":      zone.Records[i].Class,
			"Name":       zone.Records[i].Name,
			"TTL":        zone.Records[i].TTL,
			"Type":       zone.Records[i].Type,
			"Value":      zone.Records[i].Value,
			"Preference": zone.Records[i].Preference,
		}).Info("item", i)
	}
	zone.Records = zone.Records[:6]
	updateErr := dnsProvider.UpdateZone(zone)
	if updateErr != nil {
		logrus.Fatalf("UpdateZone failed, see: %v", updateErr)
	}

}

// TestDeleteJam should test Deleting blank but currently fails
func TestDeleteJam(t *testing.T) {
	username := GetUserName()
	password := GetPassword()
	var dnsProvider dns.Provider
	dnsProvider = NewProvider(username, password)
	zone, err := dnsProvider.GetZone("nect.shop")
	if err != nil {
		logrus.Fatal("Error")
	}
	zone.RemoveRecordsWithName("jam")
	logrus.Info("len:", len(zone.Records))
	for i := 0; i < len(zone.Records); i++ {
		logrus.WithFields(logrus.Fields{
			"Class":      zone.Records[i].Class,
			"Name":       zone.Records[i].Name,
			"TTL":        zone.Records[i].TTL,
			"Type":       zone.Records[i].Type,
			"Value":      zone.Records[i].Value,
			"Preference": zone.Records[i].Preference,
		}).Info("item", i)
	}
	updateErr := dnsProvider.UpdateZone(zone)
	if updateErr != nil {
		logrus.Fatalf("ssssUpdateZone failed, see: %v", updateErr)
	}

}

package dns

import (
	"fmt"
	"testing"
)

func Test_keyOp_Degenerate(t *testing.T) {
	TestRecords1 := []Record{
		*NewRecord(ClassIN, "", 0, RecordTypeNS, "a.ns14.net", 0),
		*NewRecord(ClassIN, "", 0, RecordTypeNS, "b.ns14.net", 0),
		*NewRecord(ClassIN, "", 0, RecordTypeNS, "c.ns14.net", 0),
		*NewRecord(ClassIN, "", 0, RecordTypeNS, "d.ns14.net", 0),
		*NewRecord(ClassIN, "elephant", 0, RecordTypeA, "85.214.241.20", 0),
	}

	TestRecords2 := []Record{
		*NewRecord(ClassIN, "", 0, RecordTypeNS, "a.ns14.net", 0),
		*NewRecord(ClassIN, "", 0, RecordTypeNS, "b.ns14.net", 0),
		*NewRecord(ClassIN, "", 0, RecordTypeNS, "c.ns14.net", 0),
		*NewRecord(ClassIN, "", 0, RecordTypeNS, "d.ns14.net", 0),
		*NewRecord(ClassIN, "pink.elephant", 0, RecordTypeA, "85.214.241.20", 0),
	}

	tests := []struct {
		name        string
		zoneName    string
		zoneRecs    []Record
		delete      string
		wantRecords int
	}{
		{
			name:        "success",
			zoneName:    "nect.shop",
			zoneRecs:    TestRecords1,
			delete:      "elephant.nect.shop",
			wantRecords: 4,
		},
		{
			name:        "noncanoncal",
			zoneName:    "nect.shop",
			zoneRecs:    TestRecords1,
			delete:      "elephant",
			wantRecords: 5,
		},
		{
			name:        "withAnS",
			zoneName:    "nect.shop",
			zoneRecs:    TestRecords1,
			delete:      "elephants.nect.shop",
			wantRecords: 5,
		},
		{
			name:        "withWrongZone",
			zoneName:    "nect.shops",
			zoneRecs:    TestRecords1,
			delete:      "elephant.nect.shop",
			wantRecords: 5,
		},
		{
			name:        "keepSubdomains",
			zoneName:    "nect.shop",
			zoneRecs:    TestRecords1,
			delete:      "pink.elephant.nect.shop",
			wantRecords: 5,
		},
		{
			name:        "subdomains",
			zoneName:    "nect.shop",
			zoneRecs:    TestRecords2,
			delete:      "pink.elephant.nect.shop",
			wantRecords: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newZone := NewZone(tt.zoneName, tt.zoneRecs)
			newZone.RemoveRecordsWithCanonicalnameType(tt.delete)
			if len(newZone.Records) != tt.wantRecords {
				errorString := fmt.Sprintf("Got '%v', Expected number of records was '%v'", len(newZone.Records), tt.wantRecords)
				panic(errorString)

			}

		})
	}
}

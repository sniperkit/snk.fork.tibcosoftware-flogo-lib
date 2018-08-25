/*
Sniperkit-Bot
- Status: analyzed
*/

package json

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sniperkit/snk.fork.tibcosoftware-flogo-lib/core/mapper/exprmapper/json/field"
)

var jsonData = `{
    "City": [
        {
            "Array": [
                {
                    "id": "11111"
                },
                {
                    "id": "2222"
                }
            ],
            "InUS": true,
            "Name": "Sugar Land",
            "Park": {
                "Emails": null,
                "Location": "location",
                "Maps": {
                    "bb": "bb",
                    "cc": "cc",
                    "dd": "dd"
                },
                "Name": "Name"
            }
        }
    ],
    "Emails": [
        "123@123.com",
        "456@456.com"
    ],
    "Id": 1234,
    "In": "string222",
    "Maps": {
        "bb": "bb",
        "cc": "cc",
        "dd": "dd"
    },
    "State": "Taxes",
    "Streat": "311 wind st",
    "ZipCode": "77477",
    "hello world":"CHINA",
    "tag  world":"CHINA"
}`

func TestRootChildArray(t *testing.T) {
	mappingField := &field.MappingField{HasArray: true, HasSpecialField: false}
	mappingField.Fields = []string{"City[0]"}
	value, err := GetFieldValueFromIn(jsonData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	log.Debug("Value is:", value)
}

func TestRoot(t *testing.T) {
	mappingField := &field.MappingField{HasArray: false, HasSpecialField: false}
	mappingField.Fields = []string{"City"}
	value, err := GetFieldValueFromIn(jsonData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	log.Debug("Value is:", value)
}

func TestGetFieldWithSpaces(t *testing.T) {
	mappingField := &field.MappingField{HasArray: false, HasSpecialField: false}
	mappingField.Fields = []string{"hello world"}
	value, err := GetFieldValueFromIn(jsonData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	log.Debug("Value is:", value)
}

func TestGetFieldWithTag(t *testing.T) {
	mappingField := &field.MappingField{HasArray: false, HasSpecialField: false}
	mappingField.Fields = []string{"tag  world"}
	value, err := GetFieldValueFromIn(jsonData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	log.Info("Value is:", value)
}

func TestGetArray(t *testing.T) {
	mappingField := &field.MappingField{HasArray: true, HasSpecialField: false}
	mappingField.Fields = []string{"Emails[0]"}
	value, err := GetFieldValueFromIn(jsonData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	log.Debug("Value is:", value)

}

func TestMultipleLevelArray(t *testing.T) {
	mappingField := &field.MappingField{HasArray: true, HasSpecialField: false}
	mappingField.Fields = []string{"City[0]", "Array[1]", "id"}
	value, err := GetFieldValueFromIn(jsonData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	log.Debug("Value:", value)
}

func TestMultipleLevelObject(t *testing.T) {
	mappingField := &field.MappingField{HasArray: true, HasSpecialField: false}
	mappingField.Fields = []string{"City[0]", "Park", "Maps", "bb"}
	value, err := GetFieldValueFromIn(jsonData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	log.Debug("Value:", value)
}

func TestID(t *testing.T) {
	mappingField := &field.MappingField{HasArray: false, HasSpecialField: false}
	mappingField.Fields = []string{"Id"}
	value, err := GetFieldValueFromIn(jsonData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	log.Debug("Value:", value)
}

func TestGetFieldValue(t *testing.T) {
	account := `{
    "Account": {
        "records": [
            {
                "AccountNumber": "12356",
                "AccountSource": "Test Source",
                "Active__c": "Yes",
                "AnnualRevenue": "324556",
                "BillingCity": "Palo Alto",
                "BillingCountry": "USA",
                "BillingGeocodeAccuracy": null,
                "BillingLatitude": null,
                "BillingLongitude": null,
                "BillingPostalCode": "94207",
                "BillingState": "California",
                "BillingStreet": "3330 hillview ave",
                "CleanStatus": "Pending",
                "CustomerPriority__c": "High",
                "Description": "Sample Description for the account",
                "DunsNumber": "32653",
                "Fax": "345272",
                "Industry": "Engineering",
                "Jigsaw": "Test",
                "NaicsCode": "34583",
                "NaicsDesc": "Test Description",
                "Name": "may24_a",
                "Ownership": "Private",
                "ParentId": null,
                "Phone": "1234567890",
                "Rating": "Warm",
                "SLAExpirationDate__c": "2017-08-27",
                "SLA__c": "23453",
                "ShippingCity": "San Francisco",
                "ShippingCountry": "USA",
                "ShippingGeocodeAccuracy": null,
                "ShippingLatitude": null,
                "ShippingLongitude": null,
                "ShippingPostalCode": 45692,
                "ShippingState": "California",
                "ShippingStreet": "1234 Hillview Ave",
                "Sic": "Gold",
                "SicDesc": null,
                "Site": "www.example2.com",
                "TickerSymbol": null,
                "Tradestyle": "Regular",
                "Type": "Custumer-Direct",
                "UpsellOpportunity__c": "Yes",
                "Website": "www.example.com",
                "YearStarted": "2015"
            }
        ]
    }
}
`
	mappingField := &field.MappingField{HasArray: true, HasSpecialField: false}
	mappingField.Fields = []string{"Account", "records[0]", "Name"}
	value, err := GetFieldValueFromIn(account, mappingField)
	log.Infof("Value:%s", value)

	assert.Nil(t, err)
	assert.NotNil(t, value)
}

func TestConcurrentGetk(t *testing.T) {
	w := sync.WaitGroup{}
	var recovered interface{}
	//Create factory

	for r := 0; r < 100000; r++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			defer func() {
				if r := recover(); r != nil {
					recovered = r
				}
			}()
			mappingField := &field.MappingField{HasArray: true, HasSpecialField: false}
			mappingField.Fields = []string{"City[0]", "Park", "Maps", "bb"}
			value, err := GetFieldValueFromIn(jsonData, mappingField)
			assert.Nil(t, err)
			assert.NotNil(t, value)
		}(r)

	}
	w.Wait()
	assert.Nil(t, recovered)
}

func TestRootArray(t *testing.T) {
	jsonArray := `[
    {
        "Body": "test from WI",
        "MD5OfBody": "ec7d5c27e25bcd3d6a65362b71bd0525",
        "MD5OfMessageAttributes": "50df80e5fea57210bb8167abfd053899",
        "MessageAttributes": {
            "MA1": "test"
        },
        "MessageId": "1c0483d9-8166-4df0-be9f-cd03177a38c6",
        "ReceiptHandle": "AQEBE6elNqdJKrTz5A2X/gQJETxPdtJgAktTAuT4pvBTjQgnJpSEPhfMI08fHCMrEX6ILD0fTY2FVPCCJ8LfMvAxp+LO2/Bsi1uZhUyesFoj11Y/4jvdYSCQhqWEuAI1q1pxpSj2d2QbL5SUlX979ZG+Abv/IYeDvPO8nyuZ0IWgVhZWaGcoOwADvj3mNJZ9XJh8mS3vL8EQlUO6dhIRn9PxVet2fGRmm3iY1YI4N7bZsw9nxIqIYgl5kfuBNegSRcrrTOb6u9vTnHK2uiiCwJi+Io6WNGuJGF4fyFi3skk/AvCS7fjl+4MFqoHKsm1nR06Rel7017m0/Dg5KaOJCRAJ92gV4iuUMynG1WfmELMMg/sS19hrNvcgdKW5Vd3Snn/oNcoP2Ebb7CQA08XzVcoO0maVt2KqUWgvqf0DDxVArEE="
    }
]`

	mappingField := &field.MappingField{HasArray: true, HasSpecialField: false}
	mappingField.Fields = []string{"[0]", "MessageId"}
	value, err := GetFieldValueFromIn(jsonArray, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
}

func TestRootArrayInvalid(t *testing.T) {
	jsonArray := `[
    {
        "Body": "test from WI",
        "MD5OfBody": "ec7d5c27e25bcd3d6a65362b71bd0525",
        "MD5OfMessageAttributes": "50df80e5fea57210bb8167abfd053899",
        "MessageAttributes": {
            "MA1": "test"
        },
        "MessageId": "1c0483d9-8166-4df0-be9f-cd03177a38c6",
        "ReceiptHandle": "AQEBE6elNqdJKrTz5A2X/gQJETxPdtJgAktTAuT4pvBTjQgnJpSEPhfMI08fHCMrEX6ILD0fTY2FVPCCJ8LfMvAxp+LO2/Bsi1uZhUyesFoj11Y/4jvdYSCQhqWEuAI1q1pxpSj2d2QbL5SUlX979ZG+Abv/IYeDvPO8nyuZ0IWgVhZWaGcoOwADvj3mNJZ9XJh8mS3vL8EQlUO6dhIRn9PxVet2fGRmm3iY1YI4N7bZsw9nxIqIYgl5kfuBNegSRcrrTOb6u9vTnHK2uiiCwJi+Io6WNGuJGF4fyFi3skk/AvCS7fjl+4MFqoHKsm1nR06Rel7017m0/Dg5KaOJCRAJ92gV4iuUMynG1WfmELMMg/sS19hrNvcgdKW5Vd3Snn/oNcoP2Ebb7CQA08XzVcoO0maVt2KqUWgvqf0DDxVArEE="
    },
	    {
        "Body": "test from WI2",
        "MD5OfBody": "ec7d5c27e25bcd33d6a65362b71bd0525",
        "MD5OfMessageAttributes": "50df80e5fea57210bb8167abfd053899",
        "MessageAttributes": {
            "MA1": "test"
        },
        "MessageId": "==1c04833d9-8166-4df0-be9f-cd03177a38c6",
        "ReceiptHandle": "AQE3BE6elNqdJKrTz5A2X/gQJETxPdtJgAktTAuT4pvBTjQgnJpSEPhfMI08fHCMrEX6ILD0fTY2FVPCCJ8LfMvAxp+LO2/Bsi1uZhUyesFoj11Y/4jvdYSCQhqWEuAI1q1pxpSj2d2QbL5SUlX979ZG+Abv/IYeDvPO8nyuZ0IWgVhZWaGcoOwADvj3mNJZ9XJh8mS3vL8EQlUO6dhIRn9PxVet2fGRmm3iY1YI4N7bZsw9nxIqIYgl5kfuBNegSRcrrTOb6u9vTnHK2uiiCwJi+Io6WNGuJGF4fyFi3skk/AvCS7fjl+4MFqoHKsm1nR06Rel7017m0/Dg5KaOJCRAJ92gV4iuUMynG1WfmELMMg/sS19hrNvcgdKW5Vd3Snn/oNcoP2Ebb7CQA08XzVcoO0maVt2KqUWgvqf0DDxVArEE="
    }
]`

	mappingField := &field.MappingField{HasArray: true, HasSpecialField: false}
	mappingField.Fields = []string{"[0]", "MessageId[0]"}
	value, err := GetFieldValueFromIn(jsonArray, mappingField)
	assert.NotNil(t, err)
	assert.Nil(t, value)
}

func TestGetValue(t *testing.T) {
	value := struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		IntV   int    `json:"int_v"`
		Int64V int64  `json:"int_64"`
	}{
		ID:     "12222",
		Name:   "name",
		Int64V: int64(123),
		IntV:   int(12),
	}

	mappingField := &field.MappingField{HasArray: false, HasSpecialField: false}
	mappingField.Fields = []string{"id"}

	v, err := GetFieldValueFromIn(value, mappingField)
	assert.Nil(t, err)
	assert.Equal(t, "12222", v)

	testMap := make(map[string]string)
	testMap["id"] = "id"
	testMap["id2"] = "id2"

	testMap2 := make(map[string]interface{})
	testMap2["id"] = value
	testMap2["id2"] = int(2)

	mappingField2 := &field.MappingField{HasArray: false, HasSpecialField: false}
	mappingField2.Fields = []string{"id"}
	v, err = GetFieldValueFromIn(testMap, mappingField2)
	assert.Nil(t, err)
	assert.Equal(t, "id", v)

	mappingField3 := &field.MappingField{HasArray: false, HasSpecialField: false}
	mappingField3.Fields = []string{"id2"}
	v, err = GetFieldValueFromIn(testMap2, mappingField3)
	assert.Nil(t, err)
	assert.Equal(t, float64(2), v)

	mappingField4 := &field.MappingField{HasArray: false, HasSpecialField: false}
	mappingField4.Fields = []string{"id", "id"}
	v, err = GetFieldValueFromIn(testMap2, mappingField4)
	assert.Nil(t, err)
	assert.Equal(t, "12222", v)

	////Int64
	mappingFieldInt64 := &field.MappingField{HasArray: false, HasSpecialField: false}
	mappingFieldInt64.Fields = []string{"id", "int_64"}
	v, err = GetFieldValueFromIn(testMap2, mappingFieldInt64)
	assert.Nil(t, err)
	assert.Equal(t, float64(123), v)
	//Int
	mappingFieldint := &field.MappingField{HasArray: false, HasSpecialField: false}
	mappingFieldint.Fields = []string{"id", "int_v"}
	v, err = GetFieldValueFromIn(testMap2, mappingFieldint)
	assert.Nil(t, err)
	assert.Equal(t, float64(12), v)
}

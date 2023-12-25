package entities

import (
	"encoding/json"
	"testing"
)

func TestShouldUnmarshalValue(t *testing.T) {
	var res []User
	err := json.Unmarshal([]byte(fetchUsersJson), &res)

	if err != nil {
		t.Errorf("error should be nil. found %v", err)
	}

	if res == nil {
		t.Errorf("should have unmarshalled values")
	}
}

const fetchUsersJson = `[
    {
        "updatedAt": 1700849299336,
        "createdAt": 1700849299336,
        "createdBy": {
            "userID": "",
            "username": "",
            "email": ""
        },
        "updatedBy": {
            "userID": "",
            "username": "",
            "email": ""
        },
        "isRemoved": false,
        "userID": "dbf961dc-fe4c-473f-8254-8793ee1e2352",
        "username": "fakeuser",
        "role": "admin"
    },
    {
        "updatedAt": 1701175229469,
        "createdAt": 1701175229469,
        "createdBy": {
            "userID": "",
            "username": "",
            "email": ""
        },
        "updatedBy": {
            "userID": "",
            "username": "",
            "email": ""
        },
        "isRemoved": false,
        "userID": "81f9729a-55eb-4667-b1ea-42f7b0de0606",
        "username": "fake.name@fakecompany.net",
        "email": "fake.name@fakecompany.net",
        "name": "Fake Name",
        "role": "user"
    }
]`

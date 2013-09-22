package eventhub

import (
    "time"
    "testing"
)

func RunDataBackendTest(t *testing.T, d DataBackend) {

	data := struct {
		Foo string
	}{
		"bar",
	}

	e := Event{
		Key:         "foo.bar",
		Created:     time.Now(),
		Payload:     data,
		Description: "My event",
		Importance:  3,
		Origin:      "mysystem",
		Entities:    []string{"ns/foo", "ns/moo"},
		Actors:      []string{"someone"},
	}

	d.Save(&e)

	if e.ID != 1 {
		t.Errorf("Expected '%d', got %v", 1, e.ID)
	}

	newE, err := d.GetById(1)

	if err != nil {
		t.Error("PostgresDataSource has error:", err)
		return
	}

	t.Logf("%v", newE)

}
package eventhub

import (
	"github.com/google/go-querystring/query"
	"github.com/gorilla/schema"
	"testing"
)

func TestEncodeDecodeQuery(t *testing.T) {

	expected := "entities=c%2F1&entities=c%2F2&origin=mysystem"
	q := Query{}

	q.Origin = "mysystem"
	q.Entities = []string{"c/1", "c/2"}

	v, err := query.Values(q)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(v.Encode())

	if v.Encode() != expected {
		t.Fatalf("Expected %s, got %s", expected, v.Encode())
	}

	q2 := new(Query)
	decoder := schema.NewDecoder()
	err = decoder.Decode(q2, v)

	if err != nil {
		t.Fatal(err)
	}

	if q.Origin != q2.Origin {
		t.Fatalf("Expected %s, got %s", q.Origin, q2.Origin)
	}
}

func TestMatchFilter(t *testing.T) {

	q := Query{}
	q.Entities = []string{"ns/moo"}
	t.Logf("Query filter: %+v", q)

	filtered := NewEvent(
		"Should filter",
		nil,
		nil,
		"This event should be filtered",
		3,
		"myapp",
		[]string{"ns/foo", "ns/boo"},
		nil,
		nil,
		nil)

	if q.Match(*filtered) == true {
		t.Errorf("Query %+v should not pass Event: %+v", q, filtered)
	}

}
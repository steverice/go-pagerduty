package pagerduty

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestRulesetList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/global-event-rules/rulesets", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"rulesets":[{"id": "1", "name": "Ruleset", "teamId": "POOPBUG"}]}`))
	})

	resp, _, err := client.Rulesets.List()
	if err != nil {
		t.Fatal(err)
	}

	want := &ListRulesetsResponse{
		Rulesets: []*ListedRuleset{
			{
				ID:     "1",
				Name:   "Ruleset",
				TeamID: "POOPBUG",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestRulesetCreate(t *testing.T) {
	setup()
	defer teardown()
	input := &Ruleset{
		Name:   "Ruleset",
		TeamID: "POOPBUG",
	}

	mux.HandleFunc("/global-event-rules/rulesets", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(Ruleset)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"id":"RULESET_GUID","name":"Ruleset","team_id":"POOPBUG","type":"team","routing_keys":["ROUTING_KEY"],"rules":[{"actions":[["suppress","true"]],"catch_all":true,"condition": null,"advanced_condition":null,"disabled":false,"id":"RULE_GUID"}]}`))
	})

	resp, _, err := client.Rulesets.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &Ruleset{
		Name:           "Ruleset",
		ID:             "RULESET_GUID",
		DefaultRuleset: false,
		TeamID:         "POOPBUG",
		Type:           "team",
		RoutingKeys:    []interface{}{"ROUTING_KEY"},
		Rules: []*EventRule{
			{
				Actions:           []interface{}{[]interface{}{"suppress", "true"}},
				Condition:         nil,
				CatchAll:          true,
				AdvancedCondition: nil,
				ID:                "RULE_GUID",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestRulesetDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/global-event-rules/rulesets/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.Rulesets.Delete("1"); err != nil {
		t.Fatal(err)
	}
}

func TestRulesetGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/global-event-rules/rulesets/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"id": "1"}`))
	})

	resp, _, err := client.Rulesets.Get("1")
	if err != nil {
		t.Fatal(err)
	}

	want := &Ruleset{
		ID: "1",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestRulesetUpdate(t *testing.T) {
	setup()
	defer teardown()

	input := &Ruleset{
		Name: "foo",
	}

	mux.HandleFunc("/global-event-rules/rulesets/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		v := new(Ruleset)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write([]byte(`{"id":"RULESET_GUID","name":"Ruleset","team_id":"POOPBUG","type":"team","routing_keys":["ROUTING_KEY"],"rules":[{"actions":[["suppress","true"]],"catch_all":false,"condition": null,"advanced_condition":null,"disabled":false,"id":"ROUTE"},{"actions":[["suppress","true"]],"catch_all":true,"condition": null,"advanced_condition":null,"disabled":false,"id":"FALLBACK"}]}`))
	})

	resp, _, err := client.Rulesets.Update("1", input)
	if err != nil {
		t.Fatal(err)
	}

	want := &Ruleset{
		Name:           "Ruleset",
		ID:             "RULESET_GUID",
		DefaultRuleset: false,
		TeamID:         "POOPBUG",
		Type:           "team",
		RoutingKeys:    []interface{}{"ROUTING_KEY"},
		Rules: []*EventRule{
			{
				Actions:           []interface{}{[]interface{}{"suppress", "true"}},
				Condition:         nil,
				CatchAll:          false,
				AdvancedCondition: nil,
				ID:                "ROUTE",
			},
			{
				Actions:           []interface{}{[]interface{}{"suppress", "true"}},
				Condition:         nil,
				CatchAll:          true,
				AdvancedCondition: nil,
				ID:                "FALLBACK",
			},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

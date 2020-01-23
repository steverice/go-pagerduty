package pagerduty

import "fmt"

// RulesetService handles the communication with rulesets
// related methods of the PagerDuty API.
type RulesetService service

// Ruleset represents a ruleset.
type Ruleset struct {
	Name           string        `json:"name,omitempty"`
	DefaultRuleset bool          `json:"default_ruleset,omitempty"`
	RoutingKeys    []interface{} `json:"routing_keys,omitempty"`
	Rules          []*EventRule  `json:"rules,omitempty"`
	TeamID         string        `json:"team_id,omitempty"`
	Type           string        `json:"type,omitempty"`
	ObjectVersion  string        `json:"object_version,omitempty"`
	FormatVersion  int           `json:"format_version,string,omitempty"`
	ID             string        `json:"id,omitempty"`
}

// ListedRuleset handles the inconsistency in rulesets fetched in a list
// from those fetched by ID
type ListedRuleset struct {
	Name   string `json:"name,omitempty"`
	TeamID string `json:"teamId,omitempty"`
	ID     string `json:"id,omitempty"`
}

// ListRulesetsResponse represents a list response of rulesets.
type ListRulesetsResponse struct {
	Rulesets []*ListedRuleset `json:"rulesets,omitempty"`
}

// List lists existing rulesets.
func (s *RulesetService) List() (*ListRulesetsResponse, *Response, error) {
	u := "/global-event-rules/rulesets"
	v := new(ListRulesetsResponse)

	resp, err := s.client.newRequestDo("GET", u, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// Create creates a new ruleset.
func (s *RulesetService) Create(ruleset *Ruleset) (*Ruleset, *Response, error) {
	u := "/global-event-rules/rulesets"
	v := new(Ruleset)

	resp, err := s.client.newRequestDo("POST", u, nil, ruleset, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// Delete deletes an existing ruleset.
func (s *RulesetService) Delete(id string) (*Response, error) {
	u := fmt.Sprintf("/global-event-rules/rulesets/%s", id)
	return s.client.newRequestDo("DELETE", u, nil, nil, nil)
}

// Update updates an existing ruleset.
func (s *RulesetService) Update(id string, ruleset *Ruleset) (*Ruleset, *Response, error) {
	u := fmt.Sprintf("/global-event-rules/rulesets/%s", id)
	v := new(Ruleset)

	resp, err := s.client.newRequestDo("PUT", u, nil, ruleset, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

// Get retrieves information about a ruleset.
func (s *RulesetService) Get(id string) (*Ruleset, *Response, error) {
	u := fmt.Sprintf("/global-event-rules/rulesets/%s", id)
	v := new(Ruleset)

	resp, err := s.client.newRequestDo("GET", u, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

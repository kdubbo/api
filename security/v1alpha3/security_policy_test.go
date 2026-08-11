// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0.

package v1alpha3

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestAuthorizationPolicyExtendedFieldsJSONRoundTrip(t *testing.T) {
	in := &AuthorizationPolicy{
		Action: AuthorizationPolicy_CUSTOM,
		Rules: []*Rule{{
			From: []*From{{Source: &Source{
				Principals:        []string{"cluster.local/ns/default/sa/client"},
				RemoteIpBlocks:    []string{"203.0.113.0/24"},
				ServiceAccounts:   []string{"default/client"},
				NotIpBlocks:       []string{"10.0.0.0/8"},
				RequestPrincipals: []string{"issuer/subject"},
			}}},
			To: []*To{{Operation: &Operation{
				Ports:   []string{"8080"},
				Methods: []string{"GET"},
				Paths:   []string{"/admin/*"},
			}}},
		}},
		Provider: &ExtensionProvider{Name: "opa"},
		DryRun:   true,
	}

	data, err := json.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}
	for _, field := range []string{"\"remoteIpBlocks\"", "\"serviceAccounts\"", "\"provider\"", "\"dryRun\""} {
		if !strings.Contains(string(data), field) {
			t.Fatalf("JSON %s does not contain %s", data, field)
		}
	}

	var out AuthorizationPolicy
	if err := json.Unmarshal(data, &out); err != nil {
		t.Fatal(err)
	}
	if out.GetAction() != AuthorizationPolicy_CUSTOM || out.GetProvider().GetName() != "opa" || !out.GetDryRun() {
		t.Fatalf("round trip = %#v", &out)
	}
}

func TestRequestAuthenticationClaimHeadersJSONRoundTrip(t *testing.T) {
	in := &RequestAuthentication{JwtRules: []*JWTRule{{
		Issuer:                "https://issuer.example",
		FromCookies:           []string{"access_token"},
		ForwardOriginalToken:  true,
		OutputPayloadToHeader: "x-jwt-payload",
		OutputClaimToHeaders: []*ClaimToHeader{{
			Claim:  "nested.group",
			Header: "x-jwt-group",
		}},
	}}}

	data, err := json.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "\"outputClaimToHeaders\"") {
		t.Fatalf("JSON %s does not contain outputClaimToHeaders", data)
	}

	var out RequestAuthentication
	if err := json.Unmarshal(data, &out); err != nil {
		t.Fatal(err)
	}
	rule := out.GetJwtRules()[0]
	if !rule.GetForwardOriginalToken() || rule.GetOutputClaimToHeaders()[0].GetHeader() != "x-jwt-group" {
		t.Fatalf("round trip = %#v", &out)
	}
}

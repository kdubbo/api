// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0.

package v1alpha3

import (
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
)

func TestAuthorizationPolicyPreservesWorkloadPrincipal(t *testing.T) {
	policy := &AuthorizationPolicy{Rules: []*Rule{{
		From: []*From{{Source: &Source{
			Principals: []string{"cluster.local/ns/payments/sa/checkout"},
		}}},
	}}}
	data, err := protojson.Marshal(policy)
	if err != nil {
		t.Fatal(err)
	}
	var decoded AuthorizationPolicy
	if err := protojson.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	got := decoded.GetRules()[0].GetFrom()[0].GetSource().GetPrincipals()
	if len(got) != 1 || got[0] != "cluster.local/ns/payments/sa/checkout" {
		t.Fatalf("principals = %v", got)
	}
}

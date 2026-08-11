// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0.

package kubernetes

import (
	"os"
	"strings"
	"testing"
)

func TestAuthorizationPolicyUsesIrregularPlural(t *testing.T) {
	content, err := os.ReadFile("customresourcedefinitions.gen.yaml")
	if err != nil {
		t.Fatal(err)
	}
	crds := string(content)
	if strings.Contains(crds, "authorizationpolicys.security.dubbo.apache.org") {
		t.Fatal("AuthorizationPolicy CRD uses invalid plural authorizationpolicys")
	}
	for _, want := range []string{
		"name: authorizationpolicies.security.dubbo.apache.org",
		"plural: authorizationpolicies",
		"singular: authorizationpolicy",
	} {
		if !strings.Contains(crds, want) {
			t.Fatalf("generated CRD missing %q", want)
		}
	}
}

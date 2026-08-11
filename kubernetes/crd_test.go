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

func TestAuthorizationPolicyUsesKubernetesPlural(t *testing.T) {
	data, err := os.ReadFile("customresourcedefinitions.gen.yaml")
	if err != nil {
		t.Fatal(err)
	}
	text := string(data)
	if !strings.Contains(text, "name: authorizationpolicies.security.dubbo.apache.org") ||
		!strings.Contains(text, "plural: authorizationpolicies") {
		t.Fatal("AuthorizationPolicy CRD does not use authorizationpolicies")
	}
	if strings.Contains(text, "authorizationpolicys") {
		t.Fatal("AuthorizationPolicy CRD contains invalid plural authorizationpolicys")
	}
}

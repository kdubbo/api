// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0.

package v1alpha1

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestMeshMinimumTLSVersionJSONRoundTrip(t *testing.T) {
	in := &MeshConfig{
		MeshMtls: &MeshMTLS{MinProtocolVersion: MeshMTLS_TLSV1_3},
		ExtensionProviders: []*MeshExtensionProvider{{
			Name: "opa",
			Provider: &MeshExtensionProvider_EnvoyExtAuthzHttp{
				EnvoyExtAuthzHttp: &ExternalAuthorizationProvider{
					Service:  "opa.default.svc.cluster.local",
					Port:     9191,
					FailOpen: false,
				},
			},
		}},
	}
	data, err := json.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), `"minProtocolVersion":"TLSV1_3"`) {
		t.Fatalf("JSON = %s", data)
	}
	if !strings.Contains(string(data), `"envoyExtAuthzHttp"`) {
		t.Fatalf("JSON = %s", data)
	}

	var out MeshConfig
	if err := json.Unmarshal(data, &out); err != nil {
		t.Fatal(err)
	}
	if out.GetMeshMtls().GetMinProtocolVersion() != MeshMTLS_TLSV1_3 {
		t.Fatalf("round trip = %#v", &out)
	}
	if out.GetExtensionProviders()[0].GetEnvoyExtAuthzHttp().GetPort() != 9191 {
		t.Fatalf("round trip = %#v", &out)
	}
}

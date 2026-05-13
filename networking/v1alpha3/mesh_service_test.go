//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha3

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestMeshServiceJSONRoundTrip(t *testing.T) {
	in := &MeshService{
		Hosts: []string{"nginx.app.svc.cluster.local"},
		TrafficPolicy: &TrafficPolicy{
			Tls: &ClientTLSSettings{
				Mode: ClientTLSSettings_DUBBO_MUTUAL,
			},
		},
		Routes: []*MeshServiceRoute{
			{
				Service: []*ServiceDestination{
					{
						Name:   "v1",
						Host:   "nginx.app.svc.cluster.local",
						Labels: map[string]string{"version": "v1"},
						Port:   &ServicePort{Number: 80},
						Weight: 20,
					},
					{
						Name:   "v2",
						Host:   "nginx.app.svc.cluster.local",
						Labels: map[string]string{"version": "v2"},
						Port:   &ServicePort{Number: 80},
						Weight: 80,
					},
				},
			},
		},
	}

	raw, err := json.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}
	oldCamel := "visible" + "To"
	oldSnake := "visible" + "_to"
	if strings.Contains(string(raw), oldCamel) || strings.Contains(string(raw), oldSnake) {
		t.Fatalf("MeshService JSON = %s, must not contain removed visibility fields", raw)
	}
	if !strings.Contains(string(raw), `"trafficPolicy"`) {
		t.Fatalf("MeshService JSON = %s, want trafficPolicy field", raw)
	}

	var out MeshService
	if err := json.Unmarshal(raw, &out); err != nil {
		t.Fatal(err)
	}
	if got := out.GetTrafficPolicy().GetTls().GetMode(); got != ClientTLSSettings_DUBBO_MUTUAL {
		t.Fatalf("tls mode = %v, want DUBBO_MUTUAL", got)
	}
	if got := out.GetRoutes()[0].GetService()[1].GetLabels()["version"]; got != "v2" {
		t.Fatalf("route label = %q, want v2", got)
	}
	if got := out.GetRoutes()[0].GetService()[1].GetWeight(); got != 80 {
		t.Fatalf("route weight = %d, want 80", got)
	}
}

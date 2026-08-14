// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0.

package v1alpha3

import (
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
)

func TestMetricRuleJSON(t *testing.T) {
	var telemetry Telemetry
	err := protojson.Unmarshal([]byte(`{
		"metrics": [{
			"providers": [{"name": "prometheus"}],
			"rules": [{
				"metric": "REQUEST_COUNT",
				"scope": "CLIENT_AND_SERVER",
				"tags": {
					"grpc_response_status": {"action": "REMOVE"}
				}
			}]
		}]
	}`), &telemetry)
	if err != nil {
		t.Fatalf("unmarshal metric rule: %v", err)
	}

	rule := telemetry.GetMetrics()[0].GetRules()[0]
	if rule.GetMetric() != StandardMetric_REQUEST_COUNT {
		t.Fatalf("metric = %s, want REQUEST_COUNT", rule.GetMetric())
	}
	if rule.GetScope() != MetricScope_CLIENT_AND_SERVER {
		t.Fatalf("scope = %s, want CLIENT_AND_SERVER", rule.GetScope())
	}
	if got := rule.GetTags()["grpc_response_status"].GetAction(); got != TagOverride_REMOVE {
		t.Fatalf("tag action = %s, want REMOVE", got)
	}
}

// +build e2e

/*
 * Copyright 2020 The Knative Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package e2e_new

import (
	"testing"

	"knative.dev/eventing-kafka-broker/control-plane/pkg/reconciler/kafka"

	"knative.dev/pkg/system"
	_ "knative.dev/pkg/system/testing"
	"knative.dev/reconciler-test/pkg/k8s"
	"knative.dev/reconciler-test/pkg/knative"

	brokerfeatures "knative.dev/eventing/test/rekt/features/broker"
	"knative.dev/eventing/test/rekt/resources/broker"
)

func TestBrokerSourceToSink(t *testing.T) {
	t.Parallel()

	ctx, env := global.Environment(
		knative.WithKnativeNamespace(system.Namespace()),
		knative.WithLoggingConfig,
		knative.WithTracingConfig,
		k8s.WithEventListener,
	)

	// Install and wait for a Ready Broker.
	env.Prerequisite(ctx, t, brokerfeatures.GoesReady("default", broker.WithBrokerClass(kafka.BrokerClass)))

	// Test that a Broker can act as middleware.
	env.Test(ctx, t, brokerfeatures.SourceToSink("default"))

	env.Finish()
}

func TestBrokerIngressConformance(t *testing.T) {
	t.Parallel()

	ctx, env := global.Environment(
		knative.WithKnativeNamespace(system.Namespace()),
		knative.WithLoggingConfig,
		knative.WithTracingConfig,
		k8s.WithEventListener,
	)

	for _, f := range brokerfeatures.BrokerIngressConformanceFeatures(kafka.BrokerClass) {
		env.Test(ctx, t, f)
	}

	env.Finish()
}
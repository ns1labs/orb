// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

// Adapted for Orb project, modifications licensed under MPL v. 2.0:
/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package producer

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/ns1labs/orb/sinks"
)

const (
	streamID  = "orb.sinks"
	streamLen = 1000
)

var _ sinks.Service = (*eventStore)(nil)

type eventStore struct {
	svc    sinks.Service
	client *redis.Client
}

func (es eventStore) ListSinks(ctx context.Context, token string, pm sinks.PageMetadata) (sinks.Page, error) {
	return es.svc.ListSinks(ctx, token, pm)
}

func (es eventStore) CreateSink(ctx context.Context, token string, s sinks.Sink) (sinks.Sink, error) {
	return es.svc.CreateSink(ctx, token, s)
}

func (es eventStore) ListBackends(ctx context.Context, token string) (_ []string, err error) {
	return es.svc.ListBackends(ctx, token)
}

func (es eventStore) DeleteSink(ctx context.Context, token, id string) error {
	if err := es.svc.DeleteSink(ctx, token, id); err != nil {
		return err
	}

	event := deleteSinkEvent {
		id: id,
	}

	record := &redis.XAddArgs{
		Stream: streamID,
		MaxLenApprox: streamLen,
		Values: event.Encode(),
	}

	es.client.XAdd(ctx, record).Err()

	return nil
}

// NewEventStoreMiddleware returns wrapper around sinks service that sends
// events to event store.
func NewEventStoreMiddleware(svc sinks.Service, client *redis.Client) sinks.Service {
	return eventStore{
		svc:    svc,
		client: client,
	}
}

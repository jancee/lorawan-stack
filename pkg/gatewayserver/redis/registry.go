// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package redis

import (
	"context"
	"runtime/trace"

	"github.com/go-redis/redis"
	"go.thethings.network/lorawan-stack/pkg/errors"
	io "go.thethings.network/lorawan-stack/pkg/gatewayserver/io"
	ttnredis "go.thethings.network/lorawan-stack/pkg/redis"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/pkg/unique"
)

// GatewayConnectionStatsRegistry implements the GatewayConnectionStatsRegistry interface.
type GatewayConnectionStatsRegistry struct {
	Redis *ttnredis.Client
}

func (r *GatewayConnectionStatsRegistry) key(uid string) string {
	return r.Redis.Key("uid", uid)
}

// Set sets or clears the connection stats for a gateway.
func (r *GatewayConnectionStatsRegistry) Set(ctx context.Context, ids ttnpb.GatewayIdentifiers, stats *ttnpb.GatewayConnectionStats, newTraffic io.Traffic) error {
	uid := unique.ID(ctx, ids)
	key := r.key(uid)

	defer trace.StartRegion(ctx, "set gateway connection stats").End()

	err := r.Redis.Watch(func(tx *redis.Tx) error {
		stored := &ttnpb.GatewayConnectionStats{}
		var pipelined func(redis.Pipeliner) error

		// Delete if nil
		if stats == nil {
			pipelined = func(p redis.Pipeliner) error {
				return p.Del(key).Err()
			}
		} else {
			// Get current stats
			err := ttnredis.GetProto(tx, uid).ScanProto(stored)
			if err != nil {
				// Redis error, cannot proceed
				if !errors.IsNotFound(err) {
					return err
				}

				// No stats in store yet. Simply set.
				pipelined = func(p redis.Pipeliner) error {
					_, err := ttnredis.SetProto(p, key, stats, 0)
					return err
				}
			} else {
				// Existing stats are in stored. Update based on new traffic
				if newTraffic.Up {
					stored.LastUplinkReceivedAt = stats.LastUplinkReceivedAt
					stored.UplinkCount = stats.UplinkCount
				}
				if newTraffic.Down {
					stored.LastDownlinkReceivedAt = stats.LastDownlinkReceivedAt
					stored.DownlinkCount = stats.DownlinkCount
					stored.RoundTripTimes = stats.RoundTripTimes
				}
				if newTraffic.Status {
					stored.LastStatus = stats.LastStatus
					stored.LastStatusReceivedAt = stats.LastStatusReceivedAt
				}

				pipelined = func(p redis.Pipeliner) error {
					_, err := ttnredis.SetProto(p, key, stored, 0)
					return err
				}
			}
		}

		_, err := tx.Pipelined(pipelined)
		return err
	}, key)

	if err != nil {
		return ttnredis.ConvertError(err)
	}
	return nil
}

// Get returns the connection stats for a gateway.
func (r *GatewayConnectionStatsRegistry) Get(ctx context.Context, ids ttnpb.GatewayIdentifiers) (*ttnpb.GatewayConnectionStats, error) {
	uid := unique.ID(ctx, ids)
	stats := &ttnpb.GatewayConnectionStats{}
	err := ttnredis.GetProto(r.Redis, r.key(uid)).ScanProto(stats)
	if err != nil {
		stats = nil
	}
	return stats, err
}

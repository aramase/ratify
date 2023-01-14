/*
Copyright The Ratify Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package memory

import (
	"context"
	"time"

	et "github.com/deislabs/ratify/pkg/executor/types"
)

// Cache describes an in-memory cache with automatic expiration
type Cache struct {
	syncMap *SyncMapWithExpiration
}

func (c Cache) GetVerifyResult(ctx context.Context, subjectRefString string) (et.VerifyResult, bool) {
	item, ok := c.syncMap.GetEntry(subjectRefString)
	if !ok {
		return et.VerifyResult{}, false
	}
	return item.(et.VerifyResult), true
}

func (c Cache) SetVerifyResult(ctx context.Context, subjectRefString string, verifyResult et.VerifyResult, ttl time.Duration) {
	c.syncMap.SetEntry(subjectRefString, verifyResult, ttl)
}

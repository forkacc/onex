// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package authz

import "context"

type Authorizer interface {
	Authorize(ctx context.Context, token string, obj, act string) (userID string, allowed bool, err error)
}

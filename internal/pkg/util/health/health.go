// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package health

import (
	"net/http"

	"k8s.io/klog/v2"
)

// ServeHealthCheck runs a http server used to provide a api to check pump health status.
func ServeHealthCheck(healthPath string, healthAddress string) {
	http.HandleFunc("/"+healthPath, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	if err := http.ListenAndServe(healthAddress, nil); err != nil {
		klog.Fatalf("Error serving health check endpoint: %s", err.Error())
	}
}

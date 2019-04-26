// Copyright 2014 mqant Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package gate

import "time"

type Option func(*Options)

type Options struct {
	ConcurrentTasks int
	BufSize         int
	Heartbeat       time.Duration
	OverTime        time.Duration
	RouteHandler    RouteHandler
	StorageHandler  StorageHandler
	AgentLearner    AgentLearner
	SessionLearner  SessionLearner
	GateHandler     GateHandler
}

func NewOptions(opts ...Option) Options {
	opt := Options{
		ConcurrentTasks: 20,
		BufSize:         2048,
		Heartbeat:       time.Minute,
		OverTime:        time.Second * 10,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func ConcurrentTasks(s int) Option {
	return func(o *Options) {
		o.ConcurrentTasks = s
	}
}
func BufSize(s int) Option {
	return func(o *Options) {
		o.BufSize = s
	}
}
func Heartbeat(s time.Duration) Option {
	return func(o *Options) {
		o.Heartbeat = s
	}
}

func OverTime(s time.Duration) Option {
	return func(o *Options) {
		o.OverTime = s
	}
}

func SetRouteHandler(s RouteHandler) Option {
	return func(o *Options) {
		o.RouteHandler = s
	}
}
func SetStorageHandler(s StorageHandler) Option {
	return func(o *Options) {
		o.StorageHandler = s
	}
}
func SetAgentLearner(s AgentLearner) Option {
	return func(o *Options) {
		o.AgentLearner = s
	}
}
func SetGateHandler(s GateHandler) Option {
	return func(o *Options) {
		o.GateHandler = s
	}
}

func SetSessionLearner(s SessionLearner) Option {
	return func(o *Options) {
		o.SessionLearner = s
	}
}

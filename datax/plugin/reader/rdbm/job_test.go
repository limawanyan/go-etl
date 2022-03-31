// Copyright 2020 the go-etl Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rdbm

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Breeze0806/go-etl/config"
	"github.com/Breeze0806/go-etl/datax/common/plugin"
)

func newMockDbHandler(newQuerier func(name string, conf *config.JSON) (Querier, error)) DbHandler {
	return NewBaseDbHandler(newQuerier, nil)
}

func TestJob_Init(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		j       *Job
		conf    *config.JSON
		jobConf *config.JSON
		args    args
		wantErr bool
	}{
		{
			name: "1",
			j: NewJob(newMockDbHandler(func(name string, conf *config.JSON) (Querier, error) {
				return &MockQuerier{}, nil
			})),
			args: args{
				ctx: context.TODO(),
			},
			conf: TestJSON(),
			jobConf: TestJSONFromString(`{
			}`),
		},
		{
			name: "2",
			j: NewJob(newMockDbHandler(func(name string, conf *config.JSON) (Querier, error) {
				return &MockQuerier{}, nil
			})),
			args: args{
				ctx: context.TODO(),
			},
			conf:    TestJSONFromString(`{}`),
			jobConf: TestJSONFromString(`{}`),
			wantErr: true,
		},
		{
			name: "3",
			j: NewJob(newMockDbHandler(func(name string, conf *config.JSON) (Querier, error) {
				return &MockQuerier{}, nil
			})),
			args: args{
				ctx: context.TODO(),
			},
			conf: TestJSON(),
			jobConf: TestJSONFromString(`{
				"username": 1
			}`),
			wantErr: true,
		},
		{
			name: "4",
			j: NewJob(newMockDbHandler(func(name string, conf *config.JSON) (Querier, error) {
				return nil, errors.New("mock error")
			})),
			args: args{
				ctx: context.TODO(),
			},
			conf: TestJSON(),
			jobConf: TestJSONFromString(`{
			}`),
			wantErr: true,
		},
		{
			name: "5",
			j: NewJob(newMockDbHandler(func(name string, conf *config.JSON) (Querier, error) {
				return &MockQuerier{
					PingErr: errors.New("mock error"),
				}, nil
			})),
			args: args{
				ctx: context.TODO(),
			},
			conf: TestJSON(),
			jobConf: TestJSONFromString(`{
			}`),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.j.SetPluginConf(tt.conf)
			tt.j.SetPluginJobConf(tt.jobConf)
			if err := tt.j.Init(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Job.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJob_Destroy(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		j       *Job
		args    args
		wantErr bool
	}{

		{
			name: "1",
			j: &Job{
				Querier: &MockQuerier{},
			},
			args: args{
				ctx: context.TODO(),
			},
		},
		{
			name: "2",
			j:    &Job{},
			args: args{
				ctx: context.TODO(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.j.Destroy(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Job.Destroy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJob_Split(t *testing.T) {
	type args struct {
		ctx    context.Context
		number int
	}
	tests := []struct {
		name    string
		j       *Job
		args    args
		jobConf *config.JSON
		want    []*config.JSON
		wantErr bool
	}{
		{
			name: "1",
			j: &Job{
				BaseJob: plugin.NewBaseJob(),
			},
			args: args{
				ctx:    context.TODO(),
				number: 1,
			},
			jobConf: TestJSONFromString(`{}`),
			want: []*config.JSON{
				TestJSONFromString(`{}`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.j.SetPluginJobConf(tt.jobConf)
			got, err := tt.j.Split(tt.args.ctx, tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("Job.Split() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Job.Split() = %v, want %v", got, tt.want)
			}
		})
	}
}

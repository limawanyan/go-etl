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
	"encoding/json"
	"time"

	"github.com/Breeze0806/go-etl/config"
	rdbmreader "github.com/Breeze0806/go-etl/datax/plugin/reader/rdbm"
	"github.com/Breeze0806/go-etl/storage/database"
	"github.com/Breeze0806/go/time2"
)

// 默认参数
var (
	defalutBatchSize    = 1000
	defalutBatchTimeout = 1 * time.Second
)

//Config 关系数据库写入器配置
type Config interface {
	GetUsername() string               //获取用户名
	GetPassword() string               //获取密码
	GetURL() string                    //获取连接url
	GetColumns() []rdbmreader.Column   //获取列信息
	GetBaseTable() *database.BaseTable //获取表信息
	GetWriteMode() string              //获取写入模式
	GetBatchSize() int                 //单次批量写入数
	GetBatchTimeout() time.Duration    //单次批量写入超时时间
}

//BaseConfig 用于实现基本的关系数据库配置，如无特殊情况采用该配置，帮助快速实现writer
type BaseConfig struct {
	Username     string                `json:"username"`     //用户名
	Password     string                `json:"password"`     //密码
	Column       []string              `json:"column"`       //列信息
	Connection   rdbmreader.ConnConfig `json:"connection"`   //连接信息
	WriteMode    string                `json:"writeMode"`    //写入模式,如插入insert
	BatchSize    int                   `json:"batchSize"`    //单次批量写入数
	BatchTimeout time2.Duration        `json:"batchTimeout"` //单次批量写入超时时间
}

//NewBaseConfig 从conf解析出关系数据库配置
func NewBaseConfig(conf *config.JSON) (c *BaseConfig, err error) {
	c = &BaseConfig{}
	err = json.Unmarshal([]byte(conf.String()), c)
	if err != nil {
		return nil, err
	}
	return
}

//GetUsername 获取用户名
func (b *BaseConfig) GetUsername() string {
	return b.Username
}

//GetPassword 获取密码
func (b *BaseConfig) GetPassword() string {
	return b.Password
}

//GetURL 获取连接url
func (b *BaseConfig) GetURL() string {
	return b.Connection.URL
}

//GetColumns 获取列信息
func (b *BaseConfig) GetColumns() (columns []rdbmreader.Column) {
	for _, v := range b.Column {
		columns = append(columns, &rdbmreader.BaseColumn{
			Name: v,
		})
	}
	return
}

//GetBaseTable 获取表信息
func (b *BaseConfig) GetBaseTable() *database.BaseTable {
	return database.NewBaseTable(b.Connection.Table.Db, b.Connection.Table.Schema, b.Connection.Table.Name)
}

//GetWriteMode 获取写入模式
func (b *BaseConfig) GetWriteMode() string {
	return b.WriteMode
}

//GetBatchTimeout 单次批量超时时间
func (b *BaseConfig) GetBatchTimeout() time.Duration {
	if b.BatchTimeout.Duration == 0 {
		return defalutBatchTimeout
	}
	return b.BatchTimeout.Duration
}

//GetBatchSize 单次批量写入数
func (b *BaseConfig) GetBatchSize() int {
	if b.BatchSize == 0 {
		return defalutBatchSize
	}

	return b.BatchSize
}

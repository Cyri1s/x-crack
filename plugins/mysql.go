/*

Copyright (c) 2017 xsec.io

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THEq
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

*/

package plugins

import (
	_ "github.com/netxfly/mysql"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"

	"x-crack/models"

	"fmt"
)

func ScanMysql(service models.Service) (err error, result models.ScanResult) {
	result.Service = service

	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", service.Username,
		service.Password, service.Ip, service.Port, "mysql")
	Engine, err := xorm.NewEngine("mysql", dataSourceName)

	if err == nil {
		Engine.SetLogLevel(core.LOG_OFF)
		// fix "[mysql] packets.go:33: unexpected EOF" error
		Engine.SetMaxIdleConns(0)
		// Engine.SetConnMaxLifetime(time.Second * 30)
		defer Engine.Close()
		err = Engine.Ping()
		if err == nil {
			result.Result = true
		}
	}
	return err, result
}

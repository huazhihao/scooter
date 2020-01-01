/*
Copyright © 2019 Hua Zhihao <ihuazhihao@gmail.com>

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
package main

import (
	"fmt"

  "github.com/huazhihao/scooter/cmd"
  _ "github.com/huazhihao/scooter/pkg/commons"
	_ "github.com/huazhihao/scooter/pkg/log"
	_ "github.com/huazhihao/scooter/pkg/api"
	_ "github.com/huazhihao/scooter/pkg/commons"
	_ "github.com/huazhihao/scooter/pkg/http"
	_ "github.com/huazhihao/scooter/pkg/tcp"
)

var (
	// Version is fetched during build time
	Version string
	// GitSHA is fetched during build time
	GitSHA string
)

func main() {
	cmd.Execute(fmt.Sprintf("%s-%s", Version, GitSHA))
}

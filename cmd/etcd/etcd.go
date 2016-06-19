// Copyright © 2016 The Remco Authors
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

package etcd

import (
	"os"
	"strings"

	"github.com/HeavyHorst/remco/backends/etcd"
	"github.com/HeavyHorst/remco/template"
	"github.com/cloudflare/cfssl/log"
	"github.com/spf13/cobra"
)

type etcdConfig struct {
	nodes     []string
	cert      string
	key       string
	caCert    string
	basicAuth bool
	username  string
	password  string
	//client    backends.StoreClient
	templateRes *template.TemplateResource
}

var Cmd = &cobra.Command{
	Use:   "etcd",
	Short: "A brief description of your command",

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.Info("Backend nodes set to " + strings.Join(config.nodes, ", "))
		client, err := etcd.NewEtcdClient(config.nodes, config.cert, config.key, config.caCert, config.basicAuth, config.username, config.password)
		if err != nil {
			log.Error(err)
		}

		t, err := template.NewTemplateResource(client, "/", cmd.Flags())
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}

		config.templateRes = t
	},
}

var config = etcdConfig{}

func init() {
	Cmd.PersistentFlags().StringSliceVar(&config.nodes, "nodes", []string{"http://127.0.0.1:2379"}, "The etcd nodes")
	Cmd.PersistentFlags().StringVar(&config.cert, "cert", "", "The client cert file")
	Cmd.PersistentFlags().StringVar(&config.key, "key", "", "The client key file")
	Cmd.PersistentFlags().StringVar(&config.caCert, "caCert", "", "The client CA key file")
	Cmd.PersistentFlags().BoolVar(&config.basicAuth, "basicAuth", false, "Enable etcd basic auth with username and password")
	Cmd.PersistentFlags().StringVar(&config.username, "username", "", "username")
	Cmd.PersistentFlags().StringVar(&config.password, "password", "", "password")
}
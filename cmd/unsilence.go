/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"github.com/alicekaerast/shush/lib"
	clientruntime "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/muesli/coral"
	"github.com/prometheus/alertmanager/api/v2/client"
	"github.com/prometheus/alertmanager/api/v2/client/silence"
	"log"
	"net/url"
)

// unsilenceCmd represents the silence command
var unsilenceCmd = &coral.Command{
	Use:   "unsilence",
	Short: "Delete silences",
	Long:  `Usage: shush unsilence [-l] [--url http://localhost] --id <id>`,
	Run: func(cmd *coral.Command, args []string) {

		parsedUrl, _ := url.Parse(urlString)
		cr := clientruntime.New(parsedUrl.Host, "/api/v2", []string{parsedUrl.Scheme})
		c := client.New(cr, strfmt.Default)

		id, _ := cmd.Flags().GetString("id")

		params := silence.NewDeleteSilenceParams().WithSilenceID(strfmt.UUID(id))
		ok, err := c.Silence.DeleteSilence(params)
		if err != nil {
			log.Println(err)
		}
		log.Println(ok)

		list, _ := cmd.Flags().GetBool("list")
		if list {
			lib.ListSilences(c)
		}

	},
}

func init() {
	unsilenceCmd.Flags().BoolP("list", "l", false, "Whether to list silences")
	unsilenceCmd.Flags().StringP("id", "i", "", "The silence ID to remove")
	unsilenceCmd.MarkFlagRequired("id")
}

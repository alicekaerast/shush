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
	"github.com/prometheus/alertmanager/api/v2/models"
	"github.com/prometheus/alertmanager/cli"
	"github.com/prometheus/alertmanager/pkg/labels"
	"gopkg.in/yaml.v2"
	"log"
	"net/url"
	"os"
	"time"
)

type Silences []struct {
	Comment  string   `yaml:"comment"`
	Matchers []string `yaml:"matchers"`
}

// silenceCmd represents the silence command
var silenceCmd = &coral.Command{
	Use:   "silence",
	Short: "Create or list silences",
	Long:  `Usage: shush silence [-l] [--url http://localhost]`,
	Run: func(cmd *coral.Command, args []string) {

		yamlFile, _ := cmd.Flags().GetString("yaml")
		yamlContent, _ := os.ReadFile(yamlFile)
		silences := Silences{}
		err := yaml.Unmarshal(yamlContent, &silences)
		if err != nil {
			log.Fatalln(err)
		}

		parsedUrl, _ := url.Parse(urlString)
		log.Println(parsedUrl)
		cr := clientruntime.New(parsedUrl.Host, "/api/v2", []string{parsedUrl.Scheme})
		c := client.New(cr, strfmt.Default)

		user := os.Getenv("LOGNAME")
		startsAt := time.Now().UTC()
		start := strfmt.DateTime(startsAt)

		endsAt := time.Now().UTC().Add(time.Duration(7200000000000))
		end := strfmt.DateTime(endsAt)

		for _, s := range silences {

			matchers, _ := parseMatchers(s.Matchers)

			ps := &models.PostableSilence{
				Silence: models.Silence{
					Matchers:  cli.TypeMatchers(matchers),
					StartsAt:  &start,
					EndsAt:    &end,
					CreatedBy: &user,
					Comment:   &s.Comment,
				},
			}
			silenceParams := silence.NewPostSilencesParams().WithSilence(ps)
			c.Silence.PostSilences(silenceParams)
		}

		list, _ := cmd.Flags().GetBool("list")
		if list {
			lib.ListSilences(c)
		}
	},
}

func init() {
	silenceCmd.Flags().BoolP("list", "l", false, "Whether to list silences")
	silenceCmd.Flags().StringP("yaml", "y", "", "The YAML file to use for silencing")
}

func parseMatchers(inputMatchers []string) ([]labels.Matcher, error) {
	matchers := make([]labels.Matcher, 0, len(inputMatchers))

	for _, v := range inputMatchers {
		matcher, err := labels.ParseMatcher(v)
		if err != nil {
			return []labels.Matcher{}, err
		}

		matchers = append(matchers, *matcher)
	}

	return matchers, nil
}

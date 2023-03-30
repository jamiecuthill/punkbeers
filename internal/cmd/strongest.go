/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"errors"
	"sort"

	"github.com/jamiecuthill/punkbeers/internal/punkapi"
	"github.com/spf13/cobra"
)

// strongestCmd represents the strongest command
var strongestCmd = &cobra.Command{
	Use:   "strongest",
	Short: "Finds the strongest beer available",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := punkapi.NewClient(punkapiURL)
		if err != nil {
			return err
		}

		number, err := cmd.Flags().GetUint("n")
		if err != nil {
			return err
		}

		var input *punkapi.AllBeersInput
		if food, err := cmd.Flags().GetString("food"); err == nil {
			input = &punkapi.AllBeersInput{Food: food}
		}

		beers, err := c.AllBeers(input)
		if err != nil {
			return err
		}

		if len(beers) == 0 {
			return errors.New("no beers found")
		}

		sort.Sort(sort.Reverse(byAbv(beers)))
		for _, beer := range beers[:number] {
			cmd.Printf("%s [%.2f]\n", beer.Name, beer.Abv)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(strongestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// strongestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	strongestCmd.Flags().Uint("n", 1, "Number of strong beers to return")
	strongestCmd.Flags().String("food", "", "Return only beers matching a food pairing")
}

type byAbv []punkapi.Beer

func (a byAbv) Len() int           { return len(a) }
func (a byAbv) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAbv) Less(i, j int) bool { return a[i].Abv < a[j].Abv }

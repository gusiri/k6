/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2016 Load Impact
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package cmd

import (
	"os"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var archiveOut = "archive.tar"

// archiveCmd represents the pause command
var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Create an archive",
	Long: `Create an archive.

An archive is a fully self-contained test run, and can be executed identically elsewhere.`,
	Example: `
  # Archive a test run.
  k6 archive -u 10 -d 10s -O myarchive.tar script.js

  # Run the resulting archive.
  k6 run myarchive.tar`[1:],
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Runner.
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}
		filename := args[0]
		fs := afero.NewOsFs()
		src, err := readSource(filename, pwd, fs, os.Stdin)
		if err != nil {
			return err
		}

		runtimeOptions, err := getRuntimeOptions(cmd.Flags())
		if err != nil {
			return err
		}

		r, err := newRunner(src, runType, fs, runtimeOptions)
		if err != nil {
			return err
		}

		cliOpts, err := getOptions(cmd.Flags())
		if err != nil {
			return err
		}
		conf, err := getConsolidatedConfig(fs, Config{Options: cliOpts}, r)
		if err != nil {
			return err
		}

		if _, cerr := deriveAndValidateConfig(conf); cerr != nil {
			return ExitCode{cerr, invalidConfigErrorCode}
		}

		err = r.SetOptions(conf.Options)
		if err != nil {
			return err
		}

		// Archive.
		arc := r.MakeArchive()
		f, err := os.Create(archiveOut)
		if err != nil {
			return err
		}
		return arc.Write(f)
	},
}

func archiveCmdFlagSet() *pflag.FlagSet {
	flags := pflag.NewFlagSet("", pflag.ContinueOnError)
	flags.SortFlags = false
	flags.AddFlagSet(optionFlagSet())
	flags.AddFlagSet(runtimeOptionFlagSet(false))
	//TODO: figure out a better way to handle the CLI flags - global variables are not very testable... :/
	flags.StringVarP(&archiveOut, "archive-out", "O", archiveOut, "archive output filename")
	return flags
}

func init() {
	RootCmd.AddCommand(archiveCmd)
	archiveCmd.Flags().SortFlags = false
	archiveCmd.Flags().AddFlagSet(archiveCmdFlagSet())
}

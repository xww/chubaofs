// Copyright 2018 The Chubao Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/chubaofs/chubaofs/sdk/master"
	"github.com/spf13/cobra"
)

const (
	cmdRootShort = "ChubaoFS Command Line Interface (CLI)"
)

func NewRootCmd(client *master.MasterClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   path.Base(os.Args[0]),
		Short: cmdRootShort,
		Args:  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(
		newClusterCmd(client),
		newVolCmd(client),
		newUserCmd(client),
		newMetaNodeCmd(client),
		newDataNodeCmd(client),
	)
	return cmd
}

func stdout(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, format, a...)
}

func errout(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, a...)
}

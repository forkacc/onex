// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package user

import (
	"context"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericiooptions"

	cmdutil "github.com/onexstack/onex/internal/onexctl/cmd/util"
	"github.com/onexstack/onex/internal/onexctl/util/templates"
	v1 "github.com/onexstack/onex/pkg/api/usercenter/v1"
)

const (
	defaultLimit = 1000
)

// ListOptions is an options struct to support list subcommands.
type ListOptions struct {
	Offset int64
	Limit  int64

	ListUserRequest *v1.ListUserRequest

	client v1.UserCenterHTTPClient
	genericiooptions.IOStreams
}

var listExample = templates.Examples(`
		# List all users
		onexctl user list

		# List users with limit and offset
		onexctl user list --offset=0 --limit=10`)

// NewListOptions returns an initialized ListOptions instance.
func NewListOptions(ioStreams genericiooptions.IOStreams) *ListOptions {
	return &ListOptions{
		IOStreams: ioStreams,
		Offset:    0,
		Limit:     defaultLimit,
	}
}

// NewCmdList returns new initialized instance of list sub command.
func NewCmdList(f cmdutil.Factory, ioStreams genericiooptions.IOStreams) *cobra.Command {
	o := NewListOptions(ioStreams)

	cmd := &cobra.Command{
		Use:                   "list",
		DisableFlagsInUseLine: true,
		Aliases:               []string{},
		Short:                 "Display all users in onex platform (Administrator rights required)",
		TraverseChildren:      true,
		Long:                  "Display all users in onex platform (Administrator rights required).",
		Example:               listExample,
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(o.Complete(f, cmd, args))
			cmdutil.CheckErr(o.Validate(cmd, args))
			cmdutil.CheckErr(o.Run(f, args))
		},
		SuggestFor: []string{},
	}

	cmd.Flags().Int64VarP(&o.Offset, "offset", "o", o.Offset, "Specify the offset of the first row to be returned.")
	cmd.Flags().Int64VarP(&o.Limit, "limit", "l", o.Limit, "Specify the amount records to be returned.")

	return cmd
}

// Complete completes all the required options.
func (o *ListOptions) Complete(f cmdutil.Factory, cmd *cobra.Command, args []string) error {
	o.ListUserRequest = &v1.ListUserRequest{
		Limit:  o.Limit,
		Offset: o.Offset,
	}

	o.client = f.UserCenterClient()
	return nil
}

// Validate makes sure there is no discrepency in command options.
func (o *ListOptions) Validate(cmd *cobra.Command, args []string) error {
	return o.ListUserRequest.Validate()
}

// Run executes a list subcommand using the specified options.
func (o *ListOptions) Run(f cmdutil.Factory, args []string) error {
	users, err := o.client.ListUser(context.Background(), o.ListUserRequest)
	if err != nil {
		return err
	}

	data := make([][]string, 0, 1)
	table := tablewriter.NewWriter(o.Out)

	for _, user := range users.Users {
		data = append(data, []string{
			user.Username,
			user.UserID,
			user.Nickname,
			user.Email,
			user.Phone,
			user.CreatedAt.AsTime().Format(time.DateTime),
			user.UpdatedAt.AsTime().Format(time.DateTime),
		})
	}

	table = setHeader(table)
	table = cmdutil.TableWriterDefaultConfig(table)
	table.AppendBulk(data)
	table.Render()

	return nil
}

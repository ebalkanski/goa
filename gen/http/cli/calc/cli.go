// Code generated by goa v3.3.1, DO NOT EDIT.
//
// calc HTTP client CLI support package
//
// Command:
// $ goa gen github.com/ebalkanski/goa/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	userc "github.com/ebalkanski/goa/gen/http/user/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `user (fetch|fetch-all|create|edit|delete)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` user fetch --name "Bob"` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		userFlags = flag.NewFlagSet("user", flag.ContinueOnError)

		userFetchFlags    = flag.NewFlagSet("fetch", flag.ExitOnError)
		userFetchNameFlag = userFetchFlags.String("name", "REQUIRED", "")

		userFetchAllFlags = flag.NewFlagSet("fetch-all", flag.ExitOnError)

		userCreateFlags    = flag.NewFlagSet("create", flag.ExitOnError)
		userCreateBodyFlag = userCreateFlags.String("body", "REQUIRED", "")

		userEditFlags    = flag.NewFlagSet("edit", flag.ExitOnError)
		userEditBodyFlag = userEditFlags.String("body", "REQUIRED", "")

		userDeleteFlags    = flag.NewFlagSet("delete", flag.ExitOnError)
		userDeleteNameFlag = userDeleteFlags.String("name", "REQUIRED", "")
	)
	userFlags.Usage = userUsage
	userFetchFlags.Usage = userFetchUsage
	userFetchAllFlags.Usage = userFetchAllUsage
	userCreateFlags.Usage = userCreateUsage
	userEditFlags.Usage = userEditUsage
	userDeleteFlags.Usage = userDeleteUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "user":
			svcf = userFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "user":
			switch epn {
			case "fetch":
				epf = userFetchFlags

			case "fetch-all":
				epf = userFetchAllFlags

			case "create":
				epf = userCreateFlags

			case "edit":
				epf = userEditFlags

			case "delete":
				epf = userDeleteFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "user":
			c := userc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "fetch":
				endpoint = c.Fetch()
				data, err = userc.BuildFetchPayload(*userFetchNameFlag)
			case "fetch-all":
				endpoint = c.FetchAll()
				data = nil
			case "create":
				endpoint = c.Create()
				data, err = userc.BuildCreatePayload(*userCreateBodyFlag)
			case "edit":
				endpoint = c.Edit()
				data, err = userc.BuildEditPayload(*userEditBodyFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = userc.BuildDeletePayload(*userDeleteNameFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// userUsage displays the usage of the user command and its subcommands.
func userUsage() {
	fmt.Fprintf(os.Stderr, `The user service process users
Usage:
    %s [globalflags] user COMMAND [flags]

COMMAND:
    fetch: Fetch user.
    fetch-all: Fetch all users.
    create: Create new user.
    edit: Edit user.
    delete: Delete user.

Additional help:
    %s user COMMAND --help
`, os.Args[0], os.Args[0])
}
func userFetchUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user fetch -name STRING

Fetch user.
    -name STRING: 

Example:
    `+os.Args[0]+` user fetch --name "Bob"
`, os.Args[0])
}

func userFetchAllUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user fetch-all

Fetch all users.

Example:
    `+os.Args[0]+` user fetch-all
`, os.Args[0])
}

func userCreateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user create -body JSON

Create new user.
    -body JSON: 

Example:
    `+os.Args[0]+` user create --body '{
      "age": 25,
      "name": "Bob"
   }'
`, os.Args[0])
}

func userEditUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user edit -body JSON

Edit user.
    -body JSON: 

Example:
    `+os.Args[0]+` user edit --body '{
      "age": 25,
      "name": "Bob"
   }'
`, os.Args[0])
}

func userDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user delete -name STRING

Delete user.
    -name STRING: 

Example:
    `+os.Args[0]+` user delete --name "Bob"
`, os.Args[0])
}

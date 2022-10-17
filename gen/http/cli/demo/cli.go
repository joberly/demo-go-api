// Code generated by goa v3.10.1, DO NOT EDIT.
//
// demo HTTP client CLI support package
//
// Command:
// $ goa gen github.com/joberly/demo-go-api/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	democ "github.com/joberly/demo-go-api/gen/http/demo/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `demo rand
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` demo rand --body '{
      "max": 6591451045836003222,
      "min": 8613984219942708242
   }'` + "\n" +
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
		demoFlags = flag.NewFlagSet("demo", flag.ContinueOnError)

		demoRandFlags    = flag.NewFlagSet("rand", flag.ExitOnError)
		demoRandBodyFlag = demoRandFlags.String("body", "REQUIRED", "")
	)
	demoFlags.Usage = demoUsage
	demoRandFlags.Usage = demoRandUsage

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
		case "demo":
			svcf = demoFlags
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
		case "demo":
			switch epn {
			case "rand":
				epf = demoRandFlags

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
		case "demo":
			c := democ.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "rand":
				endpoint = c.Rand()
				data, err = democ.BuildRandPayload(*demoRandBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// demoUsage displays the usage of the demo command and its subcommands.
func demoUsage() {
	fmt.Fprintf(os.Stderr, `Demonstration Go Service
Usage:
    %[1]s [globalflags] demo COMMAND [flags]

COMMAND:
    rand: Rand implements rand.

Additional help:
    %[1]s demo COMMAND --help
`, os.Args[0])
}
func demoRandUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] demo rand -body JSON

Rand implements rand.
    -body JSON: 

Example:
    %[1]s demo rand --body '{
      "max": 6591451045836003222,
      "min": 8613984219942708242
   }'
`, os.Args[0])
}
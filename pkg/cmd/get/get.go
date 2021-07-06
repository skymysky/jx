package get

import (
	"encoding/json"
	"fmt"

	"github.com/jenkins-x/jx/v2/pkg/cmd/get/vault"
	"github.com/jenkins-x/jx/v2/pkg/cmd/get/vault/config"

	"github.com/jenkins-x/jx/v2/pkg/cmd/helper"

	"github.com/spf13/cobra"

	"github.com/ghodss/yaml"
	"github.com/jenkins-x/jx/v2/pkg/cmd/opts"
	"github.com/jenkins-x/jx/v2/pkg/cmd/templates"
	"github.com/jenkins-x/jx/v2/pkg/util"
)

// Options is the start of the data required to perform the operation.  As new fields are added, add them here instead of
// referencing the cmd.Flags()
type Options struct {
	*opts.CommonOptions

	Output string
}

const (
	valid_resources = `Valid resource types include:

    * environments (aka 'env')
    * pipelines (aka 'pipe')
    * urls (aka 'url')
    `
)

var (
	get_long = templates.LongDesc(`
		Display one or more resources.

		` + valid_resources + `

`)

	get_example = templates.Examples(`
		# List all pipelines
		jx get pipeline

		# List all URLs for services in the current namespace
		jx get url
	`)
)

// NewCmdGet creates a command object for the generic "get" action, which
// retrieves one or more resources from a server.
func NewCmdGet(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &Options{
		CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "get TYPE [flags]",
		Short:   "Display one or more resources",
		Long:    get_long,
		Example: get_example,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			helper.CheckErr(err)
		},
		SuggestFor: []string{"list", "ps"},
	}

	cmd.AddCommand(NewCmdGetActivity(commonOpts))
	cmd.AddCommand(NewCmdGetAddon(commonOpts))
	cmd.AddCommand(NewCmdGetApps(commonOpts))
	cmd.AddCommand(NewCmdGetApplications(commonOpts))
	cmd.AddCommand(NewCmdGetBranchPattern(commonOpts))
	cmd.AddCommand(NewCmdGetBuild(commonOpts))
	cmd.AddCommand(NewCmdGetBuildPack(commonOpts))
	cmd.AddCommand(NewCmdGetChat(commonOpts))
	cmd.AddCommand(NewCmdGetConfig(commonOpts))
	cmd.AddCommand(NewCmdGetCRDCount(commonOpts))
	cmd.AddCommand(NewCmdGetCVE(commonOpts))
	cmd.AddCommand(NewCmdGetDevPod(commonOpts))
	cmd.AddCommand(NewCmdGetEnv(commonOpts))
	cmd.AddCommand(NewCmdGetGit(commonOpts))
	cmd.AddCommand(NewCmdGetHelmBin(commonOpts))
	cmd.AddCommand(NewCmdGetIssue(commonOpts))
	cmd.AddCommand(NewCmdGetIssues(commonOpts))
	cmd.AddCommand(NewCmdGetLimits(commonOpts))
	cmd.AddCommand(NewCmdGetLang(commonOpts))
	cmd.AddCommand(NewCmdGetPipeline(commonOpts))
	cmd.AddCommand(NewCmdGetPreview(commonOpts))
	cmd.AddCommand(NewCmdGetQuickstartLocation(commonOpts))
	cmd.AddCommand(NewCmdGetQuickstarts(commonOpts))
	cmd.AddCommand(NewCmdGetRelease(commonOpts))
	cmd.AddCommand(NewCmdGetStorage(commonOpts))
	cmd.AddCommand(NewCmdGetTeam(commonOpts))
	cmd.AddCommand(NewCmdGetTeamRole(commonOpts))
	cmd.AddCommand(NewCmdGetToken(commonOpts))
	cmd.AddCommand(NewCmdGetTracker(commonOpts))
	cmd.AddCommand(NewCmdGetURL(commonOpts))
	cmd.AddCommand(NewCmdGetUser(commonOpts))
	cmd.AddCommand(vault.NewCmdGetVault(commonOpts))
	cmd.AddCommand(config.NewCmdGetVaultConfig(commonOpts))
	cmd.AddCommand(NewCmdGetStream(commonOpts))
	cmd.AddCommand(NewCmdGetPlugins(commonOpts))
	return cmd
}

// Run implements this command
func (o *Options) Run() error {
	return o.Cmd.Help()
}

// AddGetFlags adds an output flag to change the format of the output
func (o *Options) AddGetFlags(cmd *cobra.Command) {
	o.Cmd = cmd
	cmd.Flags().StringVarP(&o.Output, "output", "o", "", "The output format such as 'yaml'")
}

// renderResult renders the result in a given output format
func (o *Options) renderResult(value interface{}, format string) error {
	switch format {
	case "json":
		data, err := json.Marshal(value)
		if err != nil {
			return err
		}
		_, e := o.Out.Write(data)
		return e
	case "yaml":
		data, err := yaml.Marshal(value)
		if err != nil {
			return err
		}
		_, e := o.Out.Write(data)
		return e
	default:
		return fmt.Errorf("Unsupported output format: %s", format)
	}
}

func formatInt32(n int32) string {
	return util.Int32ToA(n)
}

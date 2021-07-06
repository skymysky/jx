package get

import (
	"github.com/jenkins-x/jx-logging/pkg/log"
	"github.com/jenkins-x/jx/v2/pkg/cmd/helper"
	"github.com/jenkins-x/jx/v2/pkg/cmd/opts"
	"github.com/jenkins-x/jx/v2/pkg/cmd/templates"
	"github.com/jenkins-x/jx/v2/pkg/util"
	"github.com/spf13/cobra"
)

// GetHelmBinOptions containers the CLI options
type GetHelmBinOptions struct {
	Options
}

var (
	getHelmBinLong = templates.LongDesc(`
		Display the Helm binary name used in pipelines.

		This setting lets you switch from the stable release to early access releases (e.g. from Helm 2 <-> 3)
`)

	getHelmBinExample = templates.Examples(`
		# List the git branch patterns for the current team
		jx get helmbin
	`)
)

// NewCmdGetHelmBin creates the new command for: jx get env
func NewCmdGetHelmBin(commonOpts *opts.CommonOptions) *cobra.Command {
	options := &GetHelmBinOptions{
		Options: Options{
			CommonOptions: commonOpts,
		},
	}
	cmd := &cobra.Command{
		Use:     "helmbin",
		Short:   "Display the Helm binary name used in the pipelines",
		Aliases: []string{"helm"},
		Long:    getHelmBinLong,
		Example: getHelmBinExample,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			helper.CheckErr(err)
		},
	}

	options.AddGetFlags(cmd)
	return cmd
}

// Run implements this command
func (o *GetHelmBinOptions) Run() error {
	helm, _, _, err := o.TeamHelmBin()
	if err != nil {
		return err
	}
	log.Logger().Infof("Your team uses the helm binary: %s", util.ColorInfo(helm))
	log.Logger().Infof("To change this value use: %s", util.ColorInfo("jx edit helmbin helm3"))
	return nil
}

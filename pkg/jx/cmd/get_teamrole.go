package cmd

import (
	"io"

	"github.com/jenkins-x/jx/pkg/jx/cmd/templates"
	"github.com/jenkins-x/jx/pkg/kube"
	"github.com/jenkins-x/jx/pkg/log"
	"github.com/spf13/cobra"
)

// GetTeamRoleOptions containers the CLI options
type GetTeamRoleOptions struct {
	GetOptions
}

var (
	getTeamRoleLong = templates.LongDesc(`
		Display the roles for members of a Team
`)

	getTeamRoleExample = templates.Examples(`
		# List the team roles for the current team
		jx get teamrole

	`)
)

// NewCmdGetTeamRole creates the new command for: jx get env
func NewCmdGetTeamRole(f Factory, out io.Writer, errOut io.Writer) *cobra.Command {
	options := &GetTeamRoleOptions{
		GetOptions: GetOptions{
			CommonOptions: CommonOptions{
				Factory: f,
				Out:     out,
				Err:     errOut,
			},
		},
	}
	cmd := &cobra.Command{
		Use:     "teamroles",
		Short:   "Display the Team or Teams the current user is a member of",
		Aliases: []string{"teamrole"},
		Long:    getTeamRoleLong,
		Example: getTeamRoleExample,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			CheckErr(err)
		},
	}
	options.addGetFlags(cmd)
	return cmd
}

// Run implements this command
func (o *GetTeamRoleOptions) Run() error {
	kubeClient, ns, err := o.KubeClientAndDevNamespace()
	if err != nil {
		return err
	}
	teams, names, err := kube.GetTeamRoles(kubeClient, ns)
	if err != nil {
		return err
	}
	if len(teams) == 0 {
		log.Info(`
There are no Team roles defined so far!
`)
		return nil
	}

	table := o.CreateTable()
	table.AddRow("NAME")
	for _, name := range names {
		table.AddRow(name)
	}
	table.Render()
	return nil
}
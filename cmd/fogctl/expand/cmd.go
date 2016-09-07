package expand

import (
	"errors"
	"strings"

	"github.com/cheyang/fog/cluster"
	"github.com/cheyang/fog/types"
	"github.com/cheyang/fog/util"
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "expand",
		Short: "expand a cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("expand command takes no arguments")
			}
			name := args[len(args)-1]
			storage, err := util.GetStorage(name)

			//load spec
			flags := cmd.Flags()
			if !flags.Changed("config-file") {
				return errors.New("--config-file are mandatory")
			}
			configFile, err := flags.GetString("config-file")
			if err != nil {
				return err
			}
			spec, err := types.LoadSpec(configFile)

			roleString, err := flags.GetString("with-roles")
			if err != nil {
				return err
			}
			roles := strings.Split(roleString, ",")

			return cluster.ExpandCluster(storage, spec, roles)
		},
	}
)

func init() {
	flags := Cmd.Flags()
	flags.StringP("config-file", "f", "", "The config file.")
	flags.StringP("with-roles", "w", "", "If you need the inventory file also includes role")
}
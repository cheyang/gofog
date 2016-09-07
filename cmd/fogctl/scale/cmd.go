package scale

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/cheyang/fog/cluster"
	"github.com/cheyang/fog/util"
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "scale",
		Short: "scale out/in a cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("scale out/in command takes no arguments")
			}

			desireMap := make(map[string]int)
			var name string
			for i, arg := range args {
				if i == len(args)-1 {
					name = arg
					break
				}

				kv := strings.Split(arg, "=")

				if len(kv) == 2 {
					// desireMap[kv[0]]
					value, err := strconv.Atoi(kv[1])
					if err != nil {
						return err
					}
					key := kv[0]
					desireMap[key] = value
				} else {
					return fmt.Errorf("the format of %s is not correct!", arg)
				}
			}
			storage, err := util.GetStorage(name)
			if err != nil {
				return err
			}

			return cluster.Scale(storage, desireMap)
		},
	}
)

func init() {
	flags := Cmd.Flags()
	flags.BoolP("update-all", "u", false, "Only update the new node with ansible.")
	flags.StringP("with-roles", "w", "", "If you need the inventory file also includes role")
}

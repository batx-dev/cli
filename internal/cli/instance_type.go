package cli

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/batx-dev/cli/batainer"
)

type InstanceTypeCmd struct {
	Cluster string `help:"The cluster id to filter." short:"c"`
}

func (c *InstanceTypeCmd) Run(g *Globals) error {
	client := batainer.NewClient(g.BaseURL).SetToken(g.Token)
	res, err := client.ListInstanceTypes(g.Context, &batainer.ListInstanceTypesRequest{
		ClusterID: c.Cluster,
	})
	if err != nil {
		return fmt.Errorf("list instance types: %v", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintf(tw, "ClusterID\tID\tCPUCores\tCPUModel\tMemoryGB\tGPUCount\tGPUModel\n")
	for _, it := range res.InstanceTypes {
		fmt.Fprintf(tw, "%s\t%s\t%d\t%s\t%d\t%d\t%s\n", it.ClusterID, it.ID, it.CPU.Cores, it.CPU.Model,
			it.Memory.SizeGB, it.GPU.Count, it.GPU.Model)
	}
	return tw.Flush()
}

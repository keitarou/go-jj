package command

import (
	"fmt"
	"strconv"
	"time"
)

type ToDateCommand struct{}

func (c *ToDateCommand) Run(args []string) int {
	value := args[0]
	timestamp, _ := strconv.ParseInt(value, 10, 64)
	t := time.Unix(timestamp, 0)
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	return 0
}
func (c *ToDateCommand) Synopsis() string {
	return "to_date command"
}
func (c *ToDateCommand) Help() string {
	return "Usage: jj to_date unix_timestamp"
}

package command

import (
	"fmt"
	"strings"
	"time"
)

type ToTimestampCommand struct{}

func (c *ToTimestampCommand) Run(args []string) int {
	value := fmt.Sprintf("%s %s", args[0], args[1])
	value = strings.Replace(value, "/", "-", -1)
	t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	fmt.Println(t.Unix())
	return 0
}
func (c *ToTimestampCommand) Synopsis() string {
	return "to_timestamp command"
}
func (c *ToTimestampCommand) Help() string {
	return "Usage: jj to_timestamp yyyy/mm/dd hh:mm:ss"
}

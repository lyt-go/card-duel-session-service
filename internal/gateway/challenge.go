package gateway

import "fmt"

type RemoteError struct {
	Kind    string
	Message string
}

func (e *RemoteError) Error() string { return e.Message }

type Client struct{ Calls int }

func (c *Client) Challenge(mode string) error {
	c.Calls++
	if mode == "reject" {
		return normalize(&RemoteError{Kind: "rejected", Message: "牌组被规则服务拒绝"})
	}
	return nil
}
func normalize(err error) error { return fmt.Errorf("下游调用失败: %v", err) }

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
// normalize 用 %w 包裹，保留 *RemoteError 的错误链，
// 使调用方能通过 errors.As 识别拒绝等具体类型。
func normalize(err error) error { return fmt.Errorf("下游调用失败: %w", err) }

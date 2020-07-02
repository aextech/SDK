package lib

import "ws-api-cli/handle"

type Handle interface {
	GetBody() *handle.Body

	Work()
}

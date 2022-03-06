package response

import "catering/config"

type SysConfigResponse struct {
	Config config.Config `json:"config"`
}

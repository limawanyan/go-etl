package reader

import "github.com/Breeze0806/go-etl/datax/common/plugin"

type Task interface {
	StartRead(reader plugin.RecordSender) error
}

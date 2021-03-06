// +build !go1.7

package pg

import (
	"github.com/mddos/pg/internal/pool"
	"github.com/mddos/pg/orm"
)

// DB is a database handle representing a pool of zero or more
// underlying connections. It's safe for concurrent use by multiple
// goroutines.
type DB struct {
	opt   *Options
	pool  pool.Pooler
	fmter orm.Formatter

	queryProcessedHooks []queryProcessedHook
}

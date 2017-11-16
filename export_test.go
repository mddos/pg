package pg

import "github.com/mddos/pg/internal/pool"

func (db *DB) Pool() pool.Pooler {
	return db.pool
}

func (ln *Listener) CurrentConn() *pool.Conn {
	return ln.cn
}

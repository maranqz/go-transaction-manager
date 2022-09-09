package sqlx

type logger interface {
	Printf(format string, a ...interface{})
}

// WithLog sets logger for transaction.Manager.
func WithLog(l logger) Option {
	return func(m *TrManager) {
		if l == nil {
			l = defaultLog
		}

		m.log = l
	}
}

//nolint:gochecknoglobals // initializing default log, which does nothing
var defaultLog = log{}

type log struct{}

func (l log) Printf(_ string, _ ...interface{}) {}

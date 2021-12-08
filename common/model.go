package common

type Model interface {
	AutoMigrate() error
	Validate() error
}

func AutoMigrate(ms ...Model) {
	for _, m := range ms {
		if err := m.AutoMigrate(); err != nil {
			panic(err)
		}
	}
}

func Validate(m Model) error {
	return m.Validate()
}
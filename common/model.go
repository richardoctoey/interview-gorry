package common

type Model interface {
	AutoMigrate() error
	Validate() error
}

func AutoMigrate(m Model) {
	if err := m.AutoMigrate(); err != nil {
		panic(err)
	}
}

func Validate(m Model) error {
	return m.Validate()
}
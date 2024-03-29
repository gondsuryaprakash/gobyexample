package solidprinciple

/*
Dependency Inversion : 
High level class should not depend on low level classes instead both are depend upon abstraction 

*/

type Logger interface {
	Log(message string) error
}

type FileLogger struct{}

func (f *FileLogger) Log(message string) (err error) {
	// Write log to a file
	return
}

type UserManager struct {
	logger Logger
}

func NewUserManager(logger Logger) *UserManager {
	return &UserManager{
		logger: logger,
	}
}

func (u *UserManager) CreateUser(user *User) {
	// Create user
	u.logger.Log("User created")
}

package user

var AppMode string

func GetAppMode() string {
	return AppMode
}

func IsProduction() bool {
	return AppMode == "production"
}

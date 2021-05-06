package database_repo

//
//go:generate gormgen -structs Database -input .
type Database struct {
	Id       int32  //
	Name     string //
	Addr     string //
	Ports    int32  //
	Account  string //
	Password string //
}

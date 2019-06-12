package models

type Permission struct {

}

func (pm *Permission) TableEngine() string {
    return "INNODB"
}
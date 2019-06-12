package models


type Action struct {

}

func (act *Action) TableEngine() string {
    return "INNODB"
}

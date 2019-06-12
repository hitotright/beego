package models


type WebSet struct {

}

func (ws *WebSet) TableEngine() string {
    return "INNODB"
}
package jiudian2000w

import (
	"errors"
	"fmt"
)

type People struct {
	Name     string `json:"name"`
	CtfId    string `json:"ctf_id"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
	Address  string `json:"address"`
	Mobile   string `json:"mobile"`
	//Tel       string
	Email string `json:"email"`
	//Nation    string
	//Education string
	//Company   string
	//Version   string
}

func (p *People) Type() string {
	return "people"
}

func parsePeopleInfo(record []string) (people People, err error) {
	if len(record) < 26 {
		err = errors.New("该行字段不够：" + fmt.Sprintln(record))
		return
	}
	people = People{
		Name:     record[0],
		CtfId:    record[3],
		Gender:   record[4],
		Birthday: record[5],
		Address:  record[7],
		Mobile:   record[19],
		//Tel:       record[20],
		Email: record[22],
		//Nation:    record[23],
		//Education: record[25],
		//Company:   record[26],
	}
	return
}

func (p People) String() string {
	return fmt.Sprintf("Name: %s Gender: %s Birthday: %s Address: %s Mobile: %s Email: %s Ctfid: %s",
		p.Name, p.Gender, p.Birthday, p.Address, p.Mobile, p.Email, p.CtfId)
}

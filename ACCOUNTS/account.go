package ACCOUNTS

import (
	"encoding/base64"
	"strconv"
	"strings"
	"time"
)

type EmployeeSignIn struct {
	Id          string
	Agent       string
	requestTime time.Time
}

type User struct {
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Mobile        string    `json:"mobile"`
	Password      string    `json:"password"`
	DOB           time.Time `json:"dob"`
	Gender        string    `json:"gender"`
	MaritalStatus string    `json:"marital_status"`
	CreatedAt     time.Time `json:"created_at"`
	Status        string    `json:"status"`
	ArchivedAt    time.Time `json:"archived_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	LastLoginAt   time.Time `json:"last_login_at"`
}

type Return struct {
	Message string `json:"message"`
	Id      string `json:"id"`
	Details *User  `json:"details"`
}

type Auth struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

var Users = []User{}

func (acc User) AddUser() (int, Return) {
	if acc.isAvlable() == -1 {
		acc.Id = "A" + strconv.Itoa((1000000 + len(Users)))
		acc.CreatedAt = time.Now()
		acc.UpdatedAt = time.Now()
		acc.Status = "Active"
		Users = append(Users[0:len(Users)], acc)
		return 200, Return{Message: "User Added Successfully", Id: acc.Id}
	}
	return 400, Return{Message: "User exist"}
}

func (acc User) isAvlable() int {
	for i := range Users {
		if Users[i].Id == acc.Id && Users[i].Status == "Active" {
			return i
		}
	}
	return -1
}

func (acc User) CheckUser(agent string) (int, Auth) {
	id := acc.isAvlable()
	if id != -1 {
		if acc.Password == Users[id].Password {
			var req EmployeeSignIn
			req.Id = acc.Id
			req.Agent = agent
			req.requestTime = time.Now()
			Users[id].LastLoginAt = req.requestTime
			var res = Auth{Message: "User Login Successfully"}
			res.Token = base64.StdEncoding.EncodeToString([]byte(req.Id + "\n" + req.Agent + "\n" + req.requestTime.String()))
			return 200, res
		}
		return 400, Auth{Message: "User Id or Password wrong!"}
	}
	return 400, Auth{Message: "User not Exist!"}
}

func (acc User) GetDetail(agent string, token string) (int, Return) {
	var tokenInfo EmployeeSignIn
	info, err := base64.StdEncoding.DecodeString(token)
	if err == nil {
		str := string(info)
		stringInfo := strings.Split(str, "\n")
		tokenInfo.Id = stringInfo[0]
		tokenInfo.Agent = stringInfo[1]
		timeInfo, er := time.Parse(time.Layout, stringInfo[2])
		if er != nil {
			tokenInfo.requestTime = timeInfo
		}
	}
	acc.Id = tokenInfo.Id
	id := acc.isAvlable()
	if id != -1 {
		if tokenInfo.Agent == agent {
			if tokenInfo.requestTime == Users[id].LastLoginAt {
				acc = Users[id]
				acc.Password = "Headen"
				return 200, Return{Message: "User Details taken Successfully", Details: &acc}
			}
			return 401, Return{Message: "Login Expired!"}
		}
		return 401, Return{Message: "Invalid User!"}
	}
	return 401, Return{Message: "User not Exist!"}
}

func (acc User) UpdateUser() (int, Return) {
	acc.UpdatedAt = time.Now()
	id := acc.isAvlable()
	if id != -1 {
		Users[id] = acc
		return 200, Return{Message: "User Update Successfully"}
	}
	return 400, Return{Message: "User not Exist"}
}

func (acc User) RemoveUser() (int, Return) {
	acc.Status = "Inactive"
	acc.ArchivedAt = time.Now()
	acc.UpdatedAt = time.Now()
	id := acc.isAvlable()
	if id != -1 {
		Users[id] = acc
		return 200, Return{Message: "User Delete Successfully"}
	}
	return 400, Return{Message: "User not Exist"}

}

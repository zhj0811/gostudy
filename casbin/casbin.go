package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//a, _ := xormadapter.NewAdapter("sqlite3", "role.db")
	a, _ := xormadapter.NewAdapterWithTableName("sqlite3", "role.db", "role")

	e, _ := casbin.NewEnforcer("conf/rbac_model.conf", a)

	// Load the policy from DB.
	e.LoadPolicy()
	//e.AddRoleForUser("org1user", "org1")
	//ok, err := e.Enforce("org1user", "data1", "read")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(ok)
	//e.RemovePolicy("org1user", "data1", "read")
	////e.AddPolicy("alice2", "data1", "read")
	// Check the permission.
	e.AddPolicy("alice2", "data1", "read")
	ok, err := e.Enforce("alice2", "data1", "read")

	//fmt.Println(ok)
	//e.RemovePolicy("alice2", "data1", "read")
	//ok, err = e.Enforce("alice2", "data1", "read")

	//ok, err = e.Enforce("org1user", "data1", "read")
	if err != nil {
		panic(err)
	}
	fmt.Println(ok)
	//e.DeleteRoleForUser("org1user", "org1")
	//ok, err = e.Enforce("org1user", "data1", "read")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(ok)
	// Modify the policy.
	//e.AddPolicy(...)
	//e.RemovePolicy(...)

	// Save the policy back to DB.
	//e.SavePolicy()
}

//e.AddPolicy("org1", "data1", "read")
//e.RemovePolicy("org1user", "data1", "read")
//e.AddRoleForUser("org1user", "org1")
//e.DeleteRoleForUser("org1user", "org1")

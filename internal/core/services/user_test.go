package services

import (
	"github.com/a-korkin/ecommerce/internal/core/models"
	"testing"
)

func TestCreateUser(t *testing.T) {
	in := models.UserIn{
		LastName:  "Smith",
		FirstName: "John",
	}
	out, err := userService.Create(&in)
	if err != nil {
		t.Errorf("Failed to create user: %s", err)
	}
	if out.LastName != "Smith" || out.FirstName != "John" {
		t.Errorf("Wrong user created: %v", out)
	}
}

func TestUpdateUser(t *testing.T) {
	in := models.UserIn{
		LastName:  "upd_last_name",
		FirstName: "upd_first_name",
	}
	id := "d3f729cb-43c0-40c4-9084-74fb2b0bd408"
	out, err := userService.Update(id, &in)
	if err != nil {
		t.Errorf("Failed to update user: %s", err)
	}
	if out.LastName != "upd_last_name" || out.FirstName != "upd_first_name" {
		t.Errorf("Wrong user updated: %v", out)
	}
}

func TestGetAllUsers(t *testing.T) {
	pageParams := models.NewPageParams("")
	out, err := userService.GetAll(pageParams)
	if err != nil {
		t.Errorf("Failed to get all users: %s", err)
	}
	if len(out) < 2 {
		t.Errorf("Wrong count of users returned")
	}
}

func TestGetUserByID(t *testing.T) {
	id := "4636a25d-02ee-4eb8-9757-efd677677076"
	out, err := userService.GetByID(id)
	if err != nil {
		t.Errorf("Failed to get user by id: %s", err)
	}
	if out.LastName != "Ivanov" || out.FirstName != "Ivan" {
		t.Errorf("Wrong user returned: %v", out)
	}
}

func TestDeleteUser(t *testing.T) {
	id := "5e782875-4d9c-4641-be3c-afddeb05c083"
	if err := userService.Delete(id); err != nil {
		t.Errorf("Failed to delete user: %s", err)
	}
}

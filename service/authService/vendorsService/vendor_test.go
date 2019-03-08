package vendorsService

import (
	"../../../model"

	"fmt"
	"testing"
)

func TestCreateVendor(t *testing.T) {
	vendor, err := CreateVendor(&model.Vendor{
		UserName: "test1",
		Password: "admin1234",
	})
	fmt.Println("user:", vendor, "err:", err)
}

func TestReadVendor(t *testing.T) {
	vendor, err := ReadVendor(6)
	fmt.Println("user:", vendor, "err:", err)
}

func TestUpdateVendor(t *testing.T) {
	vendor, err := UpdateVendor(&model.Vendor{
		ID:       6,
		Name:     "xiang",
		UserName: "tian",
	})
	fmt.Println("user:", vendor, "err:", err)
}

func TestDeleteVendor(t *testing.T) {
	err := DeleteVendor(6)
	fmt.Println("err:", err)
}

func TestReadVendors(t *testing.T) {
	//	vendors, total, err := ReadVendors("admin", 0, 0, "", 0)
	//	fmt.Println("users:", vendors, "total:", total, "err:", err)
	//	ReadByField("ShenYang", 0, 0, "city", 0)
}

package service

import(
	"project01/app/model"
	"fmt"
	"log"
	"context"
)


func AddKategori(conn model.InventoryClient) {
	var (
		idKategori int32 =0
		namaKategori string
	)

	fmt.Println("Masukkan nama kategori")
	fmt.Scan(&namaKategori)


	req := &model.Kategori{
		IdKategori :idKategori,
		NamaKategori: namaKategori,
	}

	res,err := conn.AddKategori(context.Background(), req)
	if err != nil {
		log.Fatalf("Tidak bisa menerima response terkait add Kategori karena error: ", err)
	}

	fmt.Println("Status Anda adalah ", res.GetStatus())
	fmt.Println("Message Anda adalah ", res.GetMessage())
}

func EditKategori(conn model.InventoryClient){
	var(
		idKategori int32
		namaKategori string
	)

	fmt.Println("Masukkan id kategori")
	fmt.Scan(&idKategori)

	fmt.Println("Masukkan nama Kategori")
	fmt.Scan(&namaKategori)

	req := &model.Kategori{
		IdKategori: idKategori,
		NamaKategori : namaKategori,
	}

	res,err := conn.EditKategori(context.Background(), req)
	if err != nil {
		log.Fatalf("Tidak bisa menerima response terkait Edit Kategori dikarenaka karena error ", err)
	}

	fmt.Println("Status Anda adalah ", res.GetStatus())
	fmt.Println("Message Anda adalah ", res.GetMessage())
}

func DeleteKategori(conn model.InventoryClient) {
	var (
		idKategori int32
	)

	fmt.Println("Masukkan id Kategori")
	fmt.Scan(&idKategori)

	req := &model.Kategori {
		IdKategori: idKategori,
	}

	res,err := conn.DeleteKategori(context.Background(), req)
	if err != nil {
		log.Fatalf("Tidak bisa menerima response terkait Delete Kategori dikarenakan error ", err)
	}

	fmt.Println("Status Anda adalah ", res.GetStatus())
	fmt.Println("Message Anda adalah ", res.GetMessage())
}

func ShowKategori(conn model.InventoryClient) {
	res,err := conn.ShowKategori(context.Background(),new(model.Empty))
	if err != nil {
		log.Fatalf("Tidak bisa menerima response terkait Show Kategori ", err)
	}


	if len(res.KategoriList) < 1 {
		fmt.Println("Belum ada kategori yang disimpan")
	} else {
		fmt.Println("Barang yang telah diSimpan: ")
		for _,val:= range res.KategoriList {
			fmt.Println("=======================")
			fmt.Println("Id Kategori: ", val.IdKategori)
			fmt.Println("Nama Kategori: ", val.NamaKategori)
		}
	}

}

func MenuKategori (conn model.InventoryClient) {

	Menu()
	var flag bool = false
	var angkaSwitch int

	for !flag {
		
		fmt.Println("Masukkan angka pilihan")
		fmt.Println(&angkaSwitch)
		
		switch angkaSwitch {
		case 0:
			Menu()
			break
		case 1:
			AddKategori(conn)
			break
		case 2:
			EditKategori(conn)
			break
		case 3:
			DeleteKategori(conn)
			break
		case 4:
			ShowKategori(conn)
			break
		case 5:
			fmt.Println("Keluar Menu Kategori")
			flag = true
			break

		}
	}
}

func Menu(){
	fmt.Println("0. Menu Kategori")
	fmt.Println("1. Add Kategori")
	fmt.Println("2. Edit Kategori")
	fmt.Println("3. Delete Kategori")
	fmt.Println("4. Show Kategori")
	fmt.Println("5. Keluar menu")
}
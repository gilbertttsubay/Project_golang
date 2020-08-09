package service

import (
  "context"
  "project01/app/model"
  "fmt"
  "net"
  "log"
  "google.golang.org/grpc"

  // "project01/app/config"
)

var Inventory *model.Item
var InventoryList *model.ItemList

var InventoryListShowAll *model.ItemList



func init()  {
  Inventory = new(model.Item)
  InventoryList = new(model.ItemList)
  InventoryListShowAll = new(model.ItemList)
  InventoryListShowAll.ItemList =make([]*model.Item, 0)
  InventoryList.ItemList = make([]*model.Item, 0)
}

func idFilterInventory (idItem int32) *model.Item {
  for i := 0; i < len(InventoryList.ItemList); i++ {
    InventoryDariList := InventoryList.ItemList[i]
    if idItem == InventoryDariList.IdItem {
      return InventoryDariList
    } else {
      return nil
    }
  }
  return nil
}

func IdItemFilter (idItem int32) int32 {
  for i := 0; i < len(InventoryList.ItemList); i++ {
    InventoryDariList := InventoryList.ItemList[i]
      if idItem == InventoryDariList.IdItem {
        return 1
      } else {
        return 2
      }
    
  }
  return 0
}

func removeInventory (slice []*model.Item, s int32) []*model.Item {
  return append(slice[:s], slice[s+1:]...)

}

func main(){
  fmt.Println("Hello")

	lis, error := net.Listen("tcp", "0.0.0.0:50051")

	if error != nil {
		log.Fatal("Failed to listen: %v", error)
	}

	s := grpc.NewServer()

  model.RegisterInventoryServer(s, &Server{})



	if error := s.Serve(lis); error != nil {
		log.Fatalf("failed to serve: %v", error)
	}
}

type Server struct{}

func (*Server) AddItem(ctx context.Context, item *model.Item) (*model.Status, error) {
  hasilFilterId := idFilterInventory(item.GetIdItem())

  fmt.Println(item.GetIdItem())

  if hasilFilterId != nil {
    res := &model.Status {
      Status: 400,
      Message: "Id Item sudah ada yang sama",
    }

    return res, nil

  } else if hasilFilterId == nil{
    Inventory.NamaItem = item.GetNamaItem()
    Inventory.IdItem = item.GetIdItem()
    Inventory.Jumlah = item.GetJumlah()
    Inventory.Kategori = item.GetKategori()

    InventoryList.ItemList = append(InventoryList.ItemList, Inventory)

    res := &model.Status {
      Status : 200,
      Message : "Barang Berhasil Disimpan",
    }

    return res, nil
  } 
  res := &model.Status {
    Status : 404,
    Message : "Not Found",
  }
  return res, nil
}

func (*Server) GetItem(ctx context.Context, item *model.Item) (*model.Item, error) {
  InventoryDariClient := idFilterInventory(item.GetIdItem())
  var err error
  if InventoryDariClient != nil {
    res:= &model.Item {
      IdItem: InventoryDariClient.IdItem,
      NamaItem: InventoryDariClient.NamaItem,
      Jumlah: InventoryDariClient.Jumlah,
      Kategori: InventoryDariClient.Kategori,
    }

    removeInventory(InventoryList.ItemList, item.GetIdItem())

    return res, nil
  } else {
    res:= &model.Item {
      IdItem:0,
    }
    return res, err
  }

}

func(*Server) Show(ctx context.Context, item *model.Item) (*model.Item, error) {
  InventoryDariClient := idFilterInventory(item.GetIdItem())
  var err error
  if InventoryDariClient != nil {
    res:= &model.Item {
      IdItem: InventoryDariClient.IdItem,
      NamaItem: InventoryDariClient.NamaItem,
      Jumlah: InventoryDariClient.Jumlah,
      Kategori: InventoryDariClient.Kategori,
    }

    removeInventory(InventoryList.ItemList, item.GetIdItem())

    return res, nil
  } else {
    res :=&model.Item {
      IdItem: 0,
    }
    return res, err
  }
}

func(*Server) ShowAll(ctx context.Context, item *model.Item) (*model.ItemList, error) {
  idUser := item.GetIdUser()

  
  if idUser > 0{
    for i := 0; i < len(InventoryList.ItemList); i++ {
      
      InventoryDariClient:= InventoryList.ItemList[i]
      if idUser == InventoryDariClient.IdUser {
        InventoryListShowAll.ItemList = append(InventoryListShowAll.ItemList, InventoryDariClient)
      }

    }
    res:= &model.ItemList{
      ItemList:InventoryListShowAll.ItemList,
    }
    return res, nil
  }

  res:= &model.ItemList{
    ItemList:nil,
  }
  return res,nil

}




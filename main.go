package main

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocarina/gocsv"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Vehicleinfo struct {
	gorm.Model
	ID                   uint   `gorm:"id"`
	Image_Name           string `gorm:"image_name"`
	Date                 string `gorm:"date"`
	Time                 string `gorm:"time"`
	License_Plate_Column string `gorm:"license_plate_column"`
	Output               string `gorm:"output"`
	NAS_Image_Path       string `gorm:"nas_image_path"`
	Camera_Name          string `gorm:"camera_name"`
	Brand                string `gorm:"brand"`
	with_helmet          uint32 `gorm:"with_helmet"`
}

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

var db *sql.DB

func createTable() {

	query := `create table vehicleinfo (
		image_name  varchar(256) not null,
		date  datetime not null,
		time  datetime not null,
		license_plate_column varchar(256) not null,
		output  varchar(256) not null,
		nas_image_path  varchar(256) not null,
		camera_name  varchar(256) not null,
		brand  varchar(256) not null,
         with_helmet   int not null
			 );`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	//fmt.Println(" table created....")
}

func main() {
	//createTable()
	file, err := os.Open("C:/Users/LENOVO/Desktop/vehicleinfo.csv")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	var vehicleinfo []Vehicleinfo
	err = gocsv.Unmarshal(file, &vehicleinfo)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/weather"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&Vehicleinfo{})
	if err != nil {
		panic(err)
	}

	result := db.Create(vehicleinfo)
	if result.Error != nil {
		panic(result.Error)
	}

}

package main

import (
	"database/sql"
	"log"
	"regexp"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type GeneLoci struct {
	RSID         string `json:"RSID"`
	Genotype     string `json:"genotype"`
	ExtendedData int    `json:"extended_data"`
}

func InitDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./genetic.db")
	if err != nil {
		panic(err)
	}
	return db
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {
    gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(Cors())
	v := r.Group("api/")
	v.GET("/rsid/:rsid", GetRsid)
	r.Run(":8080")
}

func GetRsid(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	rsid := c.Params.ByName("rsid")
	matched, _ := regexp.MatchString(`^[rs\d]+$`, rsid)
	if !matched {
		c.JSON(404, gin.H{"error": "invalid rsid"})
		return
	}
	rows, err := db.Query("select rsid,Genotype,extended_data from genetic where rsid='" + rsid + "'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var geneLocis []GeneLoci
	for rows.Next() {
		var geneLoci GeneLoci
		err = rows.Scan(&geneLoci.RSID, &geneLoci.Genotype, &geneLoci.ExtendedData)
		if err != nil {
			log.Fatal(err)
		}
		geneLocis = append(geneLocis, geneLoci)
	}
	log.Printf("try find rsid:%s", rsid)
	if len(geneLocis) != 1 {
		c.JSON(200, geneLocis)
    } else if len(geneLocis) == 1{
        c.JSON(200,geneLocis[0])
	} else {
		c.JSON(404, gin.H{"error": "rsid not found"})
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Provinsi struct {
	IdProvinsi   string `json:"id_provinsi"`
	NamaProvinsi string `json:"nama_provinsi"`
}

type Kota struct {
	IdKota   string `json:"id_kota"`
	NamaKota string `json:"nama_kota"`
}

type Kecamatan struct {
	IdKecamatan   string `json:"id_kecamatan"`
	NamaKecamatan string `json:"nama_kecamatan"`
}

type Faskes struct {
	IdPuskes        string `json:"id_puskes"`
	NamaPuskesmas   string `json:"nama_puskesmas"`
	AlamatPuskesmas string `json:"alamat_puskesmas"`
	NoTelpPuskesmas string `json:"no_telp_puskesmas"`
	IdProvinsi   string `json:"id_provinsi"`
	IdKota   string `json:"id_kota"`
	IdKecamatan   string `json:"id_kecamatan"`
}

type Health struct {
	Status string
}

var Provinsis []Provinsi
var Kotas []Kota
var Kecamatans []Kecamatan
var Faskess []Faskes
//var Faskess map[string]interface{}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Welcome to the homepage")
}

func ProvinsiFunction(w http.ResponseWriter, r *http.Request){
	log.Printf("Endpoint Hit: returnAllProvinsi")
	json.NewEncoder(w).Encode(Provinsis)
}

func KotaFunction(w http.ResponseWriter, r *http.Request){
	log.Printf("Endpoint Hit: returnAllKota")
	Kotas = []Kota{
		Kota{IdKota: "210", NamaKota: "Peru"},
		Kota{IdKota: "220", NamaKota: "Barujaya"},
	}
	//json.NewEncoder(w).Encode(Kotas)
	resp, err := json.Marshal(Kotas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "aplication/json")
	w.Write(resp)
}

func KecamatanFunctions(w http.ResponseWriter, r *http.Request){
	log.Printf("Endpoint Hit: returnAllKecamatan")
	//json.NewEncoder(w).Encode(Kecamatans)
	resp, err := json.Marshal(Kecamatans)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func FaskesFunction(w http.ResponseWriter, r *http.Request){
	log.Printf("Endpoint Hit: returnAllFaskes")
	json.NewEncoder(w).Encode(Faskess)
}

func HealthCheck(w http.ResponseWriter, r *http.Request){
	s := Health{"Good"}
	resp, err := json.Marshal(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}


func handleRequest(){
	// create a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/faskes", FaskesFunction)
	myRouter.HandleFunc("/kecamatan", KecamatanFunctions)
	myRouter.HandleFunc("/kota", KotaFunction)
	myRouter.HandleFunc("/provinsi", ProvinsiFunction)
	myRouter.HandleFunc("/health",HealthCheck)
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main() {
	fmt.Println("Rest Api v2.0 - Mux Router - moehiqbal golang")
	Kecamatans = []Kecamatan{
		Kecamatan{IdKecamatan: "110", NamaKecamatan: "Pesanggarahan"},
		Kecamatan{IdKecamatan: "120", NamaKecamatan: "Barujaya"},
	}
	Provinsis = []Provinsi{
		Provinsi{IdProvinsi: "310", NamaProvinsi: "Bimonten West"},
		Provinsi{IdProvinsi: "320", NamaProvinsi: "Bimonten Center"},
	}
	Faskess = []Faskes{
		Faskes{IdPuskes: "001", NamaPuskesmas: "PKC Bimonten jaya", AlamatPuskesmas: "Bimonten Center",
			NoTelpPuskesmas: "0895107723412", IdProvinsi: "320", IdKota: "220", IdKecamatan: "120"},
	}

	//Faskess = map[string]interface{}{
	//
	//}
	handleRequest()
}

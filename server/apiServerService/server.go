package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Satici struct {
	Id             string
	Name           string `json:"name"`
	Foto           string `json:"foto"`
	MainLocation   string `json:"mainlocation"`
	BazaarLocation string `json:"bazaarlocation"`
	BazaarName     string `json:"bazaarname"`
	Crops          string `json:"crops"`
	Phone          string `json:"phone"`
	Days           string `json:"days"`
}

type SaticiHandlers struct {
	sync.Mutex
	store map[string]Satici
}

func (h *SaticiHandlers) saticilar(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method desteklenmiyor.."))
		return

	}
}

func (h *SaticiHandlers) get(w http.ResponseWriter, r *http.Request) {
	saticilar := make([]Satici, len(h.store))

	h.Lock()
	i := 0
	for _, satici := range h.store {
		saticilar[i] = satici
		i++
	}
	h.Unlock()

	jsonBytes, err := json.Marshal(saticilar)
	if err != nil {
		//TODO
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *SaticiHandlers) getSatici(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	h.Lock()

	satici, ok := h.store[parts[2]]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	h.Unlock()

	jsonBytes, err := json.Marshal(satici)
	if err != nil {
		//TODO
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *SaticiHandlers) post(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		//TODO
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	//content-type
	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("content-type: 'application/json' olmalıydı fakat sen bunu kullandın:,%s", ct)))
		return
	}

	var satici Satici
	if err := json.Unmarshal(bodyBytes, &satici); err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(err.Error()))
	}

	satici.Id = fmt.Sprintf("%d", time.Now().UnixNano())

	h.Lock()
	h.store[satici.Id] = satici
	defer h.Unlock()

}

func newSaticiHandlers() *SaticiHandlers {
	return &SaticiHandlers{
		store: map[string]Satici{
			"id1": Satici{
				Name:           "abdullah admin",
				Foto:           "none",
				MainLocation:   "SAMSUN ATAKUM ÖMÜREVLERİ CUMHURİYET MAHALLESİ",
				BazaarLocation: "/Samsun/Atakum",
				BazaarName:     "Ömürevleri_Pazarı,Denizevleri_Pazarı",
				Crops:          "Marul,Maydanoz,YeşilBiber,Karpuz,Elma,Çilek",
				Phone:          "05301215386",
				Days:           "1458",
			},
		},
	}
}

type adminPortal struct {
	password string
}

func newAdminPortal() *adminPortal {
	password := "secret"
	if password == "" {
		panic("ADMIN_PASSWORD gerekli, set edilmemiş ADMIN_PASSWORD durumu var")
	}

	return &adminPortal{password: password}
}

func (a adminPortal) handler(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if !ok || user != "admin" || pass != a.password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - unauthorized"))
		return
	}

	w.Write([]byte("<html><h1>SUPER - ADMIN PORTALI</h1></html>"))

}

func main() {
	admin := newAdminPortal()

	SaticiHandlers := newSaticiHandlers()

	http.HandleFunc("/saticilar", SaticiHandlers.saticilar)
	http.HandleFunc("/saticilar/", SaticiHandlers.getSatici)
	http.HandleFunc("/admin", admin.handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)

	}

	fmt.Print("Server is runnning...")
}

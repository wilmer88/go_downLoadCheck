package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/wilmer88/go_downLoadCheck/pkg/utils"
	"github.com/wilmer88/go_downLoadCheck/pkg/models"
)
 var NewMember models.Fammember

 func GetMembers(w http.ResponseWriter, r *http.Request){
	NewMember:= models.GetALLMembers()
	res, _ :=json.Marshal(NewMember)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
 }

 func GetFamMemById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	famMem:= vars["famMem"]
	FamID, err:= strconv.ParseInt(famMem,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	familyMdetails, _ := models.GetMemberByID(FamID)
	res, _ := json.Marshal(familyMdetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
 }

 func CreatefamiliaMiembro(w http.ResponseWriter, r *http.Request){
	CreateFM := &models.Fammember{}
	utils.ParseBody(r, CreateFM)
	b:= CreateFM.CreateMember()
	res,_:= json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
 }

 func eleminaMember(w http.ResponseController, r *http.Request){
	vars := mux.vars(r)
	memId := vars["famID"]
	ID, err := strconv.ParseInt(memId,0,0)
	 err != nil {
		fmt.Println("error while parsing delete")
	}
	fmember := models.DeleteMember(ID)
	res, _ := json.Marshal(fmember)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
 }

 fun UpdateFamMember(w http.ResponseWriter, r *http.Request){
	var updateMember = &models.Fammember{}
	utils.Parse(r, updateMember)
	vars := mux.Vars(r)
	memId : vars["famID"]
	ID, err := strconv.ParseInt(memId,0,0)
	 err != nil {
		fmt.Println("error while parsing update")
    }
	memberDetails, db:=models.GetFamMemById(ID)
	if updateMember.FirstName != ""{
		memberDetails.FirstName = updateMember.FirstName
	}
	if updateMember.Happiness != ""{
		memberDetails.Happiness = updateMember.Happiness
	}
	if updateMember.UrlStr != ""{
		memberDetails.UrlStr = updateMember.UrlStr
	}

	db.Save(&memberDetails)
	res, _:= json.Marshal(memberDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	
}
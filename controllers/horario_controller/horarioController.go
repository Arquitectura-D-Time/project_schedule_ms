package horario_controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	driver "project_schedule_ms/common"
	repository "project_schedule_ms/data"
	horario "project_schedule_ms/data/horario_mysql"
	model "project_schedule_ms/model"
)

func NewHorarioHandler(db *driver.DB) *Horario {
	return &Horario{
		repo: horario.NewSQLHorario(db.SQL),
	}
}

type Horario struct {
	repo repository.Horario
}

func (p *Horario) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := p.repo.Fetch(r.Context(), 5)

	respondwithJSON(w, http.StatusOK, payload)
}

func (p *Horario) Create(w http.ResponseWriter, r *http.Request) {
	horario := model.Horario{}
	json.NewDecoder(r.Body).Decode(&horario)

	newID, err := p.repo.Create(r.Context(), &horario)
	fmt.Println(newID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
	fmt.Println(err)
}

func (p *Horario) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "IDtutoria"))
	data := model.Horario{IDtutoria: int64(id)}
	json.NewDecoder(r.Body).Decode(&data)
	payload, err := p.repo.Update(r.Context(), &data)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

func (p *Horario) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "IDtutoria"))
	fmt.Println(id)
	payload, err := p.repo.GetByID(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

func (p *Horario) GetByNombre(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "NombreMateria"))
	fmt.Println(id)
	payload, err := p.repo.GetByNombre(r.Context(), string(id))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

func (p *Horario) GetByFecha(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "Fecha"))
	fmt.Println(id)
	payload, err := p.repo.GetByFecha(r.Context(), string(id))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

func (p *Horario) GetByHora(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "HoraInicio"))
	fmt.Println(id)
	payload, err := p.repo.GetByHora(r.Context(), string(id))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

func (p *Horario) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "IDtutoria"))
	_, err := p.repo.Delete(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
package models

type CriminalRecord struct {
	ID               int    `json:"id"`
	SuspectName      string `json:"suspect_name"`
	CrimeDescription string `json:"crime_description"`
	CrimeDate        string `json:"crime_date"`
	CrimeStatus      string `json:"crime_status"`
	PrisonID         int    `json:"prison_id"`
}

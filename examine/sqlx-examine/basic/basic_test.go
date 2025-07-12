package basic

import (
	"testing"
	// "github.com/jmoiron/sqlx"
)

func TestBasic(t *testing.T) {

	db, err := GetDB()
	if err != nil {
		t.Fatalf("failed to establish dstabase connection: %v", err)
	}
	defer db.Close()

	db.MustExec(schema)

	tx := db.MustBegin()
	tx.MustExec(insertLocation, 1, "Dubuque, Iowa", "Millwork District")
	tx.MustExec(insertLocation, 2, "Dubuque, Iowa", "Valley View")
	tx.NamedExec(insertLocationNamed, &Location{Id: 3, Addr: "Dubuque, Iowa", Name: "Center Grove"})
	tx.MustExec(insertPlace, 1, 1, "Charlotte's Coffee House")
	tx.MustExec(insertPlace, 2, 1, "Backpocket Dubuque")
	tx.NamedExec(insertPlaceNamed, &Place{Id: 3, LocationId: 1, Name: "Driftless Pizza"})
	tx.MustExec(insertPlace, 4, 2, "Trees Plus")
	tx.MustExec(insertPlace, 5, 2, "Spear Logistics")
	tx.NamedExec(insertPlaceNamed, &Place{Id: 6, LocationId: 3, Name: "Hampton Inn Dubuque"})
	tx.NamedExec(insertPlaceNamed, &Place{Id: 7, LocationId: 3, Name: "Warren Plaza"})
	tx.Commit()

	locations := []Location{}
	err = db.Select(&locations, selectLocations)
	if err != nil {
		t.Errorf("failed to fetch locations: %v", err)
	} else {
		if actual, expected := len(locations), 3; actual != expected {
			t.Errorf("locations count: get %d, expect %d", actual, expected)
		}
	}

	places := []Place{}
	err = db.Select(&places, "SELECT id, location_id, name from place where location_id=$1 ORDER BY name", 2)
	if err != nil {
		t.Errorf("failed to fetch places: %v", err)
	} else {
		if actual, expected := places[0].Name, "Spear Logistics"; actual != expected {
			t.Errorf("First place in Valley View: get %q, expect %q", actual, expected)
		}
	}

	db.MustExec(cleanup)

}

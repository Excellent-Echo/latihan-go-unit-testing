package helper

import "testing"

// Unit Testing latihan 1
func TestTotal(t *testing.T) {
	t.Run("Test Total hasil dari Pajak PPN", func(t *testing.T) {
		got1 := Total(110000)
		want1 := 10000

		got2 := Total(220000)
		want2 := 20000

		t.Logf("Test : %d", want1)
		t.Logf("Test : %d", want2)

		if got1 != want1 {
			t.Errorf("got1 %d, but want1 %d", got1, want1)
		}

		if got2 != want2 {
			t.Errorf("got2 %d, but want2 %d", got2, want2)
		}
	})
}

// Unit testing latihan 2
func TestChangeToIdr(t *testing.T) {
	t.Run("Test Total hasil convert rupiah", func(t *testing.T) {
		got1 := ChangeToIdr(100000)
		want1 := "100,000"

		got2 := ChangeToIdr(1200000)
		want2 := "1,200,000"

		t.Logf("Test : %s", want1)
		t.Logf("Test : %s", want2)

		if got1 != want1 {
			t.Errorf("got1 %s, but want1 %s", got1, want1)
		}

		if got2 != want2 {
			t.Errorf("got2 %s, but want2 %s", got2, want2)
		}
	})
}

// Unit testing latihan 3
func TestGenapGanjil(t *testing.T) {
	t.Run("Test Total hasil ganjil genap", func(t *testing.T) {
		got1 := GenapGanjil(1, 2, 3, 4, 5)
		want1 := "Angka terbanyak adalah ganjil"

		got2 := GenapGanjil(4, 2)
		want2 := "Angka terbanyak adalah genap"

		t.Logf("Test : %s", want1)
		t.Logf("Test : %s", want2)

		if got1 != want1 {
			t.Errorf("got1 %s, but want1 %s", got1, want1)
		}

		if got2 != want2 {
			t.Errorf("got2 %s, but want2 %s", got2, want2)
		}
	})
}

// Unit testing latihan 4
func TestGovermentHelper(t *testing.T) {
	t.Run("Test Total hasil dapat bantuan atau tidak", func(t *testing.T) {
		var data = map[string]string{
			"name":        "andi",
			"umur":        "30",
			"jarakRumah":  "50",
			"berkeluarga": "ya",
		}

		var data2 = map[string]string{
			"name":        "santi",
			"umur":        "19",
			"jarakRumah":  "80",
			"berkeluarga": "ya",
		}
		got1 := GovermentHelper(data)
		want1 := "andi layak mendapat bantuan dari pemerintah"

		got2 := GovermentHelper(data2)
		want2 := "santi tidak layak mendapat bantuan dari pemerintah"

		t.Logf("Test : %s", want1)
		t.Logf("Test : %s", want2)

		if got1 != want1 {
			t.Errorf("got1 %s, but want1 %s", got1, want1)
		}

		if got2 != want2 {
			t.Errorf("got2 %s, but want2 %s", got2, want2)
		}
	})
}

// Unit testing latihan 6
func TestTotalPajak(t *testing.T) {
	t.Run("Test Total hasil PPN", func(t *testing.T) {
		var belanja1 = map[string]int{
			"sepatu":   1100000,
			"jaket":    2200000,
			"topi":     594000,
			"celana":   803000,
			"sweater":  330000,
			"kausKaki": 110000,
			"sabuk":    55000,
		}

		got1 := TotalPajak(belanja1)
		want1 := "total PPN yang diterima sebesar 472000"

		t.Logf("Test : %s", want1)

		if got1 != want1 {
			t.Errorf("got1 %s, but want1 %s", got1, want1)
		}

	})
}

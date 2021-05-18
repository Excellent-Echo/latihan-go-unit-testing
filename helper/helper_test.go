package helper

import "testing"

// Unit Testing latihan 1
func TestTotal(t *testing.T) {
	t.Run("Test perhitungan biaya PPN dan pengeluaran saat belanja 110000", func(t *testing.T) {
		got1, got2 := Total(110000)
		want1, want2 := 10000, 100000

		t.Logf("Test : %v", want1)
		t.Logf("Test : %v", want2)

		if got1 != want1 && got2 != want2 {
			t.Errorf("got1 %v and got1 %v, but want1 %v and want2 %v", got1, got2, want1, want2)
		}

	})

	t.Run("Test perhitungan biaya PPN dan pengeluaran saat belanja 220000", func(t *testing.T) {
		got1, got2 := Total(220000)
		want1, want2 := 20000, 200000

		t.Logf("Test : %v", want1)
		t.Logf("Test : %v", want2)

		if got1 != want1 && got2 != want2 {
			t.Errorf("got1 %v and got1 %v, but want1 %v and want2 %v", got1, got2, want1, want2)
		}

	})

	t.Run("Test perhitungan biaya PPN dan pengeluaran tidak boleh nol", func(t *testing.T) {
		got1, got2 := Validate(0)
		want := "tidak boleh nol"

		t.Logf("Test : %s", want)

		if got1 != false && got2 != nil {
			t.Errorf("got1 %t, got2 %t but want %s", got1, got2, want)
		}
	})
}

// Unit testing latihan 2
func TestChangeToIdr(t *testing.T) {
	t.Run("Test Total hasil convert rupiah 100000", func(t *testing.T) {
		got := ChangeToIdr(100000)
		want := "100,000"

		t.Logf("Test : %s", want)

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}

	})

	t.Run("Test Total hasil convert rupiah 1200000", func(t *testing.T) {
		got := ChangeToIdr(1200000)
		want := "1,200,000"

		t.Logf("Test : %s", want)

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})

	t.Run("Test konversi dari nominal menjadi nilai rupiah tidak boleh tanpa value", func(t *testing.T) {
		got1, got2 := ValidateChangeToIdr(0)
		want := "tidak boleh tanpa value"

		t.Logf("Test : %s", want)

		if got1 != false && got2 != nil {
			t.Errorf("got1 %t, got2 %t but want %s", got1, got2, want)
		}
	})
}

// Unit testing latihan 3
func TestGenapGanjil(t *testing.T) {
	t.Run("Test Total hasil ganjil genap deretan 1, 2, 3, 4, 5", func(t *testing.T) {
		got := GenapGanjil(1, 2, 3, 4, 5)
		want := "Angka terbanyak adalah ganjil"

		t.Logf("Test : %s", want)

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}

	})

	t.Run("Test Total hasil ganjil genap deretean 4, 2", func(t *testing.T) {
		got := GenapGanjil(4, 2)
		want := "Angka terbanyak adalah genap"

		t.Logf("Test : %s", want)

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})

	t.Run("Isi deretean angka tidak boleh kosong", func(t *testing.T) {
		got1, got2 := ValidateGanjilGenap(0, 0)
		want := "tidak boleh kosong"

		t.Logf("Test : %s", want)

		if got1 != false && got2 != nil {
			t.Errorf("got1 %t, got2 %t but want %s", got1, got2, want)
		}
	})
}

// Unit testing latihan 4
func TestGovermentHelper(t *testing.T) {
	t.Run("Test Total hasil dapat bantuan atau tidak dari andi", func(t *testing.T) {
		var data = map[string]string{
			"name":        "andi",
			"umur":        "30",
			"jarakRumah":  "50",
			"berkeluarga": "ya",
		}

		got := GovermentHelper(data)
		want := "andi layak mendapat bantuan dari pemerintah"

		t.Logf("Test : %s", want)

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})

	t.Run("Test Total hasil dapat bantuan atau tidak dari santi", func(t *testing.T) {
		var data = map[string]string{
			"name":        "santi",
			"umur":        "19",
			"jarakRumah":  "80",
			"berkeluarga": "ya",
		}

		got := GovermentHelper(data)
		want := "santi tidak layak mendapat bantuan dari pemerintah"

		t.Logf("Test : %s", want)

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	})
}

// Unit testing latihan 6
func TestTotalPajak(t *testing.T) {
	t.Run("Test Total hasil PPN", func(t *testing.T) {
		var belanja = map[string]int{
			"sepatu":   1100000,
			"jaket":    2200000,
			"topi":     594000,
			"celana":   803000,
			"sweater":  330000,
			"kausKaki": 110000,
			"sabuk":    55000,
		}

		got := TotalPajak(belanja)
		want := "total PPN yang diterima sebesar 472000"

		t.Logf("Test : %s", want)

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}

	})
}

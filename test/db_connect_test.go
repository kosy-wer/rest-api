package tests

import (
	"context"
	"testing"
	"rest_api/internal/apps/database"
  "rest_api/internal/apps/register/model/domain"
  "rest_api/internal/apps/register/repository"
	"database/sql"
)

func TestPostgresConnection(t *testing.T) {
	db, err := database.GetConnection()
	if err != nil {
		t.Fatalf("❌ Gagal konek: %v", err)
	}
	defer db.Close()

	var version string
	err = db.QueryRowContext(context.Background(), "SELECT version()").Scan(&version)
	if err != nil {
		t.Fatalf("❌ Query gagal: %v", err)
	}

	t.Logf("✅ Koneksi sukses! PostgreSQL version: %s", version)
}

func TestSelectUsers(t *testing.T) {
	db, err := database.GetConnection()
	if err != nil {
		t.Fatalf("❌ Gagal konek: %v", err)
	}
	defer db.Close()

	rows, err := db.QueryContext(context.Background(), `SELECT id, "Name" FROM students LIMIT 5`)

	if err != nil {
		t.Fatalf("❌ Query SELECT gagal: %v", err)
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		var id int
		var Name string
		if err := rows.Scan(&id, &Name); err != nil {
			t.Fatalf("❌ Gagal scan row: %v", err)
		}
		t.Logf("👤 User: ID=%d, Name=%s", id, Name)
		count++
	}

	if count == 0 {
		t.Logf("⚠️ Tidak ada data di tabel users")
	} else {
		t.Logf("✅ Berhasil baca %d user(s)", count)
	}
}
func TestInsertStudent(t *testing.T) {
    db, err := database.GetConnection()
    if err != nil {
        t.Fatalf("❌ Gagal konek: %v", err)
    }
    defer db.Close()

    tx, err := db.Begin()
    if err != nil {
        t.Fatalf("❌ Gagal mulai transaksi: %v", err)
    }
    defer tx.Rollback() // rollback biar data gak nyampah

    repo := repository.NewUserRepository()
    student := domain.Student{
        Name:  "Test User",
        Email: "test@example.com",
    }

    result := repo.Save(context.Background(), tx, student)

    if result.Name != student.Name || result.Email != student.Email {
        t.Errorf("❌ Insert gagal: %+v", result)
    } else {
        t.Logf("✅ Insert sukses: %+v", result)
    }
}

func getTx(t *testing.T) *sql.Tx {
	db, err := database.GetConnection()
	if err != nil {
		t.Fatalf("❌ Gagal konek DB: %v", err)
	}
	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("❌ Gagal mulai transaksi: %v", err)
	}
	return tx
}

func TestFindByName(t *testing.T) {
	tx := getTx(t)
	defer tx.Rollback()

	repo := repository.NewUserRepository()
	// Pastikan ada data dummy
	dummy := domain.Student{Name: "FindByName Test", Email: "findbyname@example.com"}
	_, err := tx.Exec(`INSERT INTO students("Name", "Email") VALUES ($1, $2)`, dummy.Name, dummy.Email)
	if err != nil {
		t.Fatalf("❌ Gagal insert dummy: %v", err)
	}

	user, err := repo.FindByName(context.Background(), tx, dummy.Name)
	if err != nil {
		t.Fatalf("❌ FindByName error: %v", err)
	}
	if user.Name != dummy.Name {
		t.Errorf("❌ Nama tidak sesuai. Got: %s, Expected: %s", user.Name, dummy.Name)
	} else {
		t.Logf("✅ FindByName sukses: %+v", user)
	}
}

func TestFindAll(t *testing.T) {
	tx := getTx(t)
	defer tx.Rollback()

	repo := repository.NewUserRepository()

	users := repo.FindAll(context.Background(), tx)
	if len(users) == 0 {
		t.Logf("⚠️ Tidak ada user di tabel students")
	} else {
		t.Logf("✅ FindAll sukses, jumlah user: %d", len(users))
	}
}

func TestUserExist(t *testing.T) {
	tx := getTx(t)
	defer tx.Rollback()

	repo := repository.NewUserRepository()
	email := "exist@example.com"

	// Insert data dummy
	_, err := tx.Exec(`INSERT INTO students("Name", "Email") VALUES ($1, $2)`, "Exist Test", email)
	if err != nil {
		t.Fatalf("❌ Gagal insert dummy: %v", err)
	}

	exists, err := repo.UserExist(context.Background(), tx, email)
	if err != nil {
		t.Fatalf("❌ UserExist error: %v", err)
	}
	if !exists {
		t.Errorf("❌ Harusnya user %s ada, tapi UserExist return false", email)
	} else {
		t.Logf("✅ UserExist sukses, user %s ditemukan", email)
	}
}

package main

import (
	"a21hc3NpZ25tZW50/helper"
	"fmt"
	"strings"
)

var Students string = "A1234_Aditira_TI, B2131_Dito_TK, A3455_Afis_MI"
var StudentStudyPrograms string = "TI_Teknik Informatika, TK_Teknik Komputer, SI_Sistem Informasi, MI_Manajemen Informasi"

func Login(id string, name string) string {
	studentData := strings.Split(Students, ", ")
	for _, data := range studentData {
		info := strings.Split(data, "_")
		if info[0] == id {
			if info[1] == name {
				program := GetStudyProgram(strings.Split(info[2], "")[0])
				return fmt.Sprintf("Login berhasil : %s (%s)", name, program)
			} else {
				return "Login gagal: data mahasiswa tidak ditemukan"
			}
		}
	}
	var msg string
	if id == "" || name == "" {
		msg = "ID or Name is undefined!"
	} else if len(id) != 5 {
		msg = "ID must be 5 characters long!"
	}
	return msg // TODO: replace this
}

func Register(id string, name string, major string) string {
	if info[0] == id {
		return "Registrasi gagal: id sudah digunakan"
	} else if info[1] != id {
		return "Registrasi berhasil: <nama> (<jurusan mahasiswa>)"
	}
	var msg string
	if id == "" || name == "" || major == "" {
		msg = "ID, Name or Major is undefined!"
	} else if len(id) != 5 {
		msg = "ID must be 5 characters long!"
	}
	return msg // TODO: replace this
}

func GetStudyProgram(code string) string {
	var msg string
	if code == "" {
		msg = "Code is undefined!"
	}
	return msg // TODO: replace this
}

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		fmt.Println("Students: ", Students)
		fmt.Println("Student Study Programs: ", StudentStudyPrograms)

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Keluar")

		var pilihan string
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			helper.ClearScreen()
			var id, name string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)

			fmt.Println(Login(id, name))

			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			var id, name, jurusan string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)
			fmt.Print("Masukkan jurusan: ")
			fmt.Scan(&jurusan)
			fmt.Println(Register(id, name, jurusan))

			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			var kode string
			fmt.Print("Masukkan kode: ")
			fmt.Scan(&kode)

			fmt.Println(GetStudyProgram(kode))
			helper.Delay(5)
		case "4":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

package main

import "fmt"
const NMAX int = 10
type sea struct{
	peringkat, emas, perak, perunggu int
	negara string
}
type tabSea [NMAX] sea
func main(){
	var SeaGames tabSea
	var negara string
	var option, pilih, jumlahnegara, jumlah, rank, search int
	for option != 7{
		menu()
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&option)
		for option < 1 || option > 7{
			fmt.Print("Masukkan Salah Silahkan Coba Lagi: ")
			fmt.Scan(&option)
		}
		if option == 1{
			if jumlahnegara == 0{
				fmt.Println("Array Kosong")
			}else {
				cetakPeringkat(SeaGames, jumlahnegara)
			}
		}else if option == 2{
			fmt.Print("Jumlah Negara yang Ditambahkan: ")
			fmt.Scan(&jumlah)
			tambahNegara(&SeaGames, &jumlahnegara, jumlah)
		}else if option == 3{
			var negaraasal, negarabaru string
			if jumlahnegara == 0{
				fmt.Println("Array Kosong")
			}else {
				fmt.Print("Negara Yang Akan Diubah: ")
				fmt.Scan(&negaraasal)
				fmt.Print("Nama Negara Baru: ")
				fmt.Scan(&negarabaru)
				ubahNegara(&SeaGames, &jumlahnegara, negaraasal, negarabaru)
			}
		}else if option == 4{
			if jumlahnegara == 0{
				fmt.Println("Array Kosong")
			}else {
				fmt.Print("Negara Yang Akan Dihapus: ")
				fmt.Scan(&negara)
				hapusNegara(&SeaGames, &jumlahnegara, negara)
			}
		}else if option == 5{
			if jumlahnegara == 0{
				fmt.Println("Array Kosong")
			}else{
				fmt.Println("--------MENU--------")
				fmt.Println("1. Ubah Medali")
				fmt.Println("2. Tambah Medali")
				fmt.Println("3. Hapus Medali")
				fmt.Println("--------------------")
				fmt.Print("Pilih: ")
				fmt.Scan(&pilih)
				for pilih < 1 || pilih > 3{
					fmt.Print("Masukkan Salah Silahkan Coba Lagi: ")
					fmt.Scan(&pilih)
				}
				if pilih == 1{
					fmt.Print("Ubah Medali Negara: ")
					fmt.Scan(&negara)
					ubahMedali(&SeaGames, &jumlahnegara, negara)
				}else if pilih == 2{
					fmt.Print("Tambah Medali Negara: ")
					fmt.Scan(&negara)
					tambahMedali(&SeaGames, &jumlahnegara, negara)
				}else if pilih == 3{
					fmt.Print("Hapus Medali Negara: ")
					fmt.Scan(&negara)
					hapusMedali(&SeaGames, &jumlahnegara, negara)
				}
			}
		}else if option == 6{
			if jumlahnegara == 0{
				fmt.Println("Array Kosong")
			}else{
				fmt.Println("1. Cari Berdasarkan Peringkat ")
				fmt.Println("2. Cari Berdasarkan Negara ")
				fmt.Print("Option: ")
				fmt.Scan(&search)
				for search != 1 && search != 2{
					fmt.Print("Masukkan Salah Silahkan Coba Lagi: ")
					fmt.Scan(&search)
				}
				if search == 1{
					fmt.Print("Masukkan Peringkat yang Dicari: ")
					fmt.Scan(&rank)
					cariPeringkat(SeaGames, jumlahnegara, rank)
				}else if search == 2{
					fmt.Print("Masukkan Negara yang Dicari: ")
					fmt.Scan(&negara)
					cariNegara(SeaGames, jumlahnegara, negara)
				}
			}
		}
	}
	fmt.Println("Terima Kasih")
}

func menu(){
	fmt.Println("--------MENU--------")
	fmt.Println("1. Tampilkan Peringkat")
	fmt.Println("2. Tambah Negara")
	fmt.Println("3. Ubah Nama Negara")
	fmt.Println("4. Hapus Negara")
	fmt.Println("5. Ubah Medali")
	fmt.Println("6. Cari Negara")
	fmt.Println("7. Exit")
	fmt.Println("--------------------")
}

func sequentialsearch(A tabSea, n int, negara string) int{
	var i, index int
	insertionsort(&A,n)
	index = -1
	i = 0
	for i < n && index == -1{
		if A[i].negara == negara{
			index = i
		}
		i++
	}
	return index
}

func tambahNegara(A *tabSea, n *int, jumlah int){
	var i, index int
	index = 1
	i = *n
	for i < NMAX && index <= jumlah{
		fmt.Print("Masukkan Nama Negara: ")
		fmt.Scan(&A[i].negara)
		A[i].emas = 0
		A[i].perak = 0
		A[i].perunggu = 0
		A[i].peringkat = i + 1
		i++
		index++
	}
	*n = i
}

func cetakPeringkat(A tabSea, n int){
	selectionsort(&A,n)
	fmt.Printf("%-15s %-10s %-10s %-10s %s","Peringkat", "Negara", "Emas", "Perak", "Perunggu")
	fmt.Println()
	for i := 0; i < n; i++{
		fmt.Printf("%-15d %-10s %-10d %-10d %d", A[i].peringkat, A[i].negara, A[i].emas, A[i].perak, A[i].perunggu)
		fmt.Println()
	}
}

func hapusNegara(A *tabSea, n *int, negara string){
	var index, i int
	insertionsort(&*A,*n)
	index = sequentialsearch(*A, *n, negara)
	if index != -1{
		i = index
		for i < *n-1{
			A[i].negara = A[i+1].negara
			A[i].emas = A[i+1].emas
			A[i].perak = A[i+1].perak
			A[i].perunggu = A[i+1].perunggu
			i++
		}
		*n -=1
	}else{
		fmt.Println("Negara Tidak Ada Dalam Peringkat")
	}
}

func ubahNegara(A *tabSea, n *int, negaraasal, negarabaru string){
	var index, i int
	insertionsort(&*A,*n)
	index = sequentialsearch(*A, *n, negaraasal)
	if index != -1{
		i = index
		A[i].negara = negarabaru
	}else{
		fmt.Println("Negara Tidak Ada Dalam Peringkat")
	}
}

func ubahMedali(A *tabSea, n *int, negara string){
	var index, i, emas, perak, perunggu int
	insertionsort(&*A,*n)
	index = sequentialsearch(*A, *n, negara)
	if index != -1{
		i = index
		fmt.Print("Perolehan emas: ")
		fmt.Scan(&emas)
		fmt.Print("Perolehan perak: ")
		fmt.Scan(&perak)
		fmt.Print("Perolehan perunggu: ")
		fmt.Scan(&perunggu)
		A[i].emas = emas
		A[i].perak = perak
		A[i].perunggu = perunggu
	}else{
		fmt.Println("Negara Tidak Ada Dalam Peringkat")
	}
}

func insertionsort(A *tabSea, n int) {
    var pass, i int
	var temp sea
	pass = 1
	for pass < n{
		i = pass
		temp = A[pass]
		for (i > 0 && temp.emas > A[i-1].emas) || (i > 0 && temp.emas == A[i-1].emas && temp.perak > A[i-1].perak) || (i > 0 && temp.emas == A[i-1].emas && temp.perak == A[i-1].perak && temp.perunggu > A[i-1].perunggu){
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
	
	i = 0
	for i < n{
		A[i].peringkat = i + 1
		i++
	}
}

func selectionsort(A *tabSea, n int) {
    var pass, i, idx int
	var temp sea
	pass = 1
	for pass < n{
		idx = pass - 1
		i = pass
		for i < n{
			if A[idx].emas < A[i].emas{
				idx = i
			}else if A[idx].emas == A[i].emas && A[idx].perak < A[i].perak{
				idx = i
			}else if (A[idx].emas == A[i].emas && A[idx].perak == A[i].perak) && A[idx].perunggu < A[i].perunggu{
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
	
	i = 0
	for i < n{
		A[i].peringkat = i + 1
		i++
	}
}

func binarySearch(A tabSea, n, peringkat int) int{
	var mid, left, right, found int
	found = -1
	left = 0
	right = n
	mid = (left + right)/2
	for left <= right && found == -1{
		if peringkat < A[mid].peringkat{
			right = mid - 1
		}else if peringkat > A[mid].peringkat{
			left = mid + 1
		}else{
			found = mid
		}
		mid = (left + right)/2
	}
	return found
}

func cariPeringkat(A tabSea, n, peringkat int){
	var index int
	insertionsort(&A,n)
	index = binarySearch(A, n, peringkat)
	if index != -1{
		fmt.Printf("%-15s %-10s %-10s %-10s %s","Peringkat", "Negara", "Emas", "Perak", "Perunggu")
		fmt.Println()
		fmt.Printf("%-15d %-10s %-10d %-10d %d", A[index].peringkat, A[index].negara, A[index].emas, A[index].perak, A[index].perunggu)
		fmt.Println()
	}else{
		fmt.Println("Negara Tidak Ada Dalam Peringkat")
	}
	
}

func tambahMedali(A *tabSea, n *int, negara string){
	var index, i, emas, perak, perunggu int
	insertionsort(&*A,*n)
	index = sequentialsearch(*A, *n, negara)
	if index != -1{
		i = index
		fmt.Print("Perolehan emas: ")
		fmt.Scan(&emas)
		fmt.Print("Perolehan perak: ")
		fmt.Scan(&perak)
		fmt.Print("Perolehan perunggu: ")
		fmt.Scan(&perunggu)
		A[i].emas = A[i].emas + emas
		A[i].perak = A[i].perak + perak
		A[i].perunggu = A[i].perunggu + perunggu
	}else{
		fmt.Println("Negara Tidak Ada Dalam Peringkat")
	}
}

func hapusMedali(A *tabSea, n *int, negara string){
	var index, i int
	insertionsort(&*A,*n)
	index = sequentialsearch(*A, *n, negara)
	if index != -1{
		i = index
		A[i].emas = 0
		A[i].perak = 0
		A[i].perunggu = 0
	}else{
		fmt.Println("Negara Tidak Ada Dalam Peringkat")
	}
}

func cariNegara(A tabSea, n int, negara string){
	var index int
	insertionsort(&A,n)
	index = sequentialsearch(A, n, negara)
	if index != -1{
		fmt.Printf("%-15s %-10s %-10s %-10s %s","Peringkat", "Negara", "Emas", "Perak", "Perunggu")
		fmt.Println()
		fmt.Printf("%-15d %-10s %-10d %-10d %d", A[index].peringkat, A[index].negara, A[index].emas, A[index].perak, A[index].perunggu)
		fmt.Println()
	}else{
		fmt.Println("Negara Tidak Ada Dalam Peringkat")
	}
}
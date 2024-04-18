#include <iostream>
#include <string>
using namespace std;

// Deklarasi struktur data. var = nilai MK
typedef struct
{
    string npm;
    string nama_mhs;
    double nilai;
} nilaiMK;

// Deklarasi queue pakai array. var = queue
typedef struct
{
    int first;
    int last;
    nilaiMK dat[10];
} queue;

// Deklarasi fungsi buat queue kosong. var = *q
void buatqkosong(queue *q)
{
    q->first = -1;
    q->last = -1;
}

// mengembalikan nilai TRUE jika queue kosong. var = q di atas)
bool isikosong(queue q)
{
    bool hasil = false;

    if (q.first == -1 && q.last == -1)
    {
        hasil = true;
    }
    return hasil;
}

bool isipenuh(queue q)
{
    bool hasil = false;
    
    if (q.last == 9)
    {
        hasil = true;
    }
    return hasil;
}

// Deklarasi fungsi memasukkan elemen ke dalam queue. var = npm, nama_mhs, nilai, *q
void ADD(string npm, string nama_mhs, double nilai, queue *q)
{
    if (isikosong(*q))
    {
        // kalo queue kosong
        q->first = 0;
        q->last = 0;
        q->dat[0].npm = npm;
        q->dat[0].nama_mhs = nama_mhs;
        q->dat[0].nilai = nilai;
    }
    else if (!isipenuh(*q))
    {
        //kalo queue sudah terisi
        q->last = q->last + 1;
        q->dat[q->last].npm = npm;
        q->dat[q->last].nama_mhs = nama_mhs;
        q->dat[q->last].nilai = nilai;
    }
    else
    {
        // kalo queue sudah penuh
        cout << "QUEUE SUDAH PENUH" << endl;
    }
}

// Deklarasi fungsi untuk mengeluarkan elemen keluar queue. var = q
void DEL(queue *q)
{
    if (q->last == 0)
    {
        //kalo queue kosong
        q->first = -1;
        q->last = -1;
    }
    else
    {
        //kalo queue banyak, otomatis menggeser elemen ke depan
        for (int i = q->first + 1; i <= q->last; i++)
        {
            q->dat[i - 1].npm = q->dat[i].npm;
            q->dat[i - 1].nama_mhs = q->dat[i].nama_mhs;
            q->dat[i - 1].nilai = q->dat[i].nilai;
        }
        q->last = q->last - 1;
    }
}

// Deklarasi fungsi menampilkan isi queue. var = q
void cetakqueue(queue q)
{
    if (q.first != -1)
    {
        //kalo queue terbukti ada isinya 
        cout << " MENAMPILKAN QUEUE " << endl;
        cout << "~~~~~~~ " << endl;

        for (int i = q.last; i >= q.first; i--)
        {
            cout << " ~~~~~~~~~~" << endl;
            cout << " Elemen ke : " << i << endl;
            cout << " NPM MAHASISWA : " << q.dat[i].npm << endl;
            cout << " NAMA MAHASISWA : " << q.dat[i].nama_mhs << endl;
            cout << " NILAI MAHASISWA : " << q.dat[i].nilai << endl;
            cout << " ~~~~~~~~~~" << endl;
        }
    }
    else if (q.first == -1)
    {   // kalo queue tidak isinya
        cout << " QUEUE KOSONG " << endl;
    }
}

int main()
{
    queue q;
    buatqkosong(&q);
    cetakqueue(q);

    cout << endl;
    cout << endl;
    cout << " ~~~~~~ " << endl;
    
    ADD("4523210666", "Budi", 88.75, &q);
    ADD("4523210667", "Susi", 78.85, &q);
    ADD("4523210668", "Nuri", 98.65, &q);
    ADD("4523210669", "Bimo", 78.85, &q);
    ADD("4523210670", "Arif", 68.55, &q);
    ADD("4523210671", "Rido", 98.65, &q);
    ADD("4523210672", "Ella", 68.55, &q);
    cetakqueue(q);
    
    cout << " ~~~~~~ " << endl;
    cout << endl;
    cout << endl;
    
    DEL(&q);
    cetakqueue(q);
    cout << endl;
    cout << endl;
    
    DEL(&q);
    cetakqueue(q);
    
    cout << " ~~~~~~ " << endl;
    
    return 0;
}

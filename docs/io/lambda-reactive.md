gaya pemograman reactive, ketika ada event barulah suatu line di eksekusi.
manfaat lambda:
1. membuat scope baru pada function tertentu, sehingga user dari function dengan tambahan lambda bebas membuat algoritma baru pada function dengan lambda tersebut.
extend function with limited scope.

misalkan pada ruby kita mengenal:
10.times do
    # scope_tertentu
end

times menambahkan block tertentu


2. bisa digunakan pada case reactive programming/handler, suatu fungsi akan dipanggil ketika suatu event ditrigger.


callFuncA("eventAB", func(){

})

anonymous func() baru akan dipanggil ketika "eventAB" dicall.

pada javascript, ini disebut event-loop


// console.log("hello student");
// alert("good morning");

// document.write("batch 47");

// var gelas = "air putih";
// var gelas = "kopi";
// let adalah variabel yang tidak bisa di inisiasikan lagi misal let ganti, maka tidak bisa lagi di inisiasikan di bawah dengan penulisan let ganti, tetapi
//nilai yang ada di dalamnya masih bisa diganti
// const adalah variabel yang nilainya tidak akan bisa diganti
// console.log(gelas);

// console.log(`minuman saya ${gelas}`);
// console.log("minuman saya",gelas);

// function hello2(a,b){
//     console.log(a);
//     console.log(b);

//     console.log("hasil penambahan : "+(a+b));
// }

// hello2(1,2);

function submitData(){
    let name = document.getElementById("name").value;
    let email = document.getElementById("email").value;
    let phone = document.getElementById("number").value;
    let message = document.getElementById("message").value;
    let subject =document.getElementById("subject").value;

    if(name==""){
        return alert("isi nama kalian");
    }

    else if (email==""){
        return alert("email harus di isi");
    }

    else if (phone==""){
        return alert("Nomor HP harus diisi");
    }

    else if(subject==""){
        return alert("Subject harus diisi");
    }

    else if(message==""){
        return alert("Message harus diisi");
    }


    let emailReceiver="diza.febriyan@gmail.com";
    let a = document.createElement("a");
    a.href = `mailto:${emailReceiver}?subject=${subject}&body=Halo, nama saya ${name} , ${message}. Silahkan kontak saya di nomor ${phone}, terimakasih`; 
    a.click();


    let emailer={
        name,
        email,
        phone,
        subject,
        message,
    };

    console.log(emailer);
}
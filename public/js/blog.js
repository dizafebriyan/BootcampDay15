
// let datakita = {
//     nama:"diza febriyan",
//     umur:23,
//     asal:"palembang",
// };

// console.log(datakita);

// console.log(datakita.nama);

// //array of object

// let kumpulannama=[
//     {nama: "diza febriyan",umur: 23,},
//     {nama: "ucok", umur:22},
// ];

// console.log(kumpulannama);
// console.log(kumpulannama[1].nama);

let dataproject=[];

function addblog(event){  
event.preventDefault();

let title = document.getElementById("inputTitle").value;
let startdate = document.getElementById("input-blog-date-start").value;
let enddate = document.getElementById("input-blog-date-end").value;
let content = document.getElementById("inputContent").value;
let data1 = Number(document.getElementById("inputcheckbox_iot").checked);
let data2 = Number(document.getElementById("inputcheckbox_ui").checked);
let data3 = Number(document.getElementById("inputcheckbox_full").checked);
let data4 = Number(document.getElementById("inputcheckbox_ml").checked);
let image = document.getElementById("inputImage").files;
let posttime=new Date();
image=URL.createObjectURL(image[0]);
start = new Date(startdate);
end = new Date(enddate);

if(title==""){
    return alert("isi Title!");
}

else if (startdate==""){
    return alert("Tanggal start project harus diisi!");
}

else if (enddate==""){
    return alert("Tanggal end project harus diisi!");
}

else if(content==""){
    return alert("Description harus diisi!");
}

else if(image==""){
    return alert("Image harus diisi!");
}

else if (end-start<0){
    return alert("waktu pengerjaan tidak bisa kembali ke masa lalu");
}

let datacheckbox = [data1,data2,data3,data4];

let dataform = {
    title,
    startdate,
    enddate,
    content,
    datacheckbox,
    image,
    posttime,
    
}

dataproject.push(dataform);

console.log(dataproject);
renderBlog();
}


function renderBlog(){
    document.getElementById("contents").innerHTML="";

    for(let index = 0;index<dataproject.length; index++){
        document.getElementById("contents").innerHTML+=`
        <div class="blog-list-item">
            <div class="blog-list-item-image">
                <img src="${dataproject[index].image}" alt="Foto Bahasan"/>
            </div>
            <div class="blog-list-item-content">
                <div class="blog-list-item-content-button">
                    <button class="btn-edit">Edit Project</button>
                    <button class="btn-delete">Delete Project</button>
                </div>

                <h1>
                    <a href="myproject/1" target="blank">
                        ${dataproject[index].title}
                    </a>
                </h1>
                <div class="blog-list-item-content-detail">${getTimeProject(dataproject[index].startdate,dataproject[index].enddate)}</div>
                <div class="blog-list-item-content-detail">Project Time : ${dataproject[index].startdate} - ${dataproject[index].enddate} </div>

                <p>
                    ${dataproject[index].content}
                </p>

            </div>

        </div>`
    }
}

// function gettimeago(time){
//    const databulan=[
//         "Jan",
//         "Feb",
//         "Mar",
//         "Apr",
//         "May",
//         "Jun",
//         "Jul",
//         "Aug",
//         "Sep",
//         "Oct",
//         "Nov",
//         "Dec",
//    ];
//    let monthindex=time.getMonth();

//    let date=time.getDate();
//    let month=databulan[monthindex];
//    let year = time.getFullYear();
//    let hours=date.getHours();
//    let minutes=date.getMinutes();

//    if (minutes<10){
//     minutes=`0${minutes}`;
//    }
// }

// function gettimedistance(time){
//     timenow=new Date();
//     let distance = timenow-time; //hasilnya milidetik

//     const milisecond=1000;
//     const second1hours=3600;
//     const hours1days=24;

//     let distanceday = Math.floor(
//         distance/(milisecond*second1hours*hours1days)
//         );
//     let distancehours= Math.floor(
//         distance/(milisecond*60*60)
//     );
//     let distanceminutes=Math.floor(
//         distance/(milisecond*60)
//     );
//     let distanceseconds=Math.floor(
//         distance/milisecond
//     );

//     if (distanceday>0){
//         return `${distanceday} Day Ago`;
//     }

//     else if (distancehours>0){
//         return `${distancehours} Hours Ago`;
//     }
    
//     else if (distanceminutes>0){
//         return `${distanceminutes} Minutes Ago`;
//     }

//     else{
//         return `${distanceseconds} Seconds Ago`;
//     }
    
// }

function getTimeProject(start,end){
    start=new Date(start);
    end = new Date(end);
    var durasi=end-start;
 
    const milisecond=1000;
    const second1hours=3600;
    const hours1days=24;
    const day1month=30;

    let distancemonth = Math.floor(
        durasi/(milisecond*second1hours*hours1days*day1month)
    );
    let distanceday = Math.floor(
        durasi/(milisecond*second1hours*hours1days)
        );
    let distancehours= Math.floor(
        durasi/(milisecond*60*60)
    );
    let distanceminutes=Math.floor(
        durasi/(milisecond*60)
    );
    let distanceseconds=Math.floor(
        durasi/milisecond
    );


    if (distancemonth>0){
        return `${distancemonth} Month`;
    }
    else if (distanceday>0){
        return `${distanceday} Day`;
    }

    else if (distancehours>0){
        return `${distancehours} Hours`;
    }
    
    else if (distanceminutes>0){
        return `${distanceminutes} Minutes`;
    }

    else if (distanceseconds>0){
        return `${distanceseconds} Second`;
    }
 
 

}




setInterval (function(){
    renderBlog();
},3000);


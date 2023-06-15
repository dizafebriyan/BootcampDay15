const promise = new Promise ((resolve,reject)=>{

    const xhr = new XMLHttpRequest();

    xhr.open("GET","https://api.npoint.io/1d3ad056ed00da7f1724",true);
    // parameter 1 : method
    // parameter 2 : is the URL
    // parameter 3 : true or false, if tru mean that we using asynchronous web assembly, if false mean that we using synchronous way.

    console.log(xhr);

    // xhr.onload = function(){}; // to load and check the status of the request
    // xhr.onerror = function(){}; // loaded when it is error when we requesting the data
    // xhr.send(); // send the request to the server

    xhr.onload =()=>{
        if (xhr.status === 200){
            resolve (JSON.parse(xhr.response));
        }
        else {
            reject("error Loading Data.");
        }
    };

    xhr.onerror = () => {
        reject ("Network Error.");
    };

    xhr.send();
});

async function getAllTestimonials(){
    const response=await promise;
    
    let testimonialHTML ="";
    response.forEach(function (item){
            testimonialHTML += `<div class="testimonial">
                                    <img src="${item.image}" alt="Foto User">
                                    <p class="quote">${item.quote}</p>
                                    <p class="namequote">${item.author}</p>
                                    <p class="ratting">${item.rating} <i class="fa-solid fa-star"></i></p>
                                </div>`;
});
    document.getElementById('testimonials').innerHTML=testimonialHTML;
}

getAllTestimonials();

async function getFilteredTestimonials(rating){
    const response=await promise;
    let testimonialHTML ="";

    const testimonialFiltered = response.filter((item)=>{
        return item.rating === rating;
    });

        if (testimonialFiltered.length == 0) {
            testimonialHTML+=`<h1>"Data Not Found!"</h1>`;
        }

        else {
            testimonialFiltered.forEach(function (item){
                testimonialHTML += `<div class="testimonial">
                                    <img src="${item.image}" alt="Foto User">
                                    <p class="quote">${item.quote}</p>
                                    <p class="namequote">${item.author}</p>
                                    <p class="ratting">${item.rating} <i class="fa-solid fa-star"></i></p>
                                    </div>`;
            })
        };

    document.getElementById('testimonials').innerHTML=testimonialHTML;
}
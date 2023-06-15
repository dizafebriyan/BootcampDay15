// class Testimonials {
//     #quote="";
//     #image="";

//     constructor (quote,image){
//         this.#quote=quote;
//         this.#image=image;
//     }
//     get quote (){
//         return this.#quote;
//     }
//     get image(){
//         return this.#image;
//     }
//     get author(){
//         throw new Error ("getAuthor() Method must be impelemented");
//     }
//     get testimonialHTML(){
//         return `<div class="testimonial">
//                     <img src="${this.image}" alt="Foto User">
//                     <p class="quote">${this.quote}</p>
//                     <p class="namequote">${this.author}</p>
//                 </div>`
//     }

// }

// class AuthorTestimonials extends Testimonials{
//     #author="";

//     constructor(author,quote,image){
//         super(quote,image);
//         this.#author=author;
//     }
//     get author(){
//         return this.#author;
//     }
// }

// class CompanyTestimonials extends Testimonials {
//     #company="";

//     constructor (company,quote,image){
//         super(quote,image);
//         this.#company=company;
//     }
//     get author(){
//         return this.#company + " company";
//     }
// }

// const testimonial1 = new AuthorTestimonials(
//     "- Febri",
//     "Mantappp",
//     "https://xsgames.co/randomusers/avatar.php?g=male"
//     );
// const testimonial2 =new AuthorTestimonials(
//     "- Joko",
//     "Sangat memberikan jasa yang profesional",
//     "https://xsgames.co/randomusers/avatar.php?g=female"
// )
// const testimonial3= new CompanyTestimonials(
//     "- PT Semen Baturaja",
//     "Memberikan kualitas yang baik",
//     "https://xsgames.co/randomusers/avatar.php?g=pixel"
// )

// let testimonialData = [testimonial1,testimonial2,testimonial3];
// let testimonialHTML=``;

// for (let i=0;i<testimonialData.length;i++){
//     testimonialHTML += testimonialData[i].testimonialHTML; 
// }

// document.getElementById("testimonials").innerHTML = testimonialHTML;


// // MATERI HOF
// datatestimonial=[
//     {
//         author : "Ferdi",
//         komen : "memberikan pelayanan yang profesional",
//         image : "https://xsgames.co/randomusers/avatar.php?g=male",
//         ratting : 5
//     },
//     {
//         author : "Gisel",
//         komen : "sangat mantapppp",
//         image : "https://xsgames.co/randomusers/avatar.php?g=female",
//         ratting : 4
//     },
//     {
//         author : "Anonym",
//         komen : "gak bagus",
//         image : "https://xsgames.co/randomusers/avatar.php?g=pixel",
//         ratting : 2
//     }
// ];

// function allTestimonials(){
//      let testimonialHTML =``;
//      datatestimonial.forEach(function(item){
//         testimonialHTML += `<div class="testimonial">
//                                 <img src="${item.image}" alt="Foto User">
//                                 <p class="quote">${item.komen}</p>
//                                 <p class="namequote">${item.author}</p>
//                                 <p class="ratting">${item.ratting} <i class="fa-solid fa-star"></i></p>
//                             </div>`;
//      });
//      document.getElementById("testimonials").innerHTML = testimonialHTML;

// };

// allTestimonials();


// function filterTestimonial (ratting){
//     let testimonialHTML='';

//     const testimonialfilter = datatestimonial.filter(function(item){
//         return item.ratting == ratting;
//     }); 

//     if (testimonialfilter.length == 0) {
//         testimonialHTML+=`<h1>"Data Not Found!"</h1>`;
//     }
//     else {
//         testimonialfilter.forEach(function (item){
//                 testimonialHTML += `<div class="testimonial">
//                                         <img src="${item.image}" alt="Foto User">
//                                         <p class="quote">${item.komen}</p>
//                                         <p class="namequote">${item.author}</p>
//                                         <p class="ratting">${item.ratting} <i class="fa-solid fa-star"></i></p>
//                                   </div>`;
//         })

//     }
//     document.getElementById("testimonials").innerHTML = testimonialHTML;
// }


 
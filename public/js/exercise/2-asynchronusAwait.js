let condition = true;

let promise = new Promise((resolve,reject)=>{
    if (condition){
        resolve("promise is resolved")
    }
    else {
        reject("promise is rejected")
    }
});

// if we previously call the promise like this:
// promise
// .then((value)=>{
//     console.log(value);
// })
// .catch((reason)=>{
//     console.log(reason);
// })

//ASYNCHRONOUS WAIT

async function asyncFunction(){
    const response = await promise;
    console.log(response);

}

asyncFunction();
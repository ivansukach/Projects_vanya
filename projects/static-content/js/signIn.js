let SignIn=document.getElementById("SignIn");
let name;
let surname;
let accessToken;
let refreshToken;
let coins;
let photo


SignIn.onsubmit = async(e) => {
    e.preventDefault();
    let response = await fetch("http://localhost:8081/signIn", {method: 'POST', body: new FormData(SignIn)}); // завершается с заголовками ответа
    let result = await response.json();
    console.log(result);
    console.log(response);
    console.log(result.Authorization);
    let myHeaders = new Headers();
    // let bla =Response.redirect("http://localhost:8081/restricted/main")
    // await bla.then(alert("Ivan"))
    // console.log(bla.text());
    fetch("https://jsonplaceholder.typicode.com/users", {method: 'POST',  redirect: 'follow'}).then(response2 =>{

    }).catch(function (err) { console.info(err)})
//     myHeaders.append('Authorization', 'text/html')
//     // let response2 = await fetch("https://jsonplaceholder.typicode.com/users", {method: 'POST',  headers: {
//     //         Authorization: result.Authorization, RefreshToken: result.RefreshToken}})
//     let response2 = await fetch("http://localhost:8081/restricted/main", {method: 'GET', headers: {
//             'Authorization': result.Authorization, 'RefreshToken': result.RefreshToken}}) ;
//     console.log(response2);
//     let response3=Response.redirect("http://localhost:8081/signUp", 301)
//     console.log(response3);
//     await fetch("http://localhost:8081/restricted/main", {method: 'POST', headers: {
//             'Authorization': result.Authorization, 'RefreshToken': result.RefreshToken}}).then(response2 =>{ }) ;
//     let res4 = response3.redirect("http://localhost:8081/delete", 302)
//     fetch("http://localhost:8081/restricted/main", {method: 'POST', headers: {
//             'Authorization': result.Authorization, 'RefreshToken': result.RefreshToken}}).then(response2 =>{ }) ;
//     await Response.redirect("http://localhost:8081/restricted/main", 301)
//     Response.redirect("http://localhost:8081/restricted/main", 301)
// // , 'NextURL': "http://localhost:8081/restricted/main.html"
//     // console.log(response2.text());
//     // await response2.text().then(html => document.write(html));
//     // document.URL="http://localhost:8081/test.html"

    document.location.replace("http://localhost:8081/restricted/main.html");
    // e.preventDefault();
}
async function getRequest(){

    // читать тело ответа в формате JSON
    // return result
}
getRequest().then(alert);
let currentUser = localStorage.getItem("activeUser");
let activeUser = JSON.parse(currentUser);
let emptyUser = {
    name: '',
    surname: '',
    avatar: '',
    coins: 0,
    accessToken: '',
    refreshToken:''
};
function saveToJSON(){
    let currentUser = JSON.stringify(activeUser);
    localStorage.setItem("activeUser", currentUser);
}
document.onload = () => {

}

// function setCurrentUserFields(){
//
// }
// let formSignIn = document.getElementById("SignIn");
// formSignIn.addEventListener("submit", function(event)  {
//     event.preventDefault(); setCurrentUserFields(); saveToJSON();
// });
// async function getRequest(){
//     let response = await fetch('http://localhost:8081/main'); // завершается с заголовками ответа
//     let result = await response.json();
//     console.log(result);
//     // читать тело ответа в формате JSON
//     // return result
// }
// getRequest().then(alert)
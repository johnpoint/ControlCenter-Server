var c = document.cookie.split(';');
for (let index = 0; index < c.length; index++) {
    c[index] = c[index].split('=');
}
var userjwt = "";
for (let index = 0; index < c.length; index++) {
    if (c[index][0].replace(' ', '') == "jwt") {
        userjwt = c[index][1].replace(' ', '');
    }
};

function checkuser() {

    if (userjwt == "") {
        if (window.location.pathname != "/login.html") {
            window.location.pathname = "/login.html"
        }
    }

    new Vue().$http.get(apiaddress + "/web/UserInfo", {
        headers: {
            "Content-Type": "application/x-www-form-urlencoded",
            "Authorization": "Bearer " + userjwt
        }
    }).then(function (res) {
        if (res.body.level != null) {
            userNav.name = res.body.name;
            userNav.nologin = false;
            userNav.login = true;
        }
    })
}
checkuser()

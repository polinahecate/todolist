async function myFunction() {
    let response = await fetch("http://localhost:8080/lists");

    if (response.ok) { // если HTTP-статус в диапазоне 200-299
        // получаем тело ответа (см. про этот метод ниже)
        let json = await response.json();
        //alert(json);
        for (var i = 0; i < 2; i++) {
            console.log("цикл выполняется");
            var elem = document.createElement("li");
            elem.innerHTML = json[i].id + "<button>" + json[i].name + "</button>";
            document.getElementById("lists").appendChild(elem);
        }
    } else {
        alert("Ошибка HTTP: " + response.status);
    }

}
function search_topics() {
    let input = document.getElementById('searchbar').value
    input=input.toLowerCase();
    let x = document.getElementsByClassName('topics');

    for (i = 0; i < x.length; i++) {
        if (!x[i].innerHTML.toLowerCase().includes(input)) {
            x[i].style.display="none";
        }
        else {
            x[i].style.display="list-item";
        }
    }
}

document.querySelectorAll('.like').forEach(item => {
    item.addEventListener('click', event => {
        if (item.src === "http://localhost:63342/Helpy/image/coeur(1).png") {  // to change
            item.src = "../img/coeur.png"
            item.nextElementSibling.innerHTML = parseInt(item.nextElementSibling.innerHTML) - 1
            return
        }
        item.src = "../img/coeur(1).png"
        // item.nextElementSibling.nextElementSibling.src = "../img/etoile.png"
        item.nextElementSibling.innerHTML = parseInt(item.nextElementSibling.innerHTML) + 1
    })
})

document.querySelectorAll('.fav').forEach(item => {
    item.addEventListener('click', event => {
        if (item.src == "http://localhost:63342/Helpy/image/etoile(1).png") { // to change
            item.src = "../img/etoile.png"
            item.previousElementSibling.innerHTML = parseInt(item.previousElementSibling.innerHTML) + 1
            return
        }
        item.src = "../img/etoile(1).png"
            // item.previousElementSibling.previousElementSibling.src = "../images/coeur.png"
        item.previousElementSibling.innerHTML = parseInt(item.previousElementSibling.innerHTML) - 1
    })
})
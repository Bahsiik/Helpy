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
        if (item.src === "http://localhost:63342/Helpy/Help/img/coeur(1).png") {  // to change
            item.src = "http://localhost:63342/Helpy/Help/img/coeur.png"
        } else {
            item.src = "http://localhost:63342/Helpy/Help/img/coeur(1).png"
        }
    })
})

document.querySelectorAll('.fav').forEach(item => {
    item.addEventListener('click', event => {
        if (item.src == "http://localhost:63342/Helpy/Help/img/etoile(1).png") { // to change
            item.src = "http://localhost:63342/Helpy/Help/img/etoile.png"
        } else {
            item.src = "http://localhost:63342/Helpy/Help/img/etoile(1).png"
        }
    })
})

//function that add a picture in the textarea
function add_picture(event) {
    let input = document.getElementById('boutonimage')
    let textarea = document.getElementById('description')
    textarea.value += '![alt text](/' + input.value + ')'
    input.value = localStorage.getItem('image')
}

input.addEventListener(onclick(), function () {
    localStorage.setItem('image', input.value)
})




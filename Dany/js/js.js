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
        if (item.src === "http://localhost:63342/Helpy/Dany/image/coeur(1).png") {  // to change
            item.src = "../image/coeur.png"
            item.nextElementSibling.innerHTML = parseInt(item.nextElementSibling.innerHTML) - 1
            return
        }
        item.src = "../image/coeur(1).png"
        // item.nextElementSibling.nextElementSibling.src = "../image/etoile.png"
        item.nextElementSibling.innerHTML = parseInt(item.nextElementSibling.innerHTML) + 1
    })
})

document.querySelectorAll('.fav').forEach(item => {
    item.addEventListener('click', event => {
        if (item.src == "http://localhost:63342/Helpy/Dany/image/etoile(1).png") { // to change
            item.src = "../image/etoile.png"
            item.previousElementSibling.innerHTML = parseInt(item.previousElementSibling.innerHTML) + 1
            return
        }
        item.src = "../image/etoile(1).png"
            // item.previousElementSibling.previousElementSibling.src = "../images/coeur.png"
        item.previousElementSibling.innerHTML = parseInt(item.previousElementSibling.innerHTML) - 1
    })
})

function mail() {
    document.getElementById('e-mail').innerHTML = document.getElementById('mail').value;
}

document.querySelectorAll('.modif_email').forEach(item => {
    item.addEventListener('click', event => {
        document.querySelector('.mail2').style.display = 'block';
    })
})
document.querySelectorAll('#close').forEach(item => {
    item.addEventListener('click', event => {
        document.querySelector('.mail2').style.display = 'none';
    })
})
document.querySelectorAll('.aria').forEach(item => {
    item.addEventListener('click', event => {
        if (item.ariaChecked == 'false') {
            document.querySelector('.aria').ariaChecked = 'true';
            document.querySelector('.round').style.marginLeft = '20px';
            document.querySelector('.aria').style.backgroundColor = 'black'
        } else {
            document.querySelector('.aria').ariaChecked = 'false';
            document.querySelector('.aria').style.backgroundColor = 'white'
            document.querySelector('.round').style.marginLeft = '';
        }
    })
})
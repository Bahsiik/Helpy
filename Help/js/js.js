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
            item.src = "http://localhost:63342/Helpy/Dany/image/coeur.png"
        } else {
            item.src = "http://localhost:63342/Helpy/Dany/image/coeur(1).png"
        }
    })
})

document.querySelectorAll('.fav').forEach(item => {
    item.addEventListener('click', event => {
        if (item.src == "http://localhost:63342/Helpy/Dany/image/etoile(1).png") { // to change
            item.src = "http://localhost:63342/Helpy/Dany/image/etoile.png"
            return
        }
        item.src = "http://localhost:63342/Helpy/Dany/image/etoile(1).png"
    })
})
function password(){
    document.querySelector('.message_group').style.display = 'block';
}
function leave_password() {
    document.querySelector('.message_group').style.display = 'none';
}
function mail() {
    document.getElementById('e-mail').innerHTML = document.getElementById('mail').value;
    document.querySelector('#change').innerHTML = 'Changer';
    document.querySelector('.mail2').style.display = 'none';
    document.querySelector('#change').style.color = 'black';
}
function pseudo() {
    document.getElementById('pseudo').innerHTML = document.getElementById('p_seudo').value;
    document.querySelector('#change').innerHTML = 'Changer';
    document.querySelector('.pseudo2').style.display = 'none';
    document.querySelector('#change').style.color = 'black';
}
function avatar() {
    document.getElementById('avatar').innerHTML = document.getElementById('p_seudo').value;
    document.querySelector('#change_avatar').innerHTML = 'Changer';
    document.querySelector('.container').style.display = 'none';
    document.querySelector('#change_avatar').style.color = 'black';
}
document.querySelectorAll('.modif_email').forEach(item => {
    item.addEventListener('click', event => {
        if (document.querySelector('#change').innerHTML === 'Changer'){
            document.querySelector('#change').innerHTML = 'Annuler';
            document.querySelector('#change').style.color = 'red';
            document.querySelector('.mail2').style.display = 'block';
        } else {
            document.querySelector('#change').innerHTML = 'Changer';
            document.querySelector('.mail2').style.display = 'none';
            document.querySelector('#change').style.color = 'black';
        }
    })
})

document.querySelectorAll('.modif_pseudo').forEach(item => {
    item.addEventListener('click', event => {
        if (document.querySelector('#change').innerHTML === 'Changer'){
            document.querySelector('#change').innerHTML = 'Annuler';
            document.querySelector('#change').style.color = 'red';
            document.querySelector('.pseudo2').style.display = 'block';
        } else {
            document.querySelector('#change').innerHTML = 'Changer';
            document.querySelector('.pseudo2').style.display = 'none';
            document.querySelector('#change').style.color = 'black';
        }
    })
})
document.querySelectorAll('.modif_avatar').forEach(item => {
    item.addEventListener('click', event => {
        if (document.querySelector('#change_avatar').innerHTML === 'Changer'){
            document.querySelector('#change_avatar').innerHTML = 'Annuler';
            document.querySelector('#change_avatar').style.color = 'red';
            document.querySelector('.container').style.display = 'block';
        } else {
            document.querySelector('#change_avatar').innerHTML = 'Changer';
            document.querySelector('.container').style.display = 'none';
            document.querySelector('#change_avatar').style.color = 'black';
        }
    })
})
document.querySelectorAll('.aria').forEach(item => {
    item.addEventListener('click', event => {
        if (item.ariaChecked === 'false') {
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
document.querySelectorAll('.aria1').forEach(item => {
    item.addEventListener('click', event => {
        if (item.ariaChecked === 'false') {
            document.querySelector('.aria1').ariaChecked = 'true';
            document.querySelector('.round1').style.marginLeft = '20px';
            document.querySelector('.aria1').style.backgroundColor = 'black'
        } else {
            document.querySelector('.aria1').ariaChecked = 'false';
            document.querySelector('.aria1').style.backgroundColor = 'white'
            document.querySelector('.round1').style.marginLeft = '';
        }
    })
})
document.querySelectorAll('.aria2').forEach(item => {
    item.addEventListener('click', event => {
        if (item.ariaChecked === 'false') {
            document.querySelector('.aria2').ariaChecked = 'true';
            document.querySelector('.round2').style.marginLeft = '20px';
            document.querySelector('.aria2').style.backgroundColor = 'black'
        } else {
            document.querySelector('.aria2').ariaChecked = 'false';
            document.querySelector('.aria2').style.backgroundColor = 'white'
            document.querySelector('.round2').style.marginLeft = '';
        }
    })
})
document.querySelectorAll('.aria3').forEach(item => {
    item.addEventListener('click', event => {
        if (item.ariaChecked === 'false') {
            document.querySelector('.aria3').ariaChecked = 'true';
            document.querySelector('.round3').style.marginLeft = '20px';
            document.querySelector('.aria3').style.backgroundColor = 'black'
        } else {
            document.querySelector('.aria3').ariaChecked = 'false';
            document.querySelector('.aria3').style.backgroundColor = 'white'
            document.querySelector('.round3').style.marginLeft = '';
        }
    })
})

document.querySelectorAll('.avatar_pic').forEach(item => {
    item.addEventListener('click', event => {
        if (item.id==='blue'){
            document.querySelector('#avatar').src = '../../image/profil.jpg'
        } else if (item.id === 'black') {
            document.querySelector('#avatar').src = '../../image/profil_black.jpg'

        } else if (item.id === 'green') {
            document.querySelector('#avatar').src = '../../image/profil_green.jpg'

        } else if (item.id === 'pink') {
            document.querySelector('#avatar').src = '../../image/profil_pink.jpg'

        } else if (item.id === 'red') {
            document.querySelector('#avatar').src = '../../image/profil_red.jpg'

        } else if (item.id === 'yellow') {
            document.querySelector('#avatar').src = '../../image/profil_yellow.jpg'

        }

    })
})
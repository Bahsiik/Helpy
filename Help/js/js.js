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


function nam(){
    document.querySelector('#message_name').style.display = 'block';
}
function leave_nam(){
    document.querySelector('#message_name').style.display = 'none';
}
function password(){
    document.querySelector('#message_group').style.display = 'block';
}
function leave_password() {
    document.querySelector('#message_group').style.display = 'none';
}
function mdp() {
    document.getElementById('mdp').innerHTML = document.getElementById('mail').value;
    document.querySelector('#change_mdp').innerHTML = 'Changer';
    document.querySelector('.Mdp2').style.display = 'none';
    document.querySelector('#change_mdp').style.color = 'black';
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
document.querySelectorAll('.modif_mdp').forEach(item => {
    item.addEventListener('click', event => {
        if (document.querySelector('#change_mdp').innerHTML === 'Changer'){
            document.querySelector('#change_mdp').innerHTML = 'Annuler';
            document.querySelector('#change_mdp').style.color = 'red';
            document.querySelector('.Mdp2').style.display = 'flex';
        } else {
            document.querySelector('#change_mdp').innerHTML = 'Changer';
            document.querySelector('.Mdp2').style.display = 'none';
            document.querySelector('#change_mdp').style.color = 'black';
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
            document.querySelector('.mail').style.marginTop = '10%';
        } else {
            document.querySelector('#change_avatar').innerHTML = 'Changer';
            document.querySelector('.container').style.display = 'none';
            document.querySelector('#change_avatar').style.color = 'black';
            document.querySelector('.mail').style.marginTop = '';
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
            document.querySelector('#avatar').src = '../img/profil.jpg'
        } else if (item.id === 'black') {
            document.querySelector('#avatar').src = '../img/profil_black.jpg'

        } else if (item.id === 'green') {
            document.querySelector('#avatar').src = '../img/profil_green.jpg'

        } else if (item.id === 'pink') {
            document.querySelector('#avatar').src = '../img/profil_pink.jpg'

        } else if (item.id === 'red') {
            document.querySelector('#avatar').src = '../img/profil_red.jpg'

        } else if (item.id === 'yellow') {
            document.querySelector('#avatar').src = '../img/profil_yellow.jpg'

        }

    })
})

document.querySelectorAll('.choix_profil span').forEach(item =>{
    item.addEventListener('click', event => {
        let item = event.target;
        let target = item.dataset.target;
        let spans = document.querySelectorAll('.choix_profil span');
        for (let i = 0; i < spans.length; i++) {
            spans[i].classList.remove('active');
        }
        item.classList.add('active');
        let divs = document.getElementsByClassName('contentDiv');
        for (let i = 0; i < divs.length; i++) {
            if(divs[i].classList.contains('profil'+target)) {
                divs[i].classList.remove('hidden');
            }
            else {
                divs[i].classList.add('hidden');
            }
        }
    })
})
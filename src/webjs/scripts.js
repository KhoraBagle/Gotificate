// let dogDiv = document.getElementById('doginfo')
// fetch('https://api.localhost:8080/dogs')
// .then(res => res.json())
// .then(dog => {
//     dogDiv.innetHTML += `<p> ${dog.dog}  </p>`

// })


const app = document.getElementById('root')

const logo = document.createElement('img')
logo.src = 'logo.png'

const container = document.createElement('div')
container.setAttribute('class', 'container')

app.appendChild(logo)
app.appendChild(container)

var request = new XMLHttpRequest()
request.open('GET', 'http://api.localhost:8080/dogs', true)

request.onload = function () {
  // Begin accessing JSON data here

  var data = JSON.parse(this.response)
  if (request.status >= 200 && request.status < 400) {
    data.forEach(dogs => {
      const card = document.createElement('div')
      card.setAttribute('class', 'card')

      const h1 = document.createElement('h1')
      h1.textContent = dogs.Name

      const h2 = document.createElement('h2')
      h2.textContent = dogs.Color

      const h3 = document.createElement('h3')
      h3.textContent = dogs.Size

      const d1 = document.createElement('d1')
      dogs.Disposition = dogs.Disposition.substring(0, 300)
      d1.textContent = `${dogs.Disposition}...`

      container.appendChild(card)
      card.appendChild(h1)
      card.appendChild(h2)
      card.appendChild(h3)
      card.appendChild(d1)
    })
  } else {
    const errorMessage = document.createElement('marquee')
    errorMessage.textContent = `Bark, it's not working!`
    app.appendChild(errorMessage)
  }
}

request.send()
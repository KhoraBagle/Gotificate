// javascript file to service index.html and interact with our API to populate the dog website

const app = document.getElementById('root')

// define picture of a cute dog
const logo = document.createElement('img')
logo.src = 'logo.png'

const container = document.createElement('div')
container.setAttribute('class', 'container')

app.appendChild(logo)
app.appendChild(container)

// sends GET request to get all available dogs API endpoint
var request = new XMLHttpRequest()
request.open('GET', 'http://api.localhost:8080/dogs', true)

// Browser will block CORS here and terminate script unless the functiopn is disabled per readme instructions
// This is only an issue because of local hosting and browser settings
request.onload = function () {
// Begin accessing JSON data here

  var data = JSON.parse(this.response)
// runs as long as no error code
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
  // error handling
  } else {
    const errorMessage = document.createElement('marquee')
    errorMessage.textContent = `Bark, it's not working!`
    app.appendChild(errorMessage)
  }
}

request.send()
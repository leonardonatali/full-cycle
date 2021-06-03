const express = require("express")
const { people } = require("./people")

const app = express()
const port = process.env.APP_PORT || 80

people.create()

app.get('/', (req, res) => {
    people.insert()

    people.get()
        .then((result) => {
            if (result) {
                let peopleList = ""

                for (const people of result) {
                    let id = people.id.toString()
                    let name = people.name
                    peopleList += `<strong>ID</strong>: ${id} <br> <strong>Name</strong>: ${name}<br><br>\n`
                }
                return res.send(`<h1>Full Cycle Rocks!</h1><hr>People list:<br>${peopleList}`)
            }
        }).catch((err) => {
            console.error('err', err);
            return res.send(`<h1>Full Cycle Rocks!</h1><hr><h4>An error has occurred</h4>`)
        })
})

app.listen(port, () => {
    console.log(`The app is listening at port ${port}`);
})

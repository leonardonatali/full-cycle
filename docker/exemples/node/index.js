const express = require("express")

const app = express()

const port = 80

app.get('/', (req, res) => {
    return res.send("<h3>Hello, Fullcycle</h3")
})

app.listen(port, () => {
    console.log(`The app is listening at port ${port}`);
})

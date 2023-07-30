import express from 'express'

const app = express()
const port = 8080

app.get('/', (req, res) => {
    res.send('Hello World!')
})

app.get('/health', (req, res) => {
    res.status(200).end()
})

app.listen(port, () => {
    console.log(`listening on port ${port}`)
})
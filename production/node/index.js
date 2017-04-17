let express = require('express')

let app = express();

app.get('*', (req,res)=>{

  res.send('Привет')


})


app.listen(5000,(err)=>{
  console.log(err)
})

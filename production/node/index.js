let express = require('express')

let app = express();

app.get('*', (req,res)=>{

  res.send('Привет')


})


app.listen(3000,(err)=>{
  console.log(err)
})

const http = require('http');
const Chance = require('chance');
const chance = new Chance();

const data = {
    "horses": [{
        "id": chance.postal()+chance.radio(),
        "name": chance.name(), 
        "lineID": 1,
    },{
        "id": chance.postal()+chance.radio(),
        "name": chance.name(), 
        "lineID": 2,
    },{
        "id": chance.postal()+chance.radio(),
        "name": chance.name(), 
        "lineID": 3,
    },{
        "id": chance.postal()+chance.radio(),
        "name": chance.name(), 
        "lineID": 4,
    }],
    "odds": [{
        "id": 1,
        "odds": chance.floating({min: 0, max: 47}),
    },{
        "id": 2,
        "odds": chance.floating({min: 0, max: 47}),
    },{
        "id": 3,
        "odds": chance.floating({min: 0, max: 47}),
    },{
        "id": 4,
        "odds": chance.floating({min: 0, max: 47}),
    }],
};

const server = http.createServer(function(request, response){
  response.writeHead(200, {'Content-Type':'application/json'});
  response.write(JSON.stringify(data));
  response.end();

})

console.log(`Server is listening on ${process.env.PORT}`)
server.listen(process.env.PORT);


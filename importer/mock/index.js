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
        "odd": chance.floating({min: 0, max: 47}),
    },{
        "id": 2,
        "odd": chance.floating({min: 0, max: 47}),
    },{
        "id": 3,
        "odd": chance.floating({min: 0, max: 47}),
    },{
        "id": 4,
        "odd": chance.floating({min: 0, max: 47}),
    }],
};

http.createServer(function(request, response){
  response.writeHead(200, {'Content-Type':'application/json'});
  response.write(JSON.stringify(data));
  response.end();

}).listen(7000);


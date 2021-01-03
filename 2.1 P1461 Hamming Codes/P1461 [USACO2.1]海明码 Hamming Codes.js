const fs = require('fs')
const data = fs.readFileSync('/dev/stdin', {encoding: 'ascii'})
// const data = fs.readFileSync('./testdata.in', {encoding: 'ascii'})
const tmp = data.trim().split(' ')
const [n, b, d] = [tmp[0], tmp[1], tmp[2]];
// console.log(n, b, d)

let res = [0]
for(let i=0; i< 2 ** b; i++) {
    let flag =1
    for(let j=0; j<res.length; j++){
        const length = (res[j] ^ i).toString(2).split(1).length - 1;
        if(length < d){
            flag =0;
            break
        }
    }
    flag && res.push(i)
    if(res.length >= n) {
        break;
    }
}

res = res.slice(0,n)
for(let i = 0; i<res.length/10; i++){
    console.log(res.slice(i*10, i*10+10).join(' '))
}

process.exit()
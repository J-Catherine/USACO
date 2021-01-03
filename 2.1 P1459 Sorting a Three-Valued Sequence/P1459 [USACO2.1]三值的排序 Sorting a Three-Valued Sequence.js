const fs = require('fs')
// const data = fs.readFileSync('/dev/stdin', {encoding: 'ascii'})
const data = fs.readFileSync('./testdata.in', {encoding: 'ascii'});
const list = data.trim().split('\n').map(x => parseInt(x));

const n = list[0];
let [a1, a2,a3] = [0,0,0]
list.splice(0,1)
list.map(item => {
    if(item === 1) {
        a1 += 1;
    } else if(item === 3) {
        a3+=1
    } else {
        a2+=1
    }
})
let [a12, a13, a21, a23,a31,a32] = [0,0,0,0,0,0];

list.map((item,idx) => {
    if(idx < a1) {
        item === 2 && a12++
        item === 3 && a13++
    } else if(idx >= a1 + a2) {
        item === 1 && a31++
        item === 2 && a32++
    } else {
        item === 1 && a21++
        item === 3 && a23++
    }
})
const result = Math.min(a31, a13) + Math.max(a21,a12) + Math.max(a23,a32)
console.log(result);
process.exit() // 请注意必须在出口点处加入此行

// [ 1, 1, 1, 1, 1, 1, 1, 1, 2, 1
//     3, 2, 2, 2,
//     3, 3, 3, 3, 3, 1 ]


const fs = require('fs')
const temp_data = fs.readFileSync('/dev/stdin')
// const temp_data = '1'
const data = parseInt(temp_data.toString('ascii').trim());
const  compareValue  = (props) => {
    return (a,b) => {
        let [va, vb] = [a[props], b[props]];
        if(va > vb) {
            return 1;
        } else if(va < vb) {
            return -1;
        } else {
            return 0;
        }
    }
}
// calculate all
let res = [];
for(let i =1; i<data+1; i++) {
    let cur = new Array(i);
    for(let j=0; j<i; j++) {
        cur[j] = {
            idx: `${j}/${i}`,
            value:  j/i,
        };
    }
    res = res.concat(cur);
}
res.sort(compareValue("value"))
const result = res.reduce((prev, curr, idx) => {
    if(idx === 0) {
        return [curr.idx];
    }
    if(res[idx-1].value === curr.value) {
        return prev;
    }
    return prev.concat(curr.idx);
}, [])

result.forEach((item)=>{
    console.log(item)
})
console.log('1/1')
process.exit() // 请注意必须在出口点处加入此行
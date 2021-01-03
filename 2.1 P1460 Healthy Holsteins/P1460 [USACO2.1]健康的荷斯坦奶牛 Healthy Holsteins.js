const fs = require('fs')
const data = fs.readFileSync('/dev/stdin', {encoding: 'ascii'})
// const data = fs.readFileSync('./testdata.in', {encoding: 'ascii'})
const result = data.trim().split('\n')
const v = parseInt(result[0]);
const vList = result[1].trim().split(' ').map(item => parseInt(item));
const g = parseInt(result[2]);
const gList = result.slice(3).map(item => item.split(' ').map(item => parseInt(item)));
// console.log(v, vList, g,gList)

const compareNormal = (v) => {
    const normal = vList;
    if(v.length !== normal.length) {
        return false;
    }
    for(let i=0; i<normal.length; i++) {
        if(v[i] < normal[i]) {
            return false;
        }
    }
    return true;
}

const addition = (a, b) => {
    // console.log(a,b)
    if(a.length !== b.length) { return a }
    return a.map((item,idx) => item+b[idx])
}

let res = {
    type: v,
    typeList: []
}
const dfs = (idx, currSum, currType, typeList) => {
    // console.log(idx, currSum, currType, typeList)
    if(compareNormal(currSum)) {
        if(currType <= res.type) {
            res = {
                type: currType,
                typeList: typeList.slice()
            }
            return 
        };
    }
    if(idx >= g) { return }
    typeList[idx] = 1
    // console.log(idx, gList[idx])
    dfs(idx+1, addition(currSum, gList[idx]), currType+1, typeList);
    typeList[idx] = 0
    dfs(idx+1, currSum, currType, typeList);
}

const a = new Array(v);
const b = new Array(v);
for(let i=0;i<v;i++){
    a[i] = 0
    b[i] = 0
}
dfs(0, b, 0, a)

let r = res.type.toString()
res.typeList.map((item, idx)=>{if(item) {r+=` ${idx+1}`}})
console.log(r);
process.exit() // 请注意必须在出口点处加入此行
import benchmark from "../../../util/benchmark.js"

// time complexity O(n)
function addUptoA(n) {
  let total = 0
  for (let i = 0; i <= n; i++) {
    total += i
  }
  return total
}

// tc: O(1)
function addUptoB(n) {
  return n * (n + 1) / 2
}

console.log(addUptoA(25))
console.log(addUptoA(1000))
benchmark("addUptoA", addUptoA, [10000000000])

console.log(addUptoB(25))
console.log(addUptoB(1000))
benchmark("addUptoB", addUptoB, [10000000000])
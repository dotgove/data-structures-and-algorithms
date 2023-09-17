// O(n)
function countUpAndDown(n) {
  for (let i = 0; i < n; i++) {
    console.log(i)
  }
  console.log("At the top. Going down...")
  for (let i = n - 1; i >= 0; i--) {
    console.log(i)
  }
  console.log("Back down. Exiting...")
}

countUpAndDown(10)
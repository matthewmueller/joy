async function main(ms) {
  await sleep(ms)
}

function sleep(ms) {
  return new Promise(function(resolve) {
    setTimeout(resolve, ms)
  })
}

main(10000).then(function() {
  console.log('all done')
})

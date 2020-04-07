require('ts-node').register({})


// const ok = (e)=> {
//   console.log(e)
//   console.log(chalk.blue("done"))
// }

if (process.argv.length>3) {
  console.log("• production mode")
  require('./scripts/index').build(false)//.then(ok).catch(ok)
} else {
  console.log("• development mode")
  require('./scripts/index').build(true) //.then(ok).catch(ok)
}


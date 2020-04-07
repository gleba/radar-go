export const indexHtml = (head, root, scripts, hash) => `<!DOCTYPE html>
<html>
  <head>
    ${head}
  </head>
  <body>
    <div id="hud"></div>
    <div id="root">${root}</div>
    ${hash ? setRouteHash(hash) : ''}
    ${scripts}
  </body>
</html>
`

const setRouteHash = hash => `<script>window.routeHash = ${JSON.stringify(hash)}</script>`

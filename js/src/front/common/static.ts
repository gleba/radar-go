const staticHost = '/'
export const res = Object.assign((resourceName) => staticHost + resourceName, {
  logo: staticHost + 'logo.svg',
})

export const setCssVar = (key, value) =>
  document.documentElement.style.setProperty(key, value)

export const domain = 'radar.cash'

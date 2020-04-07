// import { h, Fragment } from 'preact'
// import { useCallback } from 'preact/hooks'

export const ssrNav = route => e => {
    e.preventDefault()
    const atom = window['ssrRoute']
    atom ? atom(route) : console.log("empty store")
}

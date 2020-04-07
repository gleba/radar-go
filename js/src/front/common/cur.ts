export function money(v, lv) {
  return parseInt(v).toLocaleString(navigator.language, {
    style: 'currency',
    currencyDisplay: 'symbol',
    currency: lv,
    minimumSignificantDigits: 1,
  })
}

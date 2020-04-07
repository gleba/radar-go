declare type MinMax = {
  min: number
  max: number
}
declare type CapVol = {
  cap: MinMax
  vol: MinMax
}
declare type CVLimits = LV<CapVol>

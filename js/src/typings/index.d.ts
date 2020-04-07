///<reference path="telegram.d.ts"/>
declare type KV<T> = {
  [s: string]: T
}
declare type TKV<K, V> = {
  [s: K]: V
}

declare type RouteName = string

declare type RouteBox = {
  title: string
  component: any
  time: number
  state?: string
}

declare type RouteReq = {
  state: string
  time: number
}
declare type RouteResp = {
  state: string
  data: any
}

declare type AppRoutes = 'wrong-hash' | 'tune' | 'discover' | 'token' | 'scan'

declare type FrontCommand = 'route' | 'hash' | 'some'
declare type BackCommand = 'route' | 'auth' | 'wrong-hash' | 'some'

declare type FrontAction = [FrontCommand, any]
declare type BackAction = [BackCommand, any]

// declare type Route = ''

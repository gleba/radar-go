
import {SsrRoute, upRoutes} from "./ssr/upRoutes";
import {HttpResponse} from "uWebSockets.js";
import A from "alak";

export const ramRoutes: KV<SsrRoute> = {}
upRoutes().up(route => {
  // console.log("☆",route.path)
  ramRoutes[route.path] = route
})

export function writeRoute(route: SsrRoute, res: HttpResponse, status: string) {
  // console.log("writeRoute", route.path)
  route.headers.forEach(h =>
    res.writeHeader(h[0], h[1])
  )
  res.writeStatus(status)
  res.end(route.buffer)
}

import NATS from "nats";
import {sens} from "./backStore";

export const nc = NATS.connect({url: 'nats://localhost:4222', token: '2yKnjkfXCtA8ik2yKnjkfXCtA8ik'})

nc.subscribe('land', jsonData =>
  sens.atoms.routes.land(JSON.parse(jsonData))
)
nc.publish("front")
